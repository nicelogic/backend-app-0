
import yaml from 'js-yaml';

export class Config {
	getMongoDbConfig() {
		let dbUrl;
		try {
			//'base-mongodb-svc'
			const isDebug = process.env.NODE_ENV === 'debug';
			const configMongodbPath = isDebug ? '../../../config/configmap/example/config/config-mongodb.yml'
			: '/etc/config/mongodb/config-mongodb.yml';
			const doc = yaml.load(fs.readFileSync(configMongodbPath, 'utf8'));
			dbUrl = doc['url'];
			console.log(`dbUrl: ${dbUrl}`);
			
			const dbUserNamePath = isDebug ? '../../../config/configmap/example/secret/name' : '/etc/secret/mongodb/name';
			const dbPwdPath = isDebug ? '../../../config/configmap/example/secret/password' : '/etc/secret/mongodb/password';
			const dbUserName = fs.readFileSync(dbUserNamePath, 'utf8').trim();
			const dbPwd = fs.readFileSync(dbPwdPath, 'utf8').trim();
			console.log(`dbUserName: ${dbUserName}, dbPwd: ${dbPwd}`);

			dbUrl = dbUrl.replace('name', dbUserName).replace('pwd', dbPwd);
			console.log(`final dbUrl: ${dbUrl}`);
		} catch (e) {
			console.log(e);
		}
		return {
			dbUrl,
		};
	}
}