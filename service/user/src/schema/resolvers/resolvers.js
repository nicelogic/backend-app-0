
const resolvers = {
  Query: {
    hello
  }
};

export default resolvers;

async function hello(_, args, context) {
  console.log(JSON.stringify(parent));
  console.log(JSON.stringify(args));
  console.log(JSON.stringify(context));
  return 'hello user';
}

