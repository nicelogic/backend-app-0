package graph_test

import (
	"testing"
)

func TestAddContactsApply(t *testing.T){

	type testStruct struct {
		pStr *string
	}
	type topStruct struct {
		tStruct *testStruct
	}

	tStruct := &topStruct{}
	tStruct.tStruct = &testStruct{}
	tStruct.tStruct.pStr = nil


}