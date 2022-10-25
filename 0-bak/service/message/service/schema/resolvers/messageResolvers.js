
import { AccountModel, ChatModel, MessageModel } from "../../db/account.js"
import { PubSub, withFilter } from "graphql-subscriptions";

// const GET_CHAT_SUB = 'GET_CHAT_SUB';
const NEW_MESSAGE_RECEIVED = 'NEW_MESSAGE_RECEIVED';
const pubSub = new PubSub();

const resolvers = {
  Mutation: {
    createMessage,
  },
  Query: {
    getMessages,
  },
  Subscription: {
    // getNewMessages: {
    //   subscribe: withFilter(
    //     () => pubSub.asyncIterator([GET_CHAT_SUB]),
    //     (payload, variable) => {
    //       return payload.chatId === variable.chatId;
    //     }
    //   )
    // },
    newMessageReceived: {
      subscribe: withFilter(
        () => pubSub.asyncIterator([NEW_MESSAGE_RECEIVED]),
        (payload, variable) => {
          const senderId = payload.newMessageReceived.senderId;
          const members = payload.newMessageReceived.members.toObject();
          const accountId = variable.accountId;
          const needPublish = members.indexOf(accountId) != -1 && senderId != accountId;
          return needPublish;
        }
      )
    }
  },
};

async function createMessage(_, { accountId, chatId, text }) {
  const account = await AccountModel.findById(accountId);
  const chatModel = await ChatModel.findById(chatId).populate('messages');
  const messageModel = new MessageModel({ text: text, chat: chatModel._id, sender: account._id });
  chatModel.lastMessage = messageModel;
  await messageModel.save();
  await chatModel.messages.push(messageModel);
  await chatModel.save();
  const message = (await messageModel.populate('sender')).toObject();
  console.log(`account: ${accountId} send: ${text} in chat: ${chatId}`);
  pubSub.publish(NEW_MESSAGE_RECEIVED, {
    newMessageReceived: {
      _id: chatModel.id,
      message: message,
      members: chatModel.members,
      senderId: accountId
    }
  });
  return message;
}

async function getMessages(_, { chatId }) {
  let chatModel = await (await ChatModel.findById(chatId)).populate({ path: 'messages', populate: { path: 'sender' } });
  chatModel = (await chatModel.populate({ path: 'lastMessage', populate: { path: 'sender' } }));
  chatModel = await chatModel.populate('members');
  const chat = chatModel.toObject();
  console.log(`get chat${chatId} messages`);
  return chat;
}

export default resolvers;
