
import { UserModel, ContactsModel } from "../../db/db.js";
import { kContactsCollection } from '../../constant/constant.js';
import mongoose from 'mongoose';
import { inspect } from 'util';

const resolvers = {
  Query: {
    hello,
    paginationContacts
  },
  Mutation: {
    addContacts,
    removeContacts
  }
};


async function hello(_, { val }) {
  const paginationContacts = await ContactsModel.aggregate([{
    $match: { userId: 'wzh' },
  }, {
    $lookup: {
      from: 'accounts',
      localField: 'contacts',
      foreignField: '_id',
      as: 'user'
    }
  }, {
    $unwind: "$user"
  }, {
    $sort: { 'user.name': 1 }
  }, {
    $limit: 2
  }]);
  console.log(inspect(paginationContacts, { depth: null }));

  // const contactsModel = await ContactsModel.find({ userId: 'wzh' }).populate(
  //   { path: 'contacts', select: { name: 1 }, options: { sort: { id: 1 } } });




  // const userModel = await (await UserModel.findById('wzh')).populate({
  //   path: 'contacts', populate: {
  //     path: 'contacts',
  //     select: { 'name': 1 },
  //   },
  //   // match: { 'contacts.name': { $gt: val } },
  //   options: { limit: 2, sort: { contacts.name: 1 } }

  // });
  // if (userModel === null) {
  //   return 'is null';
  // } else {
  //   console.log(inspect(userModel, { depth: null }));

  //   return 'not null';
  // }
}


async function paginationContacts(
  _,
  { userId, first, after }
) {
  let id = '000000000000000000000000';
  let name = '';
  if (after != '') {
    const decodeCursorInfo = Buffer.from(after, 'base64').toString('ascii');
    const cursorInfo = decodeCursorInfo.split(',');
    if (cursorInfo.length === 2) {
      name = cursorInfo[0];
      id = cursorInfo[1];
      console.log(`pagination cursor name: ${name}, id: ${id}`);
    }
  }
  let paginationContacts = [];
  try {
    paginationContacts = await ContactsModel.aggregate([
      {
        $lookup: {
          from: 'accounts',
          localField: 'contacts',
          foreignField: '_id',
          as: 'user'
        }
      },
      {
        $match: {
          $and: [{ userId: userId },
          {
            $or: [{
              'user.name': {
                $gt: name
              }
            }, {
              'user.name': name,
              _id: { $gt: mongoose.Types.ObjectId(id) }
            }
            ]
          }
          ]
        },
      }, {
        $unwind: "$user"
      }, {
        $sort: { 'user.name': 1, _id: 1 }
      }, {
        $limit: first + 1
      }]);
  } catch (err) {
    console.log(err);
  }
  const hasNextPage = paginationContacts.length >= first + 1;
  if (hasNextPage) {
    paginationContacts = paginationContacts.slice(0, paginationContacts.length - 1);
  }

  console.log(inspect(paginationContacts, { depth: null }));
  const edges =
    paginationContacts.map((contacts) => {
      return {
        node: {
          id: contacts.contacts,
          name: contacts.user.name
        },
        cursor: Buffer.from(contacts.user.name + ',' + contacts._id).toString('base64'),
      };
    });

  let endCursor = '';
  if (paginationContacts.length !== 0) {
    const lastPaginationContacts = paginationContacts.at(-1);
    endCursor = lastPaginationContacts.user.name + ',' + lastPaginationContacts._id;
    endCursor = Buffer.from(endCursor).toString('base64');
  }
  return {
    id: userId,
    contactsConnection: {
      totalCount: paginationContacts.length,
      edges: edges,
      pageInfo: {
        endCursor: endCursor,
        hasNextPage: hasNextPage
      }
    }
  };
}




async function addContactsImpl(_, { userId, contactsId }) {
  const userModel = await UserModel.findById(userId);
  if (userModel === null) {
    throw `collection<account> not find userId: ${userId}`;
  }

  let userContactsModel = await ContactsModel.findOne({ userId: userId, contacts: contactsId });
  if (userContactsModel === null) {
    console.log(`collection<${kContactsCollection}> not find ${userId}`);
    userContactsModel = new ContactsModel({ userId: userId, contacts: contactsId });
    userContactsModel.save();
    console.log(`collection<${kContactsCollection}> _id: ${userContactsModel._id}:${userId} add contacts ${contactsId} success`);
  } else {
    console.log(`collection<${kContactsCollection}> _id: ${userContactsModel._id}:${userId} aleady has contacts ${contactsId}`);
  }
}

async function addContacts(_, { userId, contactsId }) {
  addContactsImpl(_, { userId: userId, contactsId: contactsId });
  addContactsImpl(_, { userId: contactsId, contactsId: userId });
  return contactsId;
}

async function removeContactsImpl(_, { userId, contactsId }) {
  ContactsModel.deleteOne({ userId: userId, contacts: contactsId }).then(function () {
    console.log(`collection<${kContactsCollection}> ${userId} remove contacts ${contactsId} success`);
  }).catch(function (error) {
    console.log(error);
  });
}

async function removeContacts(_, { userId, contactsId }) {
  removeContactsImpl(_, { userId: userId, contactsId: contactsId });
  removeContactsImpl(_, { userId: contactsId, contactsId: userId });
  return contactsId;
}

export default resolvers;
