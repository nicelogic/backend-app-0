package cassandra

import (
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
					$user_id: String!
					# VariableDefinitions
					%s
					$update_time: Timestamp!
				) {
				updateUser: updateuser(
					value: {
							user_id: $user_id
							# variables
							%s
						    update_time: $update_time
					},
				ifExists: false
				) {
					applied
					accepted
					value {
						user_id
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
