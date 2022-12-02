package sql

const QuerySameMembersP2pChatWhetherExist = `
select
	id,
	type,
	members, 
	name,
	last_message 
from public.chat
where id = $1
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

const QueryChats = `

`

