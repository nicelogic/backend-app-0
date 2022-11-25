package sql

const InsertAuth = `
insert
	into
	auth  
	(auth_id,
		auth_id_type,
		auth_id_type_username_pwd ,
		user_id  
	)
values (
	$1,
	$2,
	$3,
	$4
	)
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
