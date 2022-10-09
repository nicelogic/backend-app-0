#!/bin/bash

version=$1

echo "inpurt aliyun docker reposigory password:"
docker login --username=niceice220 --password-stdin registry.cn-shanghai.aliyuncs.com
docker build -t logic-base/auth:$version ./
docker tag logic-base/auth:$version registry.cn-shanghai.aliyuncs.com/logic-base/auth:$version
docker push registry.cn-shanghai.aliyuncs.com/logic-base/auth:$version
    
