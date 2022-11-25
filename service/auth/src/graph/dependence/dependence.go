package dependence

import (
	"context"
	"log"

	contactsConfig "auth/config"

	"github.com/nicelogic/crdb"
)

func Init(serviceConfig *contactsConfig.Config)(*crdb.Client, error){
	crdbClient := &crdb.Client{}
	err := crdbClient.Init(context.Background(),
		serviceConfig.Db_config_file_path,
		serviceConfig.Db_name,
		serviceConfig.Db_pool_connections_num)
	if err != nil {
		log.Printf("crdb init err: %v", err)
		return nil, err
	}
	// pulsarClient := &pulsarclient.Client{}
	// err = pulsarClient.Init(serviceConfig.Pulsar_config_file_path,
	// 	serviceConfig.Pulsar_topic)
	// if err != nil {
	// 	log.Printf("pulsar init err: %v", err)
	// 	return crdbClient, nil, err
	// }
	return crdbClient, err
}