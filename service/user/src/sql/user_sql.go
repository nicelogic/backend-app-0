package sql

const UpsertUser = `
insert
	into
	public."user"(id,
	data,
	update_time)
values ($1,
$2,
now())
on
conflict (id)
do
update
set
	data = public.user.data || $2,
	update_time = now()
returning id, data, name, update_time
`

const QueryMe = `
select
	id,
	"data",
	"name",
	update_time
from
	public."user"
where
	id = $1
`

const QueryNameOrId = `
select
	id,
	"data",
	"name",
	update_time
from
	public."user"
where
	(name = $1 or id = $1)
`