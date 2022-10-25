
import fs from 'fs';
import Config from 'nicelogic-config';
import axios from 'axios';
import { GraphQLClient } from 'graphql-request'

class Cassandra {
	constructor() {
		this.cassandraUserName = fs.readFileSync('/etc/app-0/secret-cassandra/username', 'utf8').trim();
		this.cassandraPwd = fs.readFileSync('/etc/app-0/secret-cassandra/password', 'utf8').trim();
		const config = new Config('/etc/app-0/config/config.yml');
		this.cassandraAuthUrl = config.get('cassandra-auth-url', '');
		this.cassandraGraphqlUrl = config.get('cassandra-graphql-url-app_0', '');
	}

	async getToken() {
		const cassandraAdminNameAndPwd = { username: this.cassandraUserName, password: this.cassandraPwd };
		const response = await axios.post(this.cassandraAuthUrl,
			JSON.stringify(cassandraAdminNameAndPwd),
			{
				headers: {
					'Content-Type': 'application/json'
				},
			}
		);
		const token = response.data.authToken;
		return token;
	}

	async mutation(gql, variables){
		const token = await this.getToken();
		console.log(`token: ${token}`);

		const graphQLClient = new GraphQLClient(this.cassandraGraphqlUrl,
			{
			  method: 'POST',
			  jsonSerializer: {
			    parse: JSON.parse,
			    stringify: JSON.stringify,
			  },
			  headers: {
			    'x-cassandra-token': token,
			  },
			});
		const response = await graphQLClient.request(gql, variables);
		return response;
	}

	async query(gql, variables){
		const token = await this.getToken();
		console.log(`token: ${token}`);

		const graphQLClient = new GraphQLClient(this.cassandraGraphqlUrl,
			{
			  method: 'GET',
			  jsonSerializer: {
			    parse: JSON.parse,
			    stringify: JSON.stringify,
			  },
			  headers: {
			    'x-cassandra-token': token,
			  },
			});
		const response = await graphQLClient.request(gql, variables);
		return response;
	}
}

export default Cassandra;