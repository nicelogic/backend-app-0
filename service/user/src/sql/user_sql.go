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
returning * 
`

const QueryAuth = `
select
	auth_id,
	auth_id_type,
	user_id,
	auth_id_type_username_pwd
from
	public.auth
where
	auth_id = $1
`
