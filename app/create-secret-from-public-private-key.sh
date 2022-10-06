#!/bin/bash

pathToEnv0="../0-env"
kubeConfigFilePath=$(cat $pathToEnv0/which-env-to-apply)
kubeConfigFilePath=$pathToEnv0"/"$kubeConfigFilePath
echo "current env: $kubeConfigFilePath"
kubectl --kubeconfig $kubeConfigFilePath create secret generic secret-jwt --from-file=jwt-privatekey=./cert/2_niceice.cn.key --from-file=jwt-publickey=./cert/1_niceice.cn_bundle.crt -n app-0
