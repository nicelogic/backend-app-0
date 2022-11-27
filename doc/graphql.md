
# graphql


## 如何做到服务器只resolve部分field

比如客户端只需要获取自己的缩略信息，但是服务器已经有了一个获取用户全部信息的resolver
那么是否还需要提供一个获取缩略信息的resolver呢
如果需要的话，那工作量就大了
而且因为客户端可以随时只需要部分的查询结果，那么难道就要写多个resolver去适配了
绝不可能！

所以就有了如下概念： 

* Field Collection

//和fragment概念应该不同。fragment是为了解决 type中部分field 复用问题
//和directives概念也不同，其是为了解决客户端一个相同请求使用参数获取不同返回值
//field collection是为了解决客户端一个相同请求，对于不必要的返回值，不执行不必要返回值相关操作的问题

## go中query/mutation 是否同时运行在两个不同线程中

模拟mutation的时候，获取token, 然后sleep一段时间
query，更新token, mutation再去获取token.看是否更新


感觉go的设计挺好的： 状态始终是最重要的。 cassandra module里面，只有Client struct是状态
而其中只有token是多个线程可能会更新到的。其他都是初始化之后就不变的。所以可能就不需要做线程之间的保护