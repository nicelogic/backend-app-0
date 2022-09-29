
import { generateToken } from "../../util/token.js";

const resolvers = {
  Query: {
    hello,
    signInWithUserName
  },
  Mutation: {
    signUpWithUserName
  }
};

async function hello(_, { name }){
  return 'hello';
}

async function signUpWithUserName(_, { userName, pwd}) {
  console.log(`signup with user name: ${userName}`);
  return "";
}

async function signInWithUserName(_, { userNamePwd }) {
  return "";
}

export default resolvers;