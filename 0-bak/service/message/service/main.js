// @ts-check
import { createServer } from "http";
import express from "express";
import { execute, subscribe } from "graphql";
import { ApolloServer } from "apollo-server-express";
import { SubscriptionServer } from "subscriptions-transport-ws";
import { initDb } from "./db/account.js";
import schema  from './schema/schema.js';
import cors from 'cors';

(async () => {

  initDb({dbName: 'account'});

  const PORT = 80;
  const app = express();
  app.use(cors());
  const httpServer = createServer(app);
  const server = new ApolloServer(
    { schema, }
  );
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
