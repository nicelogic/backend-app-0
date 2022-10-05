
import Config from 'nicelogic-config';
import Cassandra from 'nicelogic-cassandra';
import { GraphQLClient, gql } from 'graphql-request'
import { nanoid } from 'nanoid';


const resolvers = {
  Query: {
    signInByUserName
  },
  Mutation: {
    signUpByUserName
  }
};
export default resolvers;

async function signUpByUserName(_, { userName, pwd }) {
  console.log(`signup by user name: ${userName}`);
  const cassandra = new Cassandra();
  const token = await cassandra.getToken();
  console.log(`token: ${token}`);

  const config = new Config('/etc/app-0/config/config.yml');
  const cassandraGraphqlUrl = config.get('cassandra-graphql-url-app_0', '');
  const graphQLClient = new GraphQLClient(cassandraGraphqlUrl,
    {
      method: 'POST',
      jsonSerializer: {
        parse: JSON.parse,
        stringify: JSON.stringify,
      },
      headers: {
        'x-cassandra-token': token,
      },
    });

  const userId = nanoid();
  const mutation = gql`
    mutation insertauth($auth_id: String!, $auth_id_type_username_pwd: String!, $user_id: String!) {
      insertauth(value: {
                    auth_id: $auth_id, 
                    auth_id_type: "username",
                    auth_id_type_username_pwd: $auth_id_type_username_pwd,
                    user_id: $user_id
                  },
                  ifNotExists: true
                  ) {
          applied,
          accepted,
          value {
            auth_id,
            user_id
          }
        }
    }
  `;
  const variables = {
    auth_id: userName,
    auth_id_type_username_pwd: pwd,
    user_id: userId
  };

  try {
    const data = await graphQLClient.request(mutation, variables);
    const jdata = JSON.stringify(data);
    console.log(jdata)

    return jdata;

  } catch (e) {
    console.log(e);
    return "error";
  }
}

async function signInByUserName(_, { userName, pwd }) {
  console.log(`signup by user name: ${userName}`);
  const cassandra = new Cassandra();
  const token = await cassandra.getToken();
  console.log(`token: ${token}`);

  const config = new Config('/etc/app-0/config/config.yml');
  const cassandraGraphqlUrl = config.get('cassandra-graphql-url-app_0', '');
  const graphQLClient = new GraphQLClient(cassandraGraphqlUrl,
    {
      method: 'GET',
      jsonSerializer: {
        parse: JSON.parse,
        stringify: JSON.stringify,
      },
      headers: {
        'x-cassandra-token': token,
      },
    });

  const query = gql`
query auth($auth_id: String!) {
  auth: auth(value: {
                    auth_id: $auth_id, 
                   
                  },
                  ) {
          pageState,
          values {
            auth_id,
						user_id,
            auth_id_type_username_pwd,
            createTime
          }
        }
}
  `;
  const variables = { auth_id: userName };

  try {
    const data = await graphQLClient.request(query, variables);
    const jdata = JSON.stringify(data);
    console.log(jdata)

    return jdata;

  } catch (e) {
    console.log(e);
    return "error";
  }
}
