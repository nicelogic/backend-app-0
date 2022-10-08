
import { graphqlHTTP } from 'express-graphql';
import express from 'express';
import schema from './schema/schema.js';
import cors from 'cors';
import Config from 'nicelogic-config';
import fs from 'fs';
import { expressjwt } from 'express-jwt';

async function main() {


  const config = new Config('/etc/app-0/config-user/config-user.yml');
  const path = config.get('path', '/');
  const publicKey = fs.readFileSync('/etc/app-0/secret-jwt/jwt-publickey');

  const rootValue = {
    publicKey: publicKey
  };

  const extensions = ({
    document,
    variables,
    operationName,
    result,
    context,
  }) => {
    return {
      runTime: Date.now() - context.startTime,
    };
  };

  const app = express();
  app.use(
    expressjwt({
      secret: publicKey,
      algorithms: ['RS256'],
      credentialsRequired: false,
    }).unless({
      path: []
    }),
  );
  app.use(cors());
  app.use(path, graphqlHTTP( (request, response, graphQLParams) => ({
    schema: schema,
    rootValue: rootValue,
    pretty: true,
    graphiql: { headerEditorEnabled: true },
    context: { startTime: Date.now() },
    extensions
  })));
  app.listen(80);
  console.log(`🚀 Running a GraphQL API server at http://localhost${path}`);

}

main().catch((error) => console.error(error));

