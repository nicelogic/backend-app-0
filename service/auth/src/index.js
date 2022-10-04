
import { graphqlHTTP } from 'express-graphql';
import express from 'express';
import schema from './schema/schema.js';
import cors from 'cors';
import Config from 'nicelogic-config';

(async () => {

  const serviceConfigFilePath = '/etc/app-0/config-auth/config-auth.yml';
  const config = new Config(serviceConfigFilePath);
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
  const port = 80;
  app.listen(port);
  console.log(`🚀 Running a GraphQL API server at http://localhost:${port}${path}`);

})();
