
# design

## design

* 用go语言
* 用graphql
* 业务基于k8s又不感知k8s(比如读取配置文件是基于k8s configmap)


## faq

### 为什么不使用多种语言并存

写的脚本还是用的python/shell
业务语言只用go
之前没学过go的就去学习
业务语言首先排除nodejs, python等动态语言
也不需要lua为了支持静态语言热更新的语言
都用过，没有静态类型，看代码首先就心累，还容易出错
一切设计的原则，为了尽可能地简单
c/c++/rust，系统性的语言
java和c++差不多
go， dart google系的，提供了通用问题的最佳解决方案和这些解决方案的工具
所以，坚定不移地走：go, dart（flutter)的道路