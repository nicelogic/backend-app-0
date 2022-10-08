
let express = require('express');
let https = require('https');
let fileSystem = require('fs');
let { graphqlHTTP } = require('express-graphql');
let { buildSchema } = require('graphql');
const expressJwt = require('express-jwt')
const mergeJSON = require("merge-json");
const { initDb, AccountModel } = require("./db/account");
const cors = require('cors')

initDb({dbName: 'account'});

let schema = buildSchema(`
  type Contact {
    id: ID!
  }

  type Account {
    id: ID!
    name: String
    info: String
    contact: [Contact!]
  }

  input ContactInput {
    id: ID!
    event: String!
  }


  type Query {
    account(id: ID!): Account!
    queryAccount(idOrName: String!): [Account!]
  }

  type Mutation {
    createAccount(id: ID!, info: String!): Account!
    updateAccount(id: ID!, info: String, contactInput: ContactInput): Account!
  }

`);

class Account {
  constructor({ id, name, info, contact }) {
    this.id = id;
    this.name = name;
    this.info = info;
    this.contact = contact;
  }

  id() { return this.id; }
  name() { return this.name; }
  info() { return this.info; }
  contact() { return this.contact; }
}

function toAccount(id, accountResult) {
  const account = new Account({
    id: id,
    name: accountResult.toObject().name,
    info: accountResult.toObject().info,
    contact: accountResult.toObject().contact
  });
  console.log(`id: ${id}`);
  console.log(`name: ${accountResult.toObject().name}`);
  console.log(`contact: ${accountResult.toObject().contact}`);
  return account;
}

let root = {
  account: async ({ id }) => {
    console.log(`get account<id:${id}>`);
    const accountResult = await AccountModel.findById(id);
    if (accountResult == null) {
      throw `user does not exist <id:${id}>`;
    }

    const account = toAccount(id, accountResult);
    return account;
  },
  createAccount: async ({ id, info }) => {
    if(id.length() < 3){
      throw `user id: ${id} length must >= 3`;
    }

    const query = await AccountModel.findById(id);
    if (query != null) {
      throw `user<id:${id}> exist`;
    }
    const accountInfo = new AccountModel({ _id: id, info: info });
    accountInfo.save().then(() => console.log('account info has save'));
    const account = new Account({ id: id, info: info });
    return account;
  },
  updateAccount: async ({ id, info, contactInput }) => {
    const queryResult = await AccountModel.findById(id);
    if (queryResult == null) {
      throw `user does not exist <id:${id}>`;
    }

    console.log(`updateAccount: info: ${info}, contactInput: ${contactInput}`);
    if (info != undefined) {
      const currentAccountInfo = queryResult.toObject().info;
      const currentAccountInfoJson = JSON.parse(currentAccountInfo);
      const newInfo = JSON.parse(info);
      console.log(`current account info: ${currentAccountInfo}`);
      console.log(`need merge info: ${info}`);
      const accountInfoJson = mergeJSON.merge(currentAccountInfoJson, newInfo);
      const accountInfo = JSON.stringify(accountInfoJson);
      console.log(`after merge, account info: ${accountInfo}`)
      await queryResult.updateOne({ info: accountInfo });
      console.log('account info has save');
      const name = newInfo.name;
      console.log(`name: ${name}`)
      if (name != undefined) {
        await queryResult.updateOne({ name: name }, { upsert: true });
        console.log(`accout name(${name}) has save`);
      }
    }
    if (contactInput != undefined) {
      const event = contactInput.event;
      const contactId = contactInput.id;
      if (event === 'add_contact') {
        await queryResult.updateOne({
          $push: {
            contact: {
              id: contactId
            }
          }
        });
        console.log('add_contact success');
      }
    }

    const accountResult = await AccountModel.findById(id);
    if (accountResult == null) {
      throw `user does not exist <id:${id}>`;
    }
    const account = toAccount(id, accountResult);
    return account;
  },
  queryAccount: async ({ idOrName }) => {
    const containIdOrName = eval(`/${idOrName}/`);
    const query = await AccountModel.find({ $or: [{ _id: containIdOrName }, { name: containIdOrName }] });
    console.log(`result: ${query}`);
    let accountList = [];
    for (let document of query) {
      const accountInfo = document.toObject();
      const account = new Account({ id: document.id, name: accountInfo.name, info: accountInfo.info });
      accountList.push(account);
    }
    return accountList;
  }
}

let options = {
  key: fileSystem.readFileSync('./cert/2_niceice.cn.key'),
  cert: fileSystem.readFileSync('./cert/1_niceice.cn_bundle.crt')
}

const { URLSearchParams } = require('url');
global.URLSearchParams = URLSearchParams;
let app = express();
let httpsServer = https.createServer(options, app);
const secretKey = options.cert;
console.log(`key: ${secretKey}`);

app.use(
  expressJwt({
    secret: secretKey,
    algorithms: ['RS256']

  }).unless({
    path: ['/test']
  }),
  cors(),
  graphqlHTTP({
    schema: schema,
    rootValue: root,
    graphiql: true,
  }));

app.listen(80);
httpsServer.listen(443);
console.log('Running a  GraphQL API server at https://niceice.cn/account');