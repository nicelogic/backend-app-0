
import yaml from 'js-yaml';
import fs from 'fs';

class Config {
	constructor(configFilePath) { 
		try {
			this.doc = yaml.load(fs.readFileSync(configFilePath, 'utf8'));
		} catch (e) {
			console.log(e);
		}
	}

	get(key, defaltVal) {
		let val;
		try {
			val = this.doc[key];
			if(val === undefined){
				val = defaltVal;
				console.log('use default value');
			}
		} catch (e) {
			console.log(e);
			console.log('use default value');
			val = defaltVal;
		}
		console.log(`${key}: ${val}`);
		return 	val;
	}
}

export default Config;