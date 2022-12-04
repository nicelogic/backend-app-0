
CREATE DATABASE message;
CREATE TYPE chat_type AS ENUM ('p2p', 'group');
CREATE TABLE public.chat(
	id STRING NOT NULL,
	type CHAT_TYPE NOT NULL,
	members STRING[] NOT NULL,
	name STRING, 
	last_message JSONB,
	last_message_time TIMESTAMPTZ,
	--last_message_update_time STRING AS (last_message->>'update_time') STORED,
	update_time TIMESTAMPTZ NOT NULL DEFAULT now():::TIMESTAMPTZ,
	CONSTRAINT "primary" PRIMARY KEY (id ASC)
	--INVERTED INDEX (members)
);


CREATE TABLE public.user_chat(
	user_id STRING NOT NULL,
	chat_id STRING NOT NULL references public.chat(id) ON DELETE CASCADE,
	priority INT DEFAULT 0, --default: 0, pinned: 1
	-- last_message_time TIMESTAMPTZ,
	update_time TIMESTAMPTZ NOT NULL DEFAULT now():::TIMESTAMPTZ,
	CONSTRAINT "primary" PRIMARY KEY (user_id ASC, chat_id ASC),
	-- UNIQUE INDEX (user_id ASC, priority DESC, last_message_time DESC, chat_id ASC),
	UNIQUE INDEX (user_id ASC, priority DESC, chat_id ASC),
	INDEX (chat_id ASC)
);

CREATE TABLE public.message(
	id STRING NOT NULL,
	chat_id STRING NOT NULL,
	content JSONB NOT NULL,
	sender_Id STRING NOT NULL,
	create_time TIMESTAMPTZ NOT NULL DEFAULT now():::TIMESTAMPTZ,
	CONSTRAINT "primary" PRIMARY KEY (id ASC),
	UNIQUE INDEX (chat_id ASC, create_time DESC, id ASC) STORING(content, sender_Id)
);


/*
这种设计，每次chat发送一条消息，会更新每个members,排序相关,需要更新时间
创建chat事务创建
查询需要join
个人的配置不影响主体chat
主体chat信息更改也不影响个人配置
没办法在一个chat表里面遍历更改所有的user setting(JSONB方式)
影响用户对chat特定状态+影响排序的一个表， chat通用信息一个表
还要一个方案，每次排序user_chat,都全部查询user所有chat,再排序
但每次查询都全量查找
还要就是未读信息数量，也是每个用户特定，又和每条消息关联的
每次有新消息，就得更新用户特定状态。。顺带也更新时间
最终选择last_message_time在user_chat维护的方案
获取可以在chat表里做last_message_time index来支持Join排序
先不做复杂。join获取chat last_message_time再做排序
last_message_time是必须的。。paginatin必须要求所有字段在一个表里
有时间思考几个问题:
	* 是否chat表members数组可以去掉
	* user chat的pagination last message time是否可以去掉，依赖chat last message time
	* 如果第二点做不到。则怎么做到cascade update user chat table last message time

每次写一条消息都要更新member的信息  vs 查询的时候 先把user_chat join chat 再排序
选择后者
因为发消息是很频繁的事情。如果group内人很多。。一次都要更新几百个人的表记录（如果真要更新也是要批处理)
而之所以要记录last_message_time 只是为了排序
我对这块的理解没有系统了解和深入的做过实验。但是目前来看，只更新chat表，代码上更简单些。
易于理解。发一条消息，只新增消息和更新chat表即可。不会涉及批量操作
但是这种设计，导致user pagination chat 的性能不是最高的
但是全局上来看, 无论读写都是分布式的部分用户数据
对这块需要深入研究和做实验才能得出科学的结论。。目前不了解的情况下，选择简单的方案
不对pagination做过早优化
至少可以得到明显的数据简单的处理逻辑的好处。性能的好处，仅凭粗浅理解，没有实验数据做支撑。不可靠

*/