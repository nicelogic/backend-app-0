
# skeleton

## 总体思想

* env项目负责物理机器构建k8s集群，及其监控k8s集群状态相关的服务： dnsutil, dashboard, promethues, 负责k8s集群的稳定
* skeleton里面的东西是所有app共享的。所有的app可以基于这个骨架去部署，运维。也可以独立。
* 以app为核心，每个服务就一个namespace.
  * 优点：mongo namespace和service同一个namespace, service可以使用Mongo的secret, 后端服务间可以共享
* traefik是所有app共享的, traefik和k8s强绑定。为所有的app转发流量。当然也可以每个app独享。
* 唯一网址luojm.com.所有app共享。appname.luojm.com

