
import { AccountModel, ChatModel } from "../../db/account.js"

const resolvers = {
  Mutation: {
    createChat,
    deleteChat,
  },
  Query: {
    getChats,
    getChat,
  },
};

async function createChat(
  _,
  { accountId, memberIds, name }
) {
  let account;
  let chats;
  try {
    account = await (await AccountModel.findById(accountId)).populate('chats');
    console.log(`account find by id: ${accountId}\n`);
    console.log(JSON.stringify(account));
    chats = account.chats;
    console.log(JSON.stringify(chats));
    for (let chatModel of chats) {
      const chatMembers = chatModel.members.toObject().sort().toString();
      const createChatMembers = [...memberIds, account._id].sort().toString();
      if (chatMembers === createChatMembers) {
        let chat = (await chatModel.populate('members'));
        chat = (await chat.populate({ path: 'lastMessage', populate: { path: 'sender' } }));
        chat = (await chat.populate({ path: 'messages', populate: { path: 'sender' } })).toObject();
        console.log(`exist repeat member chat, return this chat\n`);
        return chat;
      }
    }
  } catch (e) {
    console.log(`account populate chat failure(may old db structure), controuct a new one\n`);
    account = await AccountModel.findById(accountId);
    console.log(`account find by id: ${accountId}\n`);
    console.log(JSON.stringify(account));
    chats = account.chats;
  }

  const chatModel = new ChatModel({
    name: name,
    members: [...memberIds, account._id]
  });
  for (let memberId of memberIds) {
    try {
      const memberAccount = await AccountModel.findById(memberId);
      await memberAccount.chats.push(chatModel._id);
      await memberAccount.save();
    } catch (e) {
      throw (`create account failure: member account may invalid`);
    }
  }
  await chatModel.save();
  await chats.push(chatModel._id);
  await account.save();
  console.log(`account and chat model save success`);
  console.log(JSON.stringify(account));

  console.log(`begin to get chat:${chatModel._id}`);
  const resultChatModel = await (await ChatModel.findById(chatModel._id)).populate('members');
  const chat = resultChatModel.toObject();
  console.log(`return chat: ${JSON.stringify(chat)}`)
  return chat;
}

async function getChats(_, { accountId }) {
  const account = await (await AccountModel.findById(accountId)).populate('chats');
  console.log(JSON.stringify(account));
  const chatsModel = account.chats;
  let chats = [];
  for (let chatModel of chatsModel) {
    let chat = await chatModel.populate('members');
    chat = (await chat.populate({ path: 'lastMessage', populate: { path: 'sender' } }));
    chat = (await chat.populate({ path: 'messages', populate: { path: 'sender' } })).toObject();
    console.log(JSON.stringify(chat));
    chats.push(chat);
  }
  return chats;
}

async function getChat(
  _,
  { chatId }
) {
  const chatModel = await ChatModel.findById(chatId);
  const chat = (await chatModel.populate('members')).toObject();
  console.log(`return chat: ${JSON.stringify(chat)}`)
  return chat;
}


async function deleteChat(_, { chatId }) {
  ChatModel.findByIdAndRemove(chatId);
  return chatId;
}

export default resolvers;
