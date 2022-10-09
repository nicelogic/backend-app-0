#!/bin/bash

    
# pathToEnv0="../../0-env"
# kubeConfigFilePath=$(cat $pathToEnv0/which-env-to-apply)
# kubeConfigFilePath=$pathToEnv0"/"$kubeConfigFilePath
# echo "current env: $kubeConfigFilePath"
# kubectl --kubeconfig $kubeConfigFilePath apply -k ./k8s

version=1.0.4

./build.sh $version

sed -i '/logic-base\/auth/ s/auth:*/auth:$version/' ./k8s/deployment.yml