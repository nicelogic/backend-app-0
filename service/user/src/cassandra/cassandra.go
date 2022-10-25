package cassandra

import (
	"errors"
	"fmt"
	"strings"
	"time"
	"user/graph/model"

	"github.com/mitchellh/mapstructure"
)

func UpdateUserGql(changes map[string]interface{}) (gql string, variables map[string]interface{}, err error) {

	updatedUser := &model.User{}
	var metadata mapstructure.Metadata
	config := &mapstructure.DecoderConfig{
		Metadata: &metadata,
		Result:   &updatedUser,
	}
	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return
	}
	err = decoder.Decode(changes)
	if err != nil {
		return
	}

	const gqlFormat = `mutation updateuser(
					$id: String!
					# VariableDefinitions
					%s
					$update_time: Timestamp!
				) {
				updateuser: updateuser(
					value: {
							id: $id
							# variables
							%s
						    update_time: $update_time
					},
				ifExists: false
				) {
					applied
					accepted
					value {
						id
						# response values
						%s
						update_time
					}
			    }
	  		}`
	variables = map[string]interface{}{
		"update_time": time.Now().Format(time.RFC3339),
	}
	variableDefinitions := ""
	arguments := ""
	responseValues := ""
	for _, userKey := range metadata.Keys {
		lowerUserKey := strings.ToLower(userKey)
		variables[lowerUserKey] = changes[lowerUserKey]
		fmt.Printf("update user(%s:%v)\n", userKey, changes[lowerUserKey])

		variableDefinitions += fmt.Sprintf("$%s: String\n", lowerUserKey)
		arguments += fmt.Sprintf("%s: $%s\n", lowerUserKey, lowerUserKey)
		responseValues += fmt.Sprintf("%s\n", lowerUserKey)
	}
	gql = fmt.Sprintf(gqlFormat, variableDefinitions, arguments, responseValues)

	return gql, variables, nil
}

func GetUpdatedUserFromResponse(response map[string]interface{}) (user *model.User, err error){

	response = response["updateuser"].(map[string]interface{})
	applied := response["applied"].(bool)
	if !applied {
		err = errors.New("cassandra not applied")
		return
	}
	value := response["value"].(map[string]interface{})
	user = &model.User{}
	err = mapstructure.Decode(value, &user)
	if err != nil {
		return
	}
	return
}

const QueryUserByIdGql = `query queryuser($id: String!) {
	queryuser: user(value: {
					  id:$id, 
					},
					) {
			pageState,
			values {
			  id,
			  name,
			  signature,
			  update_time
			}
		  }
  }`

  func GetUserFromQueryUserByIdResponse(response map[string]interface{}) (user *model.User, err error){

	response = response["queryuser"].(map[string]interface{})
	values := response["values"].([]interface{})
	if len(values) == 0 {
		return
	}
	if len(values) != 1 {
		err = errors.New("response[queryuser][values] length != 1")
		return
	}
	value := values[0].(map[string]interface{})
	user = &model.User{}
	err = mapstructure.Decode(value, &user)
	if err != nil {
		return
	}
	return
}

const QueryUserByNameGql = `query queryuserbyname($name: String!, $pageState: String) {
	queryuserbyname: user(filter: {
				  name: {eq: $name}
					},
							  options: {
					pageSize: 1,
					pageState: $pageState
				  }
					) {
			pageState,
			values {
				id,
			  name,
			  signature,
			  update_time
			}
		  }
  }`

  func GetUserFromQueryUserByNameResponse(response map[string]interface{}) (users map[string]*model.User, pageState string, err error){

	response = response["queryuserbyname"].(map[string]interface{})
	pageState = response["pageState"].(string)
	values := response["values"].([]interface{})

	users = make(map[string]*model.User)
	for _, value := range values {
		value = value.(map[string]interface{})
		user := &model.User{}
		err = mapstructure.Decode(value, &user)
		if err != nil {
			fmt.Printf("error: %v\n", err)
		} else {
			users[user.ID] = user
		}
	}
	return
}
