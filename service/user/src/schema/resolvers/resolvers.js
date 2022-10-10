
import { getUserId } from 'nicelogic-auth';

const resolvers = {
  Query: {
    hello
  }
};

export default resolvers;

async function hello(rootValue) {
  console.log('root: ' + JSON.stringify(rootValue));
  const userId = getUserId(rootValue);
  console.log('user id: ' + userId);
  return 'hello user: ' + userId;
}

