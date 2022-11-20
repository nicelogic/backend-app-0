package graph_test

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func TestAddContactsApply(t *testing.T) {

	data := "aaa"
	sEnc := base64.StdEncoding.EncodeToString([]byte(data))
    fmt.Println(sEnc)

	sDec, _ := base64.StdEncoding.DecodeString(sEnc)
    fmt.Println(string(sDec))
    fmt.Println()

}
