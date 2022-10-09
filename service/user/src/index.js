
import { graphqlHTTP } from 'express-graphql';
import express from 'express';
import schema from './schema/schema.js';
import cors from 'cors';
import Config from 'nicelogic-config';
import fs from 'fs';
import { getJwtPayload } from 'nicelogic-auth';

async function main() {

  const config = new Config('/etc/app-0/config-user/config-user.yml');
  const path = config.get('path', '/');
  const publicKey = fs.readFileSync('/etc/app-0/secret-jwt/jwt-publickey');

  const app = express();


  app.use(cors());
  app.use(path, graphqlHTTP(async (req, response, graphQLParams) => ({
    schema: schema,
    pretty: true,
    graphiql: { headerEditorEnabled: true },
    rootValue: getJwtPayload(req, publicKey)
  })));
  app.listen(80);
  console.log(`🚀 Running a GraphQL API server at http://localhost${path}`);

}

main().catch((error) => console.error(error));

