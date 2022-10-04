
import Config from 'nicelogic-config';
import Cassandra from 'nicelogic-cassandra';
import { request, gql } from 'graphql-request'
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

  const userId = nanoid();
  const insertAuth = gql`
    mutation insertauth {
      ${userName}: insertauth(value: {
                        auth_id:$userName, 
                        auth_id_type:"username",
                        auth_id_type_username_pwd: ${pwd},
                        user_id: ${userId}
                      },
                      ifNotExists: true
                      ) {
              applied,
              accepted,
              value {
                auth_id,
              }
            }
    }
  `;

  request('https://api.spacex.land/graphql/', insertAuth).then((data) => console.log(data))


  return token;
}

async function signInByUserName(_, { userNamePwd }) {
  return "";
}
