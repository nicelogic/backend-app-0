package sql

const UpsertContacts = `
upsert
into
	contacts  
	(user_id,
		contacts_id,
		remark_name,
		update_time
	)
values ($1,
	$2,
	$3,
	$4)
`

const QueryUserAddedMe = `
select
	user_id,
	contacts_id
from
	contacts c
where
	user_id = $1
	and contacts_id = $2
`

const QueryContacts = `
select
	contacts_id ,
	remark_name ,
from
	contacts@default_unique_index
where 
	user_id = $1
	and (remark_name, contacts_id ) > ($2, $3)
order by
	remark_name, contacts_id 
limit $4
`

const DeleteContacts = `
delete
from
	public.contacts
where
	user_id = $1
	and contacts_id = $2
`