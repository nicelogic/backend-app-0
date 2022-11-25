
import { AuthModel } from "../../db/schema.js";
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
  const query = await AuthModel.findOne({
    'userNamePwdAuth.userName': userName
  });
  if (query !== null) {
    const error = 'username already exists';
    console.error(error);
    throw error;
  }
  const authInfo = new AuthModel({ userNamePwdAuth: {
    userName: userName,
    pwd: pwd
  }});
  authInfo.save().then(() => console.log(`sign up with user name, has save to db, id: ${authInfo._id}`));

	let key = fileSystem.readFileSync('./cert/2_niceice.cn.key');
  const token = generateToken(authInfo._id, key);
  return token;
}

async function signInWithUserName(_, { userNamePwd }) {
  return '';
}

export default resolvers;