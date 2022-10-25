
# data

## 总体思想

* auth的数据， user(profile, config, contacts)的数据， chat的数据，post的数据
contacts服务到底能不能拆分出来
毕竟contacts还包括请求加好友，同意等等，必须拆分
可以contacts里增加添加时候的name,然后如果有备注，则修改其name为备注（showname)
这样也方便数据库查询。每次点进去看详情去拉取真实用户名的时候，可能要返回修改(最终一致性)
