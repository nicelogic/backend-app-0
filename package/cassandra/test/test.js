
import Cassandra from "../index.js";
import { gql } from 'graphql-request'

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
const variables = {
auth_id: "test",
auth_id_type_username_pwd: "c",
user_id: "123"
};

const response = await cassandra.mutation(mutation, variables);
console.log(JSON.stringify(response));
  