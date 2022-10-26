#!/bin/bash

version=$1
#./build.sh $version
serviceName="user"
    
pathToEnv0="../../0-env"
kubeConfigFilePath=$(cat $pathToEnv0/which-env-to-apply)
kubeConfigFilePath=$pathToEnv0"/"$kubeConfigFilePath
echo "current env: $kubeConfigFilePath"
sed -i.bak "s/$serviceName:.*/$serviceName:${version}/" ./k8s/deployment.yml
kubectl --kubeconfig $kubeConfigFilePath apply -k ./k8s

