package config

import (
	"testing"

	"github.com/nicelogic/config"
)

func TestLoadUserConfig(t *testing.T){
	userConfig := Config{Path: "/", Listen_address: ":80"}
	config.Init("/etc/app-0/config-user/config-user.yml", &userConfig)

	path := "/user/gql"
	if userConfig.Path != path {
		t.Errorf("user config path: want: %s, but: %s\n", path, userConfig.Path)
	}

	listenAddress := "http://localhost:8080" 
	if userConfig.Listen_address != listenAddress {
		t.Errorf("user config path: want: %s, but: %s\n", listenAddress, userConfig.Listen_address)
	}
}