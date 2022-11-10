package cassandra

const UpdateAddContactsApplyGql = `

mutation  add_contacts_apply (
	$user_id: String!
	$contacts_id: String!
	$update_time: Timestamp!
	$remark_name: String!
    $message: String!
    $add_contacts_apply_ttl: Int = 604800 #one week: 60 * 60 * 24 * 7
) @atomic {
    
	updatecontacts_by_remark_name: updatecontacts_by_remark_name(
		value: {
      user_id: $user_id,
      contacts_id: $contacts_id,
      remark_name: $remark_name,
      update_time: $update_time
    }
		ifExists: false
		){
      applied
      accepted
      value{
        user_id
        contacts_id
        remark_name
        update_time
      }
  	}
  
  updatecontacts: updatecontacts(
		value: {
      user_id: $user_id,
      contacts_id: $contacts_id,
      update_time: $update_time
    }
		ifExists: false
		){
      applied
      accepted
      value{
        user_id
        contacts_id
        update_time
      }
  	}
  
  	updateadd_contacts_apply: updateadd_contacts_apply(value:{
	  user_id: $user_id
	  contacts_id: $contacts_id
	  update_time: $update_time
    message: $message
	}
	  ifExists: false
    options: {
      ttl: $add_contacts_apply_ttl
    }
	){
	  applied
	  accepted
	  value {
		contacts_id
		update_time
		user_id
		message
	  }
	}
	
  }


`


/*

{
  "user_id": "1",
  "contacts_id": "2",
  "update_time":"2022-11-05T11:38:45.000Z",
  "remark_name": "2",
  "message": "please add me",
  "add_contacts_apply_ttl": 604800
}

*/
