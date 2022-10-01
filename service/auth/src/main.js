// @ts-check

import { graphqlHTTP } from 'express-graphql';
import express from 'express';
import schema from './schema/schema.js';
import cors from 'cors';
import yaml from 'js-yaml';
import fs from 'fs';

(async () => {

  const configFilePath = '/etc/app-0/config.yml';
  const config = yaml.load(fs.readFileSync(configFilePath, 'utf8'));
  console.log(config);
  const port = config['port'];
  const path = config['path'];

  const root = {
    hello: () => {
      return 'Hello world!';
    },
  };
  const app = express();
  app.use(cors());
  app.use(path, graphqlHTTP({
    schema: schema,
    rootValue: root,
    graphiql: true
  }));
  app.listen(port);
  console.log(`🚀 Running a GraphQL API server at http://localhost:${port}${path}`);

})();
