package cassandra

import "testing"

func TestCassandraInit(t *testing.T){
	aCassandra := Cassandra{}
	err := aCassandra.Init("app_0_user")
	if err != nil{
		t.Errorf("error: %s", err)
		return
	}
	rightAuthUrl := "https://auth.cassandra.env0.luojm.com:9443/v1/auth"
	if aCassandra.authUrl != rightAuthUrl{
		t.Errorf("authUrl is wrong, want: %s, but: %s", rightAuthUrl, aCassandra.authUrl)
	}
}

func TestGetToken(t *testing.T){
	aCassandra := Cassandra{}
	err := aCassandra.Init("app_0_user")
	if err != nil{
		t.Errorf("error: %s", err)
		return
	}

	aCassandra.getToken()
}