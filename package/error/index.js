
import { GraphQLError } from 'graphql';

const errors = {
	internalError: 'server internal error',
	tokenExpired: 'token expired',
	tokenInvalid: 'invalid token',
}

function reportError(e) {
	if (typeof (e) === "string") {
		throw new GraphQLError(e);
	} else if (e.hasOwnProperty('message')) {
		throw new GraphQLError(e.message);
	} else {
		throw new GraphQLError(errors.internalError);
	}
}

export {
	errors,
	reportError
}








