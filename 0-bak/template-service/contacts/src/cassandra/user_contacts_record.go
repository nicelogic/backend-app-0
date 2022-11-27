package cassandra

const User_contacts_record = `

query user_contacts_record($user_id: String!, $contacts_id: String!){
	response: contacts(value:{
	  user_id: $user_id
	  contacts_id: $contacts_id
	}){
	 pageState
	 values {
	  user_id
	  contacts_id
	}
	}
  }


`