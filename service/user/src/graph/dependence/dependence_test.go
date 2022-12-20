package dependence_test

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func TestMinio(t *testing.T) {
	endpoint := "tenant0.minio.env0.luojm.com:9443"
	accessKeyID := "bUIacujwb7dMIALn"
	secretAccessKey := "oKG9Y1B2QBsDW6fmlyJYVARcVJJaqqs1"
	useSSL := true
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Printf("minio init err(%v)\n", err)
	}
	log.Printf("minio(%#v)\n", minioClient) // minioClient is now set up
	presignedURL, err := minioClient.PresignedPutObject(context.Background(), "app-0", "/users/test/avatar.png", time.Duration(60) * time.Minute)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(presignedURL)

	t.Log("success")
}
