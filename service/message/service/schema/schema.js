
import { loadSchemaSync } from '@graphql-tools/load';
import { GraphQLFileLoader } from '@graphql-tools/graphql-file-loader';
import { makeExecutableSchema } from "@graphql-tools/schema";
import chatResolver from './resolvers/chatResolvers.js';
import messageResolver from './resolvers/messageResolvers.js';
import { fileURLToPath } from 'url';
import { dirname } from 'path';

const __filename = fileURLToPath(import.meta.url);
const __dirname = dirname(__filename);
const typeDefs = loadSchemaSync(`${__dirname}/typeDefs.graphql`, { loaders: [new GraphQLFileLoader()] });

const resolvers = {
  Query: {
    ...chatResolver.Query,
    ...messageResolver.Query,
  },
  Mutation: {
    ...chatResolver.Mutation,
    ...messageResolver.Mutation,
  },
  Subscription: {
    ...messageResolver.Subscription
  }
};

const schema = makeExecutableSchema({ typeDefs, resolvers });

export default schema;
