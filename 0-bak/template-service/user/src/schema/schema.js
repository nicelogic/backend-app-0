
import { loadSchemaSync } from '@graphql-tools/load';
import { GraphQLFileLoader } from '@graphql-tools/graphql-file-loader';
import { makeExecutableSchema } from "@graphql-tools/schema";
import resolver from './resolvers/resolvers.js';
import { fileURLToPath } from 'url';
import { dirname } from 'path';

const __filename = fileURLToPath(import.meta.url);
const __dirname = dirname(__filename);
const typeDefs = loadSchemaSync(`${__dirname}/type_defs.gql`, { loaders: [new GraphQLFileLoader()] });

const resolvers = {
  Query: {
    ...resolver.Query,
  }
};

const schema = makeExecutableSchema({ typeDefs, resolvers });

export default schema;
