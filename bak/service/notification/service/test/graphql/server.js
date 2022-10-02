
var express = require('express');
var { graphqlHTTP } = require('express-graphql');
var { buildSchema } = require('graphql');

// Construct a schema, using GraphQL schema language
var schema = buildSchema(`
  type Account {
    id: String!
    name: String!
    wechatId: String!
    signature: String!
  }
 
  type Query {
    account(id: String): Account
  }
`);

// This class implements the RandomDie GraphQL type
class Account {
    constructor(id) {
        this.id = id;
        this.name = 'niceice';
        this.wechatId = 'niceice220';
        this.signature = 'do is finished';
    }

    id() { return this.id; }
    name() { return this.name; }
    wechatId() { return this.wechatId; }
    signature() { return this.signature; }
}

// The root provides the top-level API endpoints
var root = {
    account: ({id}) => { return new Account(id); }
}

const { URLSearchParams } = require('url');
global.URLSearchParams = URLSearchParams;
var app = express();
app.use('/graphql', graphqlHTTP({
    schema: schema,
    rootValue: root,
    graphiql: true,
}));
app.listen(4000);
console.log('Running a GraphQL API server at localhost:4000/graphql');