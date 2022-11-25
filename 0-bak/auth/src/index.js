
import { graphqlHTTP } from 'express-graphql';
import express from 'express';
import schema from './schema/schema.js';
import cors from 'cors';
import Config from 'nicelogic-config';
import fs from 'fs';

async function main(){

  const config = new Config('/etc/app-0/config-auth/config-auth.yml');
  //const expiresIn = config.get('expiresin', 60);
  const expiresIn = 60;
  const privateKey = fs.readFileSync('/etc/app-0/secret-jwt/jwt-privatekey');
  const rootValue = {
    privateKey: privateKey,
    expiresIn: expiresIn
  };

  const app = express();
  app.use(cors());
  const path = config.get('path', '/');
  app.use(path, graphqlHTTP({
    schema: schema,
    rootValue: rootValue,
    pretty: true,
    graphiql: true,
  }));
  app.listen(80);
  console.log(`🚀 Running a GraphQL API server at http://localhost${path}`);

}

main().catch((error) => console.error(error));

