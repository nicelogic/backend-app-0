const mongoose = require('mongoose');

const initDb = function initDb({dbName}){
	let mongoHost = process.env.MONGO_SERVICE_HOST;
	if (typeof mongoHost === 'undefined') {
		mongoHost = '127.0.0.1';
	}
	const connectUrl = `mongodb://${mongoHost}:27017/${dbName}`;
	console.log(connectUrl);
	const mongooseConnect = () => {
		mongoose.connect(connectUrl, {
			useNewUrlParser: true,
			useUnifiedTopology: true,
			auto_reconnect: true 
		});
	};

	try {
		mongooseConnect();
	} catch (e) {
		console.log('mongodb connect error.');
		setTimeout(mongooseConnect, 1000);
	}

	const db = mongoose.connection;
	db.on('error', console.error.bind(console, 'connection error:'));
	db.once('open', function () {
		console.log('mongodb connected');
	});
	db.on('disconnected', function () {
		console.log('mongodB disconnected');
		mongooseConnect();
	});
}

const AccountModel = mongoose.model('account', {
	_id: String,
	name: String,
	info: String,
	contact: [
		{
			id: String
		}
	],
	chats: [{
		type: mongoose.Schema.Types.ObjectId,
		ref: 'chat'
	}],
	messages: [{
		type: mongoose.Schema.Types.ObjectId,
		ref: 'message'
	}]
});


const ChatModel = mongoose.model('chat', {
	name: String,
	lastMessage: String,
	updated: String,
	members: [{
		type: mongoose.Schema.Types.ObjectId,
		ref: 'account'
	}],
	messages: [{
		type: mongoose.Schema.Types.ObjectId,
		ref: 'message'
	}]
});

const MessageModel = mongoose.model('message', {
	_id: String,
	text: String,
	data: String,
	chats: [{
		type: mongoose.Schema.Types.ObjectId,
		ref: 'chat'
	}],
});

module.exports = {
	initDb,
	AccountModel,
	ChatModel, 
	MessageModel
}

