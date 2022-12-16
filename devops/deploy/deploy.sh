#!/bin/bash

version=$1
serviceName=$2
serviceDir=$3
    
pathToEnv0="../../0-env"
kubeConfigFilePath=$(cat $pathToEnv0/which-env-to-apply)
kubeConfigFilePath=$pathToEnv0"/"$kubeConfigFilePath
echo "current env: $kubeConfigFilePath"
sed -i.bak "s/$serviceName:.*/$serviceName:${version}/" $serviceDir/k8s/deployment.yml
kubectl --kubeconfig $kubeConfigFilePath apply -k $serviceDir/k8s

