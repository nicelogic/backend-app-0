
import Cassandra from 'nicelogic-cassandra';
import { gql } from 'graphql-request'
import { nanoid } from 'nanoid';
import crypto from 'crypto';
import { generateToken } from '../../token.js';
import { reportError } from 'nicelogic-error';
import { authErrors } from '../../error.js';
import { GraphQLError } from 'graphql';

const resolvers = {
  Query: {
    signInByUserName
  },
  Mutation: {
    signUpByUserName
  }
};
export default resolvers;

async function signUpByUserName(rootValue, { userName, pwd }) {
  try {
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

    const cassandra = new Cassandra();
    const response = await cassandra.mutation(insertAuth, variables);
    console.log(JSON.stringify(response));
    const isExist = response['insertauth']['applied'] === false;
    if (isExist) {
      throw new GraphQLError(authErrors.userExist);
    }
    const user_id = response['insertauth']['value']['user_id'];
    const token = generateToken(user_id, rootValue.privateKey, rootValue.expiresIn);
    const auth_id = response['insertauth']['value']['auth_id'];
    const auth_id_type = response['insertauth']['value']['auth_id_type'];
    const create_time = response['insertauth']['value']['create_time'];
    return {
      auth: {
        auth_id: auth_id,
        auth_id_type: auth_id_type,
        user_id: user_id,
        create_time: create_time
      },
      token: token
    };
  } catch (e) {
    console.log(e);
    reportError(e);
  }
}

async function signInByUserName(rootValue, { userName, pwd }) {
  try {
    console.log(`signin by user name: ${userName}`);

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

    const cassandra = new Cassandra();
    const response = await cassandra.query(queryAuth, variables);
    console.log(JSON.stringify(response));
    if (response['auth']['values'].length === 0) {
      throw authErrors.userNotExist;
    }
    const auth_id_type_username_pwd = response['auth']['values'][0]['auth_id_type_username_pwd'];
    const md5Pwd = crypto.createHash('md5').update(pwd, 'utf8').digest("hex");
    const isPwdRight = md5Pwd === auth_id_type_username_pwd;
    if (!isPwdRight) {
      throw new GraphQLError(authErrors.pwdWrong);
    }
    const user_id = response['auth']['values'][0]['user_id'];
    const token = generateToken(user_id, rootValue.privateKey, rootValue.expiresIn);
    const auth_id = response['auth']['values'][0]['auth_id'];
    const auth_id_type = response['auth']['values'][0]['auth_id_type'];
    const create_time = response['auth']['values'][0]['create_time'];
    return {
      auth: {
        auth_id: auth_id,
        auth_id_type: auth_id_type,
        user_id: user_id,
        create_time: create_time
      },
      token: token
    };
  } catch (e) {
    console.log(e);
    reportError(e);
  }
}
