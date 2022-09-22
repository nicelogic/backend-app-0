
const mongoose = require('mongoose');

exports.initDb = function initDb({dbName, collectionName}){
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
	return new mongoose.model(collectionName, {
		_id: String,
		accountId: String,
		targetId: String,
		event: String,
		info: String,
		state: String,
		replyEvent: String,
		replyInfo: String
	});
}

// export default initDb;