#!/bin/bash

version=0.0.1
service=user

docker login --username=niceice220 --password-stdin registry.cn-shanghai.aliyuncs.com
docker build -t logic-base/$service:$version ./
docker tag logic-base/$service:$version registry.cn-shanghai.aliyuncs.com/logic-base/$service:$version
docker push registry.cn-shanghai.aliyuncs.com/logic-base/$service:$version
    
