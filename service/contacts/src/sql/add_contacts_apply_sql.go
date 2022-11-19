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
	add_contacts_apply
where
	contacts_id = $1
order by
	update_time desc
limit $2
`