
import Config from 'nicelogic-config';
import Cassandra from 'nicelogic-cassandra';
import { gql } from 'graphql-request'
import { nanoid } from 'nanoid';
import crypto from 'crypto';


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
            auth_id_type,
            user_id
          }
        }
    }
  `;
  const userId = nanoid();
  const md5Pwd = crypto.createHash('md5').update(pwd, 'utf8').digest("hex");
  const variables = {
    auth_id: userName,
    auth_id_type_username_pwd: md5Pwd,
    user_id: userId
  };

  let error_code = 0;
  let description = '';
  let auth_id = userName;
  let auth_id_type = 'username';
  let user_id = '';
  try {
    const response = await cassandra.mutation(mutation, variables);
    console.log(JSON.stringify(response));
    const isExist = response['insertauth']['applied'] === false;
    if (isExist) {
      error_code = 1;
      description = 'user name exist';
    }
    auth_id = response['insertauth']['value']['auth_id'];
    auth_id_type = response['insertauth']['value']['auth_id_type'];
    user_id = response['insertauth']['value']['user_id'];
  } catch (e) {
    console.log(e);
  }
  return {
    error_code: error_code,
    description: description,
    auth: {
      auth_id: auth_id,
      auth_id_type: auth_id_type,
      user_id: user_id
    }
  };
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

    return data;

  } catch (e) {
    console.log(e);
    return "error";
  }
}
