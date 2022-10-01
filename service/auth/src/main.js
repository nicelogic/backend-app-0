// @ts-check
import { createServer } from "http";
import express from "express";
import { ApolloServer } from "apollo-server-express";
import schema  from './schema/schema.js';
import cors from 'cors';
import yaml from 'js-yaml';
import fs from 'fs';

(async () => {

  const configFilePath = '/etc/warmth/config.yml';
  const config = yaml.load(fs.readFileSync(configFilePath, 'utf8'));
  console.log(config);
  const port = config['port'];
  const path = config['path'];

  const app = express();
  app.use(cors());
  const httpServer = createServer(app);
  const server = new ApolloServer(
    { schema, }
  );
  await server.start();
  server.applyMiddleware({
    app,
    path: path
  });

  httpServer.listen(port, () => {
    console.log(
      `🚀 Query endpoint ready at http://localhost:${port}${server.graphqlPath}`
    );
  });
})();
