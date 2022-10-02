
import { generateToken } from "../../util/token.js";

const resolvers = {
  Query: {
    signInWithUserName
  },
  Mutation: {
    signUpWithUserName
  }
};
export default resolvers;

async function signUpWithUserName(_, { userName, pwd}) {
  console.log(`signup with user name: ${userName}`);
  return "aaa";
}

async function signInWithUserName(_, { userNamePwd }) {
  return "";
}
