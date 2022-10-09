
const resolvers = {
  Query: {
    hello
  }
};

export default resolvers;

async function hello(rootValue, args) {
  console.log('root: ' + JSON.stringify(rootValue));
  console.log('args: ' + JSON.stringify(args));
  return 'hello user';
}

