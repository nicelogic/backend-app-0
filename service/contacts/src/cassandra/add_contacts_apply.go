
package cassandra

const AddContactsApplyGql = `
query add_contacts_apply($user_id: String!){
	response: add_contacts_apply(value:{
	  contacts_id: $user_id
	}){
	  pageState
	  values {
		id
		  user_id
		contacts_id
		remark_name
		message
		update_time
	  }
	}
  }
`