package dependence

import (
	"context"
	"log"

	contactsConfig "contacts/config"

	"github.com/nicelogic/authutil"
	"github.com/nicelogic/crdb"
	"github.com/nicelogic/pulsarclient"
)

func Init(serviceConfig *contactsConfig.Config)(*authutil.Auth, *crdb.Client, *pulsarclient.Client, error){
	authUtil := &authutil.Auth{}
	err := authUtil.Init(serviceConfig.Public_key_file_path,
		serviceConfig.Private_key_file_path)
	if err != nil {
		log.Printf("auth init err: %v", err)
		return nil, nil, nil, err
	}

	crdbClient := &crdb.Client{}
	err = crdbClient.Init(context.Background(),
		serviceConfig.Db_config_file_path,
		serviceConfig.Db_name,
		serviceConfig.Db_pool_connections_num)
	if err != nil {
		log.Printf("crdb init err: %v", err)
		return authUtil, nil, nil, err
	}
	pulsarClient := &pulsarclient.Client{}
	err = pulsarClient.Init(serviceConfig.Pulsar_config_file_path,
		serviceConfig.Pulsar_topic)
	if err != nil {
		log.Printf("pulsar init err: %v", err)
		return authUtil, crdbClient, nil, err
	}
	return authUtil, crdbClient, pulsarClient, err
}