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

	const gqlFormat = `mutation updateUser(
					$id: String!
					# VariableDefinitions
					%s
					$update_time: Timestamp!
				) {
				updateUser: updateuser(
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
	response, ok := response["updateUser"].(map[string]interface{})
	if !ok {
		err = errors.New("response[updateUser]'s type is not map[string]interface{}")
		return
	}
	applied, ok := response["applied"].(bool)
	if !ok {
		err = errors.New("response[updateUser][applied]'s type is not bool")
		return
	}
	if !applied {
		err = errors.New("cassandra not applied")
		return
	}
	value, ok := response["value"].(map[string]interface{})
	if !ok {
		err = errors.New("response[updateUser][value]'s type is not map[string]interface{}")
		return
	}
	user = &model.User{}
	err = mapstructure.Decode(value, &user)
	if err != nil {
		return
	}

	return
}

const QueryUserByIdGql = `query queryuser($id: String!) {
	queryUser: user(value: {
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

  func GetQueryUserFromResponse(response map[string]interface{}) (user *model.User, err error){
	response, ok := response["queryUser"].(map[string]interface{})
	if !ok {
		err = errors.New("response[queryUser]'s type is not map[string]interface{}")
		return
	}

	values, ok := response["values"].([]interface{})
	if !ok {
		err = errors.New("response[queryUser][values]'s type is not []interface{}")
		return
	}
	if len(values) != 1 {
		err = errors.New("response[queryUser][values] length != 1")
		return
	}
	value, ok := values[0].(map[string]interface{})
	if !ok {
		err = errors.New("response[queryUser][values][0]'s type is not map[string]interface{}")
		return
	}

	user = &model.User{}
	err = mapstructure.Decode(value, &user)
	if err != nil {
		return
	}
	return
}