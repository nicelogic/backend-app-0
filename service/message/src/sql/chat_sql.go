package sql

const QueryUserChat = `
select
	id,
	type,
	members, 
	name,
	last_message,
	last_message_time,
	uc."priority" 
from public.chat c
join public.user_chat uc 
on (c.id = uc.chat_id)
where uc.user_id = $1 and c.id = $2
`

const InsertChat = `
insert
	into
	public.chat
(id,
	"type",
	members,
	"name",
	update_time)
values($1,
$2,
$3,
$4,
now())
returning id, type, members, name, last_message
`
const InsertUserChat = `
insert
	into
	public.user_chat
(user_id,
	chat_id,
	"priority",
	update_time)
values($1,
$2,
0,
now())
`

const DeleteChat = `
delete
from
	public.chat
where
	id = $1
`

const QueryChats = `
select
	c.id,
	c.type,
	c.members, 
	c.name,
	c.last_message,
	c.last_message_time,
	uc."priority"
from
	public.user_chat uc
join 
	public.chat c
on
	(uc.chat_id = c.id)
where
	uc.user_id = $1
	and (
		(uc."priority" = $2 and (c.last_message_time < $3 or uc.chat_id > $4))
		or (uc."priority" < $2) 
	)
order by
	uc."priority" desc,
	c.last_message_time desc,
	uc.chat_id asc
limit $5
`

