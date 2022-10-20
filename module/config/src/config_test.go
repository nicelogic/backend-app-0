package config

import "testing"

func TestConfigerInit(t *testing.T) {

	cases := struct{
		Path string
	}{""}

	err := Init("/etc/app-0/config-user/config-user.yml", &cases)
	if err != nil{
		t.Error("has error: ", err)
	}
	if cases.Path != "/user/gql"{
		t.Error("path want /user/gql, but ", cases.Path)
	}
}
