import mongoose from 'mongoose';

const authSchema = mongoose.Schema({
	userNamePwdAuth: {
		userName: String,
		pwd: String
	}
}, { timestamps: true });
const AuthModel = mongoose.model('auth', authSchema);

export {
	AuthModel,
};

