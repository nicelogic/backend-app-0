package cassandra

const UpdateAddContactsApplyGql = `
mutation updateadd_contacts_apply(
	$user_id: String!
	$contacts_id: String!
	$id: String!
	$update_time: Timestamp!
	$remark_name: String!
  	$message: String
  ){
	response: updateadd_contacts_apply(value:{
	  user_id: $user_id
	  contacts_id: $contacts_id
	  id: $id
	  update_time: $update_time
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
		id
		user_id
		remark_name
		message
	  }
	}
  }
`


/*
{
	"user_id": "1",
	"contacts_id": "2",
	"id": "1>2",
	"update_time": "2022-10-30T11:30:25.000Z",
	"remark_name": "2",
	"message": "please add me"
}
*/
