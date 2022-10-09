#!/bin/bash

version=1.0.4
./build.sh $version
    
pathToEnv0="../../0-env"
kubeConfigFilePath=$(cat $pathToEnv0/which-env-to-apply)
kubeConfigFilePath=$pathToEnv0"/"$kubeConfigFilePath
echo "current env: $kubeConfigFilePath"
sed -i.bak "s/auth:.*/auth:${version}/" ./k8s/deployment.yml
kubectl --kubeconfig $kubeConfigFilePath apply -k ./k8s

