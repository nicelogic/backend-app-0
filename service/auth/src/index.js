
import { graphqlHTTP } from 'express-graphql';
import express from 'express';
import schema from './schema/schema.js';
import cors from 'cors';
import Config from 'nicelogic-config';

async function main(){

  const config = new Config('/etc/app-0/config-auth/config-auth.yml');
  const path = config.get('path', '/');

  const root = {
    hello: () => {
      return 'auth';
    },
  };
  const app = express();
  app.use(cors());
  app.use(path, graphqlHTTP({
    schema: schema,
    rootValue: root,
    pretty: true,
    graphiql: true
  }));
  app.listen(80);
  console.log(`🚀 Running a GraphQL API server at http://localhost${path}`);

}

main().catch((error) => console.error(error));

