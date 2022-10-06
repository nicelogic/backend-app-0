
import jwt from 'jsonwebtoken';

function generateToken(id, key) {
	let created = Math.floor(Date.now() / 1000);
	let token = jwt.sign({
		exp: created + 60 * 60 * 24, //1 day
		iat: created,
		id: id
	}, key, { algorithm: 'RS256' });
	return token;
}

export {
	generateToken
}