
let express = require('express');
let https = require('https');
let fileSystem = require('fs');
let { graphqlHTTP } = require('express-graphql');
let { buildSchema } = require('graphql');
const expressJwt = require('express-jwt')
const mergeJSON = require("merge-json");
const cors = require('cors')
const admin = require("firebase-admin");

admin.initializeApp(admin.credential.applicationDefault());

let schema = buildSchema(`
  type Notification{
    id: ID!
    sender: String!
    receivers: [String!]!
    content: String!
  }

  input NotificationInput {
    id: ID!
    sender: String!
    receivers: [String!]!
    content: String!
  }

  type Query {
    hello: String!
  }

  type Mutation {
    registerDevice(userId: ID!, token: String!, deviceInfo: String!): String!
    sendNotification(notification: NotificationInput!): Notification!
  }

`);

class Notification {
  constructor({ userId, token, deviceInfo }) {
    this.userId = userId;
    this.token = token;
    this.deviceInfo = deviceInfo;
  }

  userId() { return this.userId; }
  token() { return this.token; }
  deviceInfo() { return this.deviceInfo; }
}

let root = {
  hello: ()=>{
    return "hello";
  },

  registerDevice: async ({ userId, token, deviceInfo }) => {
    console.log(`registerDevice<userId:${userId}, token:${token}, deviceInfo:${deviceInfo}>`);

    return deviceInfo;
  },
  sendNotification: async ({ notification }) => {
    // This registration token comes from the client FCM SDKs.
    const registrationToken = 'e79Q-oFfTZWk2OSM3dYIIE:APA91bEe3zOtW_XvopWN_XiJsLoO8Cynf9221w9X-Uwdx3VJMYPz3qlr3WKj3JPP-2yN2uLbx4heL3Bkum1BH5ImiYHoOaJOHX_Aa2_t4sfAo8gWI7fuoRiBOfCicA_EVktBVoHIyn12';

    const message = {
      notification: {
        title: 'on the day',
        body: 'FooCorp gained 11.80 points to close at 835.67, up 1.43% on the day.'
      },
      android: {
        notification: {
          icon: 'stock_ticker_update',
          color: '#7e55c3'
        }
      },
      token: registrationToken
    };

    // Send a message to the device corresponding to the provided
    // registration token.
    admin.messaging().send(message)
      .then((response) => {
        // Response is a message ID string.
        console.log('Successfully sent message:', response);
      })
      .catch((error) => {
        console.log('Error sending message:', error);
      });
    return notification;
  },
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
console.log('Running a  GraphQL API server at https://niceice.cn/notification');