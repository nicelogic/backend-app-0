package dependence

import (
	"context"
	"log"

	userConfig "user/config"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/nicelogic/authutil"
	"github.com/nicelogic/config"
	"github.com/nicelogic/crdb"
)

func Init(serviceConfig *userConfig.Config) (authUtil *authutil.Auth, crdbClient *crdb.Client, minioClient *minio.Client, err error) {
	authUtil = &authutil.Auth{}
	err = authUtil.Init(serviceConfig.Public_key_file_path,
		serviceConfig.Private_key_file_path)
	if err != nil {
		log.Printf("auth init err(%v)", err)
		return
	}
	log.Printf("authutil init success")

	crdbClient = &crdb.Client{}
	err = crdbClient.Init(context.Background(),
		serviceConfig.Db_config_file_path,
		serviceConfig.Db_name,
		serviceConfig.Db_pool_connections_num)
	if err != nil {
		log.Printf("crdb init err(%v)", err)
		return
	}

	type MinioConfig struct {
		Endpoint          string
		Access_key_id     string
		Secret_access_key string
	}
	minioConfig := &MinioConfig{}
	err = config.Init(serviceConfig.S3_config_file_path, &minioConfig)
	if err != nil{
		log.Printf("minio config init err(%v), path(%v)\n", err, serviceConfig.S3_config_file_path)
		return
	}
	minioClient, err = minio.New(minioConfig.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(minioConfig.Access_key_id, minioConfig.Secret_access_key, ""),
		Secure: true,
	})
	if err != nil {
		log.Printf("minio init err(%v)\n", err)
		return
	}
	log.Printf("minio(%#v)\n", minioClient) // minioClient is now set up
	log.Printf("minio init success\n") // minioClient is now set up

	return
}
