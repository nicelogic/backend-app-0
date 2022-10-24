package cassandra

import (
	"fmt"
	"testing"
)

func TestUpdateUserGql(t *testing.T) {
	changes :=  map[string]interface{}{
		"name": "haha",
	}

	gql, variables, err := UpdateUserGql(changes)
	if err != nil {
		t.Errorf("error: %v", err)
	}

	fmt.Println(gql)
	fmt.Printf("%v\n", variables)
}