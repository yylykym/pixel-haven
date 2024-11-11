package config

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
)

var Conf *Config

type Config struct {
	minioClient *minio.Client
}

func Init() *Config {
	Conf = &Config{}
	client, err := minio.New("localhost:9000", &minio.Options{
		Creds:  credentials.NewStaticV4("minioadmin", "minioadmin", ""),
		Secure: false,
	})
	if err != nil {
		log.Fatalf("Failed to create MinIO client: %s", err)
	}
	Conf.minioClient = client
	return &Config{}
}

// MinioClient Client 获取 MinIO 客户端
func (config *Config) MinioClient() *minio.Client {
	return config.minioClient
}
