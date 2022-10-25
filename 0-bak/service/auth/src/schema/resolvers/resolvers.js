
import { generateToken } from "../../util/token.js";

const resolvers = {
  Query: {
    signInByUserName
  },
  Mutation: {
    signUpByUserName
  }
};
export default resolvers;

async function signUpByUserName(_, { userName, pwd}) {
  console.log(`signup by user name: ${userName}`);
  return "aaa";
}

async function signInByUserName(_, { userNamePwd }) {
  return "";
}
