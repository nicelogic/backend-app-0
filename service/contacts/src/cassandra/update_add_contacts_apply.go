package cassandra

const UpdateAddContactsApplyGql = `
mutation updateadd_contacts_apply(
	$contacts_id: String!
	$update_time: Timestamp!
	$user_id: String!
	$remark_name: String!
  $message: String
  ){
	response: updateadd_contacts_apply(value:{
	  contacts_id: $contacts_id
	  update_time: $update_time
	  user_id: $user_id
	  remark_name: $remark_name
    message: $message
	}
	  ifExists: false
	){
	  applied
	  accepted
	  value {
		contacts_id
		update_time
		user_id
		remark_name
		message
	  }
	}
  }
`
