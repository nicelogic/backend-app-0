#!/bin/bash

version=$1
serviceName=$2

echo "inpurt aliyun docker reposigory password:"
docker login --username=niceice220 --password-stdin registry.cn-shanghai.aliyuncs.com
docker build -t logic-base/$serviceName:$version ./
docker tag logic-base/$serviceName:$version registry.cn-shanghai.aliyuncs.com/logic-base/$serviceName:$version
docker push registry.cn-shanghai.aliyuncs.com/logic-base/$serviceName:$version
    
