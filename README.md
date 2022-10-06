
# backend app-0

## app-0

english name: warmth
中文: 抱团取暖

## 设计

* use microservice
* client每次只获取其需要的信息去渲染显示（可以多次获取)
* client缓存获取到的信息，有变化才去重新获取
* microservice之间可以存在互相依赖关系: contacts的返回可以依赖user
* 每个微服务的返回值为： error_code, error_code_description，“具体对象"

## design

### 为什么不app直接访问数据库

* 和数据库选型解耦合,数据库后端可以变更
* 微服务理念，每个功能一个服务，一个抽象。便于服务进行业务控制
  要不然业务变更全靠客户端强制升级
* 鉴权，校验等不是数据库的事情

## faq

### 为什么不需要apollo route/federation

* 因为它的好处在于有个地方把所有微服务的graphql集合在一块
  但是我不需要。每个微服务一个ingress route, 简单明了
  主要是鉴权等这类操作，有没有啥办法能过统一去做(可以放在traefik去做)
  把鉴权，数据库连接等相同的操作封装起来，最小化化调用
	这样就简单明了就好,不要引入其他7788的东西增加复杂度
  subgraph在micro-service上搞了一层。
  一个请求到subgraph,可以组合多个service的数据然后返回给client
  一个全的大的东西不是好的设计
  分而治之才好
  没必要把所有的不相关的东西都整合在一起
  如果contacts microservice need user microservice data, just contacts to get user info to 
  send response
  没必要一定强求一次页面上所有的数据都一次性全部返回。只要不是冗余的他们所需要的带宽都是一样的
  虽然会耗费多一点点的性能。但是相较于复杂度等等，根本微不足道
			
### 如何一次graphql请求只请求需要渲染的数据，而不是全部数据,客户端和服务端的model要怎么组织

* fragment

是要为每个页面写fragment吗？
这样数据就和view比较耦合在一块了。好处在于仅仅需要渲染的数据
但是这种情况库要怎么写

getPreviewData
getDetailData
每个页面的数据其实都有一个抽象，然后通过库接口去获取

服务端对于这种情况，是否只需要写一个全量的数据

### 如何缓存上次已经请求的数据，如果其没变化的话，就不返回数据，通过cache数据直接去渲染

* cache 

这个就和服务器没什么关系

### 数组数据要怎么判断是否增加成员，减少成员或者更新成员而不是每次都是去全量拿数据

* pagination

是否也适用于缓存机制。如果没变不管。如果变了，去获取差异的地方：增加，减少，更新

### 是否有必要实时监听所有修改，实时同步渲染。还是仅仅部分诸如：新消息/新好友等重要的几个消息才走订阅通知

* subscription

这种也主要针对多设备多终端情况
而且要结合cache情况看