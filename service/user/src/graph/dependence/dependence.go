package dependence

import (
	"context"
	"log"

	userConfig "user/config"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/nicelogic/authutil"
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
	crdbClient = &crdb.Client{}
	err = crdbClient.Init(context.Background(),
		serviceConfig.Db_config_file_path,
		serviceConfig.Db_name,
		serviceConfig.Db_pool_connections_num)
	if err != nil {
		log.Printf("crdb init err(%v)", err)
		return
	}

	endpoint := "https://tenant0.minio.env0.luojm.com:9443"
	accessKeyID := "bUIacujwb7dMIALn"
	secretAccessKey := "oKG9Y1B2QBsDW6fmlyJYVARcVJJaqqs1"
	useSSL := true
	minioClient, err = minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Printf("minio init err(%v)\n", err)
	}
	log.Printf("minio(%#v)\n", minioClient) // minioClient is now set up
	return
}
