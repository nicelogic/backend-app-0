package dependence

import (
	"context"
	"log"

	contactsConfig "auth/config"

	"github.com/nicelogic/authutil"
	"github.com/nicelogic/crdb"
)

func Init(serviceConfig *contactsConfig.Config) (*authutil.Auth, *crdb.Client, error) {
	authUtil := &authutil.Auth{}
	err := authUtil.Init(serviceConfig.Public_key_file_path,
		serviceConfig.Private_key_file_path)
	if err != nil {
		log.Printf("auth init err: %v", err)
		return nil, nil, err
	}
	crdbClient := &crdb.Client{}
	err = crdbClient.Init(context.Background(),
		serviceConfig.Db_config_file_path,
		serviceConfig.Db_name,
		serviceConfig.Db_pool_connections_num)
	if err != nil {
		log.Printf("crdb init err: %v", err)
		return authUtil, nil, err
	}
	return authUtil, crdbClient, err
}
