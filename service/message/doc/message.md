
# message

## 关于p2p,群聊

有区别。
创建members == 2 chat, member相同，则复用旧的chat
创建members > 2 chat, 始终创建新的chat

## create chat

* 一个聊天里面有许多member
* member是否要显示这个chat
  * chat里面至少有一条消息
    * 创建人创建的时候，会发一条消息
* 创建聊天的时候， members == 2, 检测是否已经有相同成员的chat



## chat db生命周期

如果一个chat，last_message 半年内没任何消息，则删除该chat
但是后续又有人发消息呢？
删除的chat里有各个成员的配置信息,这些信息永久丢失了
chat不断增长.其实也类似contacts, user info. 也是创建了之后就有的
group chat应该有个解散机制, 在老的chat发消息会得到明确提示。重新创建新chat

