package sql

const QuerySameMembersChatWhetherExist = `
select
	user_id,
	chat_id,
	c.members
from
	public.user_chat u
join public.chat c 
on
	(u.chat_id = c.id)
where
	u.user_id = $1
	and c.members @> array[$2]
	and array_length(c.members, 1) = $3
`

const InsertChat = `

`

const QueryChats = `

`

