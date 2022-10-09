
import jwt from 'jsonwebtoken';
import { GraphQLError } from 'graphql';
import { errors } from 'nicelogic-error';

function getJwtPayload(req, publicKey) {
	let payload = {};
	const authorizationHeader = req.headers && 'Authorization' in req.headers ? 'Authorization' : 'authorization';
	if (!req.headers.hasOwnProperty(authorizationHeader)) {
		console.log('unauthorized');
		return {};
	}
	const token = (req.headers[authorizationHeader]).split(' ')[1];
	try {
		payload = jwt.verify(token, publicKey);
	} catch (e) {
		throw new GraphQLError(e.message);
	}
	return payload;
}

function getUserId(payload) {
	if (payload.hasOwnProperty('user')
		&& payload.user.hasOwnProperty('id')) {
		return payload.user.id;
	} else {
		throw new GraphQLError(errors.tokenInvalid);
	}
}

export {
	getJwtPayload,
	getUserId
}