// @ts-check
const { createServer } = require("http");
const express = require("express");
const { execute, subscribe } = require("graphql");
const { ApolloServer, gql } = require("apollo-server-express");
const { PubSub, withFilter } = require("graphql-subscriptions");
const { SubscriptionServer } = require("subscriptions-transport-ws");
const { makeExecutableSchema } = require("@graphql-tools/schema");
const { v4: uuidv4 } = require('uuid');
const db = require('./db/db.js');
const cors = require('cors');

(async () => {

  const PublicationCollection = db.initDb({ dbName: 'publication', collectionName: 'publication' });
  const pubsub = new PubSub();

  const typeDefs = gql`
    type Publication{
      _id: ID!
      accountId: ID!
      targetId: ID!
      event: String!
      info: String!
      state: String!
      replyEvent: String
      replyInfo: String
    }

    type Query {
      getPublications(accountId: ID!, event: String!): [Publication!]
    }

    type Mutation{
      publish(accountId: ID!,
        targetId: ID!,
        event: String!,
        info: String!,
        state: String!): Publication!
      replyPublish(id: ID!
        event: String!,
        info: String!,
        state: String!): Publication!
    }

    type Subscription {
      publicationReceived(accountId: ID!): Publication!
    }
  `;

  class Publication {
    constructor({ _id, accountId, targetId, event, info, state, replyEvent, replyInfo }) {
      this._id = _id;
      this.accountId = accountId;
      this.targetId = targetId;
      this.event = event;
      this.info = info;
      this.state = state;
      this.replyEvent = replyEvent;
      this.replyInfo = replyInfo;
    }
  }

  const resolvers = {
    Query: {
      async getPublications(_, { accountId, event }) {
        const dbPublications = await PublicationCollection.find({
          $or: [
            {
              $and: [
                { targetId: accountId },
                { event: event },
                { state: 'create' }]
            },
          ]
        });
        console.log(`result: ${JSON.stringify(dbPublications)}`);
        let publications = [];
        for (let document of dbPublications) {
          const publication = document.toObject();
          publications.push(new Publication(publication));
        }
        return publications;
      },
    },
    Mutation: {
      async publish(_, { accountId, targetId, event, info, state }) {

        const id = uuidv4();
        console.log(`account: ${accountId} publish(id: ${id}) to ${targetId}:(${event},${info},${state})`);
        let publication = new Publication({
          _id: id,
          accountId: accountId,
          targetId: targetId,
          event: event,
          info: info,
          state: state,
          replyEvent: '',
          replyInfo: ''
        });

        const samePublications = await PublicationCollection.find({
          $and: [
            { accountId: accountId },
            { targetId: targetId },
            { event: event },
            { info: info },
            { state: state }]
        });
        if (samePublications.length != 0) {
          publication = new Publication(samePublications[0]);
          console.log(`has duplicate publication, will not save`);
        } else {
          const pubCollection = new PublicationCollection(publication);
          await pubCollection.save();
          console.log('pubCollection has save');
        }

        pubsub.publish('PUBLICATION', { publicationReceived: publication });
        return publication;
      },
      async replyPublish(_, { id, event, info, state}) {

        console.log(`reply the publication(id: ${id})(${event},${info},${state})`);
        const dbPublication = await PublicationCollection.findById(id);
        console.log(`result: ${JSON.stringify(dbPublication)}`);
        await dbPublication.updateOne({ $set: { replyEvent: event, replyInfo: info, state: state } });
        const upatedPublication = await PublicationCollection.findById(id);
        const publication = new Publication(upatedPublication);
        pubsub.publish('PUBLICATION', { publicationReceived: publication });
        return publication;
      }
    },
    Subscription: {
      publicationReceived: {
        subscribe: withFilter(
          () => pubsub.asyncIterator(["PUBLICATION"]),
          (payload, variables) => {
            const payloadTargetId = payload.publicationReceived.targetId;
            // const payloadAccountId = payload.publicationReceived.accountId;
            const payloadState = payload.publicationReceived.state;
            const accountId = variables.accountId;
            const isTargetSubscribe = (payloadTargetId === accountId && payloadState == 'create');
            // const isSourceSubscribe = (payloadAccountId === accountId && payloadState != 'create');
            return isTargetSubscribe;
          })
      }
    },
  };

  const schema = makeExecutableSchema({ typeDefs, resolvers });

  const PORT = 80;
  const app = express();
  const httpServer = createServer(app);
  const server = new ApolloServer(
    { schema, }
  );
  app.use(cors());
  await server.start();
  server.applyMiddleware({
    app,
    path: '/'
  });

  SubscriptionServer.create(
    {
      schema, execute, subscribe,
      onConnect: () => {
        console.log(`onConnect`);
      },
      onDisconnect: () => {
        console.log(`onDisconnect`);
      },
      keepAlive: 10 * 1000
    },
    { server: httpServer, path: server.graphqlPath }
  );

  httpServer.listen(PORT, () => {
    console.log(
      `🚀 Query endpoint ready at http://localhost:${PORT}${server.graphqlPath}`
    );
    console.log(
      `🚀 Subscription endpoint ready at ws://localhost:${PORT}${server.graphqlPath}`
    );
  });
})();
