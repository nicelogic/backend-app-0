package dependence

import (
	"context"
	"log"

	messageConfig "message/config"

	"github.com/nicelogic/authutil"
	"github.com/nicelogic/crdb"
	"github.com/nicelogic/pulsarclient"
)

func Init(serviceConfig *messageConfig.Config) (authUtil *authutil.Auth, crdbClient *crdb.Client, pulsarClient *pulsarclient.Client, err error) {
	authUtil = &authutil.Auth{}
	err = authUtil.Init(serviceConfig.Public_key_file_path,
		serviceConfig.Private_key_file_path)
	if err != nil {
		log.Printf("auth init err: %v", err)
		return
	}

	crdbClient = &crdb.Client{}
	err = crdbClient.Init(context.Background(),
		serviceConfig.Db_config_file_path,
		serviceConfig.Db_name,
		serviceConfig.Db_pool_connections_num)
	if err != nil {
		log.Printf("crdb init err: %v", err)
		return
	}
	pulsarClient = &pulsarclient.Client{}
	err = pulsarClient.Init(serviceConfig.Pulsar_config_file_path,
		serviceConfig.Pulsar_topic)
	if err != nil {
		log.Printf("pulsar init err: %v", err)
		return
	}
	return
}
