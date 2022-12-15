

import os
import yaml

configYmlPath = r'./service.yml'
configYml = open(configYmlPath)
config = yaml.safe_load(configYml)
lastVersionKey = 'last_version'
lastVersion = config[lastVersionKey]
serviceName = config['service_name']
print('last version: ' + lastVersion)
print('service name: ' + serviceName)
lastVersion = 'v' + lastVersion
os.system('./deploy.sh ' + lastVersion + ' ' + serviceName)
