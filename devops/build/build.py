
import os
import yaml
import sys

#configYmlPath = r'./service.yml'
dockerFileDir = sys.argv[1]
configYmlPath = sys.argv[1] + '/service.yml'
print('configYmlPath: ', configYmlPath)
configYml = open(configYmlPath)
config = yaml.safe_load(configYml)
lastVersionKey = 'last_version'
lastVersion = config[lastVersionKey]
serviceName = config['service_name']
print('last version: ' + lastVersion)
v1, v2, v3 = str(lastVersion).split('.')
v3 = str(int(v3) + 1)
noVprefixBuildVersion = v1 + '.' + v2 + '.' + v3
buildVersion = 'v' + noVprefixBuildVersion
print('build version: ' + buildVersion)
print('service name: ' + serviceName)
config[lastVersionKey] = noVprefixBuildVersion

os.system('./build.sh ' + buildVersion + ' ' + serviceName + ' ' + dockerFileDir)
with open(configYmlPath, 'w') as updateYmlPath:
	yaml.dump(config, updateYmlPath)
