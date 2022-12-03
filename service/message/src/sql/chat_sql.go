package sql

const QueryChat = `
select
	id,
	type,
	members, 
	name,
	last_message,
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
now());
`

const QueryChats = `

`

