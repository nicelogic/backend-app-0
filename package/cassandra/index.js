
import fs from 'fs';
import Config from 'nicelogic-config';
import axios from 'axios';

class Cassandra {
	constructor() {
		this.cassandraUserName = fs.readFileSync('/etc/app-0/secret-cassandra/username', 'utf8').trim();
		this.cassandraPwd = fs.readFileSync('/etc/app-0/secret-cassandra/password', 'utf8').trim();
		const config = new Config('/etc/app-0/config/config.yml');
		this.cassandraAuthUrl = config.get('cassandra-auth-url', '');
		this.cassandraGraphqlUrl = config.get('cassandra-graphql-url', '');
	}

	async getToken() {
		const cassandraAdminNameAndPwd = { username: this.cassandraUserName, password: this.cassandraPwd };
		const json = JSON.stringify(cassandraAdminNameAndPwd);
		try {
			const response = await axios.post(this.cassandraAuthUrl,
				json,
				{
					headers: {
						'Content-Type': 'application/json'
					},
				}
			);
			const status = response.status;
			if(status !== 201){
				console.log(`auth url: ${this.cassandraAuthUrl}, username: ${this.username}, respone: ${status}`);
				return '';
			}
			const token = response.data.authToken;
			return token;
		} catch (e) {
			console.log(e);
			return '';
		}
	}
}

export default Cassandra;