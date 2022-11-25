#!/bin/bash

pathToEnv0="../0-env"
kubeConfigFilePath=$(cat $pathToEnv0/which-env-to-apply)
kubeConfigFilePath=$pathToEnv0"/"$kubeConfigFilePath
echo "current env: $kubeConfigFilePath"
kubectl --kubeconfig $kubeConfigFilePath create secret generic secret-crdb-ca-crt --from-file=ca.crt=./crdb/cert/ca.crt -n app-0
