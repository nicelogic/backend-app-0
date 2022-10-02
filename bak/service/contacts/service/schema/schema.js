
import { loadSchemaSync } from '@graphql-tools/load';
import { GraphQLFileLoader } from '@graphql-tools/graphql-file-loader';
import { makeExecutableSchema } from "@graphql-tools/schema";
import contactsResolver from './resolvers/contacts_resolvers.js';
import { fileURLToPath } from 'url';
import { dirname } from 'path';

const __filename = fileURLToPath(import.meta.url);
const __dirname = dirname(__filename);
const typeDefs = loadSchemaSync(`${__dirname}/type_defs.graphql`, { loaders: [new GraphQLFileLoader()] });

const resolvers = {
  Query: {
    ...contactsResolver.Query,
  },
  Mutation:{
    ...contactsResolver.Mutation
  }
};

const schema = makeExecutableSchema({ typeDefs, resolvers });

export default schema;
