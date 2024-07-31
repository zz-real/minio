package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	File struct {
		Dir string
	}
	Minio struct {
		AccessKey       string
		SecretAccessKey string
		Endpoint        string
		UseSSL          bool
	}
}
