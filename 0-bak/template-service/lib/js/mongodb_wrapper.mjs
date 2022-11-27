
import mongoose from 'mongoose';

export class MongoDbWrapper {
	init({ dbUrl }) {
		const mongooseConnect = () => {
			mongoose.connect(dbUrl, {
				useNewUrlParser: true,
				useUnifiedTopology: true
			});
		};

		try {
			mongooseConnect();
		} catch (e) {
			console.log(`mongodb(${dbUrl}) connect error, try reconnect`);
			setTimeout(mongooseConnect, 3000);
		}

		const db = mongoose.connection;
		db.on('error', console.error.bind(console, `mongodb(${dbUrl}) connection error:`));
		db.once('open', function () {
			console.log(`mongodb(${dbUrl}) connected`);
		});
		db.on('disconnected', function () {
			console.log(`mongodB(${dbUrl}) disconnected, try reconnect`);
			mongooseConnect();
		});
	}
}
