
# contacts

## data 

add_contacts_req 表
存储 a 向 b 申请添加通讯录 请求

* b需要获取到所有向其申请的请求 ==> contacts_id需要做index
	然后b 同意/拒绝 同意则，同时往contacts表里加个记录 b 有联系人 a
	然后再将同意设置到 add_contacts_req 中的 是否同意字段中去
	这么做的好处在于，如果设置是否同意失败，则依然有这个记录在
	会触发b再去设置一遍
	而如果先修改add_contacts_req表。则已经设置过了。后面如果contacts添加记录失败
	则没有机会再去触发添加记录
	如果涉及到两次数据库修改的，都要遵循这个原则（同时设置数据库的操作应该要是幂等的
* b设置好答复之后，
	也可以在增加好友的时候，两边同时来增加
	这个记录可以设置成ttl，一定时间后自动删除


## 实时通知

* 关于实时通知