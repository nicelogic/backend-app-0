#!/bin/bash

# docker login --username=niceice220 --password-stdin registry.cn-shanghai.aliyuncs.com
# docker build -t logic-base/auth ./
# docker tag logic-base/auth registry.cn-shanghai.aliyuncs.com/logic-base/auth
# docker push registry.cn-shanghai.aliyuncs.com/logic-base/auth
    
pathToEnv0="../../0-env"
kubeConfigFilePath=$(cat $pathToEnv0/which-env-to-apply)
kubeConfigFilePath=$pathToEnv0"/"$kubeConfigFilePath
echo "current env: $kubeConfigFilePath"
kubectl --kubeconfig $kubeConfigFilePath apply -k ./k8s
