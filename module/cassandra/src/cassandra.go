package cassandra

import (
	"fmt"
	"os"
	"strings"

	"github.com/nicelogic/config"
)

type Cassandra struct {
	userName   string
	pwd        string
	authUrl    string
	graphqlUrl string
	token string
}

type cassandraConfig struct {
	Cassandra_auth_url    string
	Cassandra_graphql_url string
}

func (cassandra *Cassandra) Init(keyspace string) (err error) {
	byteUserName, err := os.ReadFile("/etc/app-0/secret-cassandra/username")
	if err != nil {
		return
	}
	cassandra.userName = strings.TrimSpace(string(byteUserName))
	fmt.Printf("userName: %s\n", cassandra.userName)

	bytePwd, err := os.ReadFile("/etc/app-0/secret-cassandra/password")
	if err != nil {
		return
	}
	cassandra.pwd = strings.TrimSpace(string(bytePwd))
	fmt.Printf("pwd: %s\n", cassandra.pwd)

	aConfig := cassandraConfig{}
	err = config.Init("/etc/app-0/config/config.yml", &aConfig)
	if err != nil {
		return
	}
	cassandra.authUrl = aConfig.Cassandra_auth_url
	cassandra.graphqlUrl = aConfig.Cassandra_graphql_url + keyspace
	fmt.Printf("authUrl: %s\n", cassandra.authUrl)
	fmt.Printf("authGraphqlUrl: %s\n", cassandra.graphqlUrl)

	return
}

func (cassandra *Cassandra) getToken() string {
	return ""
}

func (cassandra *Cassandra) 

