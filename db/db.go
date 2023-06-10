package db

import (
	"context"
	"io"
	"os"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DBClient struct {
	Client *gorm.DB
	Minio  *minio.Client
}

func NewDBClient() (*DBClient, error) {

	minioClient, err := minio.New("minio.sabbir.dev", &minio.Options{
		Creds:  credentials.NewStaticV4("", "", ""),
		Secure: true,
	})
	if err != nil {
		log.Error(err)
	}

	obj, err := minioClient.GetObject(context.TODO(), "remembrall", "task.sqlite", minio.GetObjectOptions{})
	if err != nil {
		log.Error("could not fetch db: ", err)
	}

	localFile, err := os.Create("task.db")
	if err != nil {
		log.Error("could not create file: ", err)
		return &DBClient{
			Client: &gorm.DB{},
			Minio:  &minio.Client{},
		}, err
	}
	defer localFile.Close()

	if _, err := io.Copy(localFile, obj); err != nil {
		log.Error("could not write db: ", err)
		return &DBClient{
			Client: &gorm.DB{},
			Minio:  &minio.Client{},
		}, err
	}

	db, err := gorm.Open(sqlite.Open("task.db"), &gorm.Config{})
	if err != nil {
		log.Error("could not open db", err)
		return &DBClient{
			Client: &gorm.DB{},
			Minio:  &minio.Client{},
		}, err
	}

	return &DBClient{
		Client: db,
		Minio:  minioClient,
	}, nil
}
