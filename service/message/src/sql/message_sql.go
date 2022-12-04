package sql

const InsertMessage = `
insert
	into
	public.message
(id,
	chat_id,
	content,
	sender_id,
	create_time)
values($1,
$2,
$3,
$4,
now())
`

const UpdateChatLastMessage = `
update
	public.chat
set
	last_message = $1,
	last_message_time = $2,
	update_time = now()
where
	id = $3
`

const UpdateUserChatLastMessage = `
update
	public.user_chat
set
	last_message_time = $1,
	update_time = now()
where
	user_id = $2
	and chat_id = $3
`

const QueryMessages = `
select
	id,
	content,
	sender_id,
	create_time
from
	public.message
where
	chat_id = $1
	and (create_time < $2
		or (create_time = $2
			and id > $3))
order by 
	create_time desc,
	id asc
limit $4
`