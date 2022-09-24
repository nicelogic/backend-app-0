
# warmth backend

## use basic-component

* cassandra

## service

* auth

## design

### 为什么不app直接访问数据库

* 和数据库选型解耦合,数据库后端可以变更
* 微服务理念，每个功能一个服务，一个抽象。便于服务进行业务控制
  要不然业务变更全靠客户端强制升级
* 鉴权，校验等不是数据库的事情