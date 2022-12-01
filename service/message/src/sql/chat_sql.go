package sql

const QuerySameMembersP2pChatWhetherExist = `
select
	id,
	members, 
	name,
	last_message 
from public.chat
where
	members @> $1
	and array_length(members, 1) = $2
	and type = 'p2p'
`

const InsertChat = `

`

const QueryChats = `

`

