package sql



const UpsertAddContactsApply = `
upsert
into
	add_contacts_apply (
	user_id,
	contacts_id,
	message,
	update_time
)
values ($1,
$2,
$3,
$4)
`

const QueryAddContactsApply = `
select
	user_id,
	message,
	update_time
from
	add_contacts_apply@default_unique_index
where
	contacts_id = $1
	and update_time < $2
order by
	update_time desc
limit $3
`

const DeleteAddContactsApply = `
delete
from
	add_contacts_apply
where
	contacts_id = $1
	and user_id = $2
`
