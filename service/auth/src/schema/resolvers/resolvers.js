
import Cassandra from 'nicelogic-cassandra';
import { gql } from 'graphql-request'
import { nanoid } from 'nanoid';
import crypto from 'crypto';
import { generateToken } from '../../util/token.js';

const resolvers = {
  Query: {
    signInByUserName
  },
  Mutation: {
    signUpByUserName
  }
};
export default resolvers;

async function signUpByUserName(context, { userName, pwd }) {
  console.log(`signup by user name: ${userName}`);
  const insertAuth = gql`
    mutation insertauth($auth_id: String!, $auth_id_type_username_pwd: String!, $user_id: String!, $create_time: Timestamp!) {
      insertauth(value: {
                    auth_id: $auth_id, 
                    auth_id_type: "username",
                    auth_id_type_username_pwd: $auth_id_type_username_pwd,
                    user_id: $user_id,
                    create_time: $create_time
                  },
                  ifNotExists: true
                  ) {
          applied,
          accepted,
          value {
            auth_id,
            auth_id_type,
            user_id,
            create_time
          }
        }
    }
  `;
  const userId = nanoid();
  const md5Pwd = crypto.createHash('md5').update(pwd, 'utf8').digest("hex");
  const createTime = (new Date()).toISOString();
  const variables = {
    auth_id: userName,
    auth_id_type_username_pwd: md5Pwd,
    user_id: userId,
    create_time: createTime
  };

  let error_code = 0;
  let error_code_description = '';
  let auth_id = userName;
  let auth_id_type = 'username';
  let user_id = '';
  let token = '';
  let create_time = '';
  try {
    const cassandra = new Cassandra();
    const response = await cassandra.mutation(insertAuth, variables);
    console.log(JSON.stringify(response));
    auth_id = response['insertauth']['value']['auth_id'];
    auth_id_type = response['insertauth']['value']['auth_id_type'];
    const isExist = response['insertauth']['applied'] === false;
    if (isExist) {
      error_code = 1;
      error_code_description = 'user name already exist';
    } else {
      user_id = response['insertauth']['value']['user_id'];
      create_time = response['insertauth']['value']['create_time'];
      token = generateToken(user_id, context.privateKey);
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
      user_id: user_id,
      create_time: create_time
    },
    token: token
  };
}

async function signInByUserName(context, { userName, pwd }) {
  console.log(`signup by user name: ${userName}`);

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
          user_id,
          create_time
        }
      }
  }
`;
  const variables = {
    auth_id: userName,
  };

  let error_code = 0;
  let error_code_description = '';
  let auth_id = userName;
  let auth_id_type = 'username';
  let user_id = '';
  let token = '';
  let create_time = '';
  try {
    const cassandra = new Cassandra();
    const response = await cassandra.query(queryAuth, variables);
    console.log(JSON.stringify(response));
    if (response['auth']['values'].length === 0) {
      error_code = 2;
      error_code_description = 'user not exist';
    } else {
      auth_id = response['auth']['values'][0]['auth_id'];
      auth_id_type = response['auth']['values'][0]['auth_id_type'];
      const auth_id_type_username_pwd = response['auth']['values'][0]['auth_id_type_username_pwd'];
      const md5Pwd = crypto.createHash('md5').update(pwd, 'utf8').digest("hex");
      const isPwdRight = md5Pwd === auth_id_type_username_pwd;
      if (isPwdRight) {
        user_id = response['auth']['values'][0]['user_id'];
        create_time = response['auth']['values'][0]['create_time'];
        token = generateToken(user_id, context.privateKey);
      } else {
        error_code = 3;
        error_code_description = 'password wrong';
      }
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
      user_id: user_id,
      create_time: create_time
    },
    token: token
  };
}
