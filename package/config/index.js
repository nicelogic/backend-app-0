
import yaml from 'js-yaml';
import fs from 'fs';

class Config {
	constructor(configFilePath) { 
		try {
			this.doc = yaml.load(fs.readFileSync(configFilePath, 'utf8'));
			console.log(`${configFilePath}: `);
			console.log(this.doc)
		} catch (e) {
			console.log(e);
		}
	}

	get(key, defaltVal) {
		try {
			const val = this.doc[key];
			console.log(`${key}: ${val}`);
			return val;
		} catch (e) {
			console.log(e);
		}
		return 	defaltVal;
	}
}

export default Config;