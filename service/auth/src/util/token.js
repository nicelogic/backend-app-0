
import jwt from 'jsonwebtoken';

function generateToken(id, key, expiresIn) {
	let token = jwt.sign({
		user: {
			id
		}
	}, key, { algorithm: 'RS256', expiresIn: expiresIn });
	return token;
}

export {
	generateToken
}