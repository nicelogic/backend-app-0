
import { graphqlHTTP } from 'express-graphql';
import { GraphQLError } from 'graphql';
import express from 'express';
import schema from './schema/schema.js';
import cors from 'cors';
import Config from 'nicelogic-config';
import fs from 'fs';
import jwt from 'jsonwebtoken';

async function main() {


  const config = new Config('/etc/app-0/config-user/config-user.yml');
  const path = config.get('path', '/');
  const publicKey = fs.readFileSync('/etc/app-0/secret-jwt/jwt-publickey');

  const app = express();
  function rootValue(req) {
    const authorizationHeader = req.headers && 'Authorization' in req.headers ? 'Authorization' : 'authorization';
    if (typeof (req.headers[authorizationHeader]) === undefined) {
      throw new GraphQLError('unauthorized');
    }
    const token = (req.headers[authorizationHeader]).split(' ')[1];
    console.log(token);
    let payload = {};
    try {
      payload = jwt.verify(token, publicKey);
      console.log(payload.id);

    } catch (e) {
      throw new GraphQLError(e.message);
    }
    return payload;
  }

  app.use(cors());
  app.use(path, graphqlHTTP((req, response, graphQLParams) => ({
    schema: schema,
    pretty: true,
    graphiql: { headerEditorEnabled: true },
    rootValue: rootValue(req)
  })));
  app.listen(80);
  console.log(`🚀 Running a GraphQL API server at http://localhost${path}`);

}

main().catch((error) => console.error(error));

