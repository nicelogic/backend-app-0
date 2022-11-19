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