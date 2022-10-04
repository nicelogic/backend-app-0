
import { generateToken } from "../../util/token.js";
import axios from 'axios';

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
  const cassandraAdminNameAndPwd = '{"username": "cassandra-cluster-env0-superuser", "password": "znk4uVfaCLm6hppEZaJl"}';
  const response = await axios.post("https://auth.cassandra.env0.luojm.com:9443/v1/auth",
    cassandraAdminNameAndPwd,
    {
      headers: {
        'Content-Type': 'application/json'
      },
    }
  );
  const token = response.data.authToken;
  return token;
}

async function signInByUserName(_, { userNamePwd }) {
  return "";
}
