
import Cassandra from "../index.js";
import { gql } from 'graphql-request'
import crypto from 'crypto';

const cassandra = new Cassandra();

async function testMutation() {
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
  const variables = {
    auth_id: "test",
    auth_id_type_username_pwd: "c",
    user_id: "123"
  };

  const response = await cassandra.mutation(mutation, variables);
  console.log(JSON.stringify(response));
}

async function testQuery(userName, pwd) {
  const queryAuth = gql`
  query auth($auth_id: String!) {
    auth(value: {
                  auth_id: $auth_id, 
                  auth_id_type: "username",
                },
                ) {
        pageState,
        values {
          auth_id,
          auth_id_type,
          auth_id_type_username_pwd,
          user_id
        }
      }
  }
`;
  const variables = {
    auth_id: userName,
  };

  let error_code = 0;
  let error_code_description = '';
  let auth_id = 'test';
  let auth_id_type = userName;
  let user_id = '';
  let token = '';
  try {
    const cassandra = new Cassandra();
    const response = await cassandra.query(queryAuth, variables);
    console.log(JSON.stringify(response));
    auth_id = response['auth']['values'][0]['auth_id'];
    auth_id_type = response['auth']['values'][0]['auth_id_type'];
    user_id = response['auth']['values'][0]['user_id'];
    const auth_id_type_username_pwd = response['auth']['values'][0]['auth_id_type_username_pwd'];
    const md5Pwd = crypto.createHash('md5').update(pwd, 'utf8').digest("hex");

    const isPwdRight = md5Pwd === auth_id_type_username_pwd;
    if (isPwdRight) {
      token = '';
    } else {
      error_code = 2;
      error_code_description = 'password wrong';
    }
  } catch (e) {
    error_code = -1;
    error_code_description = 'server internal error';
    console.log(e);
  }
  return {
    error_code: error_code,
    error_code_description: error_code_description,
    auth: {
      auth_id: auth_id,
      auth_id_type: auth_id_type,
      user_id: user_id
    },
    token: token
  };
}

const userName = 'test';
const pwd = 'c';
const response = await testQuery(userName, pwd);
console.log(JSON.stringify(response));

