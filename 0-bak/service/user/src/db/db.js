import mongoose from 'mongoose';

const initDb = function initDb({ dbName }) {
	let mongoHost = process.env.MONGO_SERVICE_HOST;
	if (typeof mongoHost === 'undefined') {
		mongoHost = '127.0.0.1';
	}
	const connectUrl = `mongodb://${mongoHost}:27017/${dbName}`;
	console.log(connectUrl);
	const mongooseConnect = () => {
		mongoose.connect(connectUrl, {
			useNewUrlParser: true,
			useUnifiedTopology: true
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

const userSchema = mongoose.Schema({
	_id: String,
	name: String
}, {
	toJSON: { virtuals: true },
	toObject: { virtuals: true }
});
userSchema.virtual('contacts', {
	ref: 'contacts',
	localField: '_id',
	foreignField: 'userId'
});
const UserModel = mongoose.model('account', userSchema);

const contactsSchema = mongoose.Schema({
	userId: String,
	contacts: {
		type: mongoose.Schema.Types.String,
		ref: 'account'
	}
}, { timestamps: true });
const ContactsModel = mongoose.model('contacts', contactsSchema);

export {
	initDb,
	UserModel,
	ContactsModel,
};

