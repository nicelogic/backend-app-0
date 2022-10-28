package cassandra

const UpdateAddContactsApplyGql = `
mutation updateadd_contacts_apply(
	$id: String!
	$user_id: String!
	$contacts_id: String!
	$remark_name: String!
	$update_time: Timestamp!
  ){
	response: updateadd_contacts_apply(value:{
	  id:$id
	  user_id: $user_id
	  contacts_id: $contacts_id
	  remark_name: $remark_name
	  update_time: $update_time
	}
	  ifExists: false
	){
	  applied
	  accepted
	  value {
		id
		user_id
		contacts_id
		remark_name
		update_time
	  }
	}
  }
`
