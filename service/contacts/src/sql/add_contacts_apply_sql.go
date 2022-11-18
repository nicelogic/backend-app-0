package sql



const UpsertAddContactsApply = `

UPSERT INTO add_contacts_apply (
	user_id,
	contacts_id,
	message,
	update_time
)
VALUES ($1, $2, $3, $4)

`

const AddContactsApply = `

SELECT user_id,
	message,
	update_time
from add_contacts_apply
where contacts_id = $1
ORDER BY update_time DESC

`