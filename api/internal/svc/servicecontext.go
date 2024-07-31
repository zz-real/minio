package svc

import (
	"context"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"minio/api/internal/config"
	"path/filepath"
	"time"
)

type ServiceContext struct {
	Config config.Config
	Mc     *minio.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	minioClient, err := minio.New(c.Minio.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(c.Minio.AccessKey, c.Minio.SecretAccessKey, ""),
		Secure: c.Minio.UseSSL,
	})
	if err != nil {
		panic(err)
	}
	//watch(c)
	svc := ServiceContext{
		Config: c,
		Mc:     minioClient,
	}
	go svc.watch()

	return &svc
}

func (c *ServiceContext) upload(fileName string) {
	_, name := filepath.Split(fileName)
	b, err := c.Mc.FPutObject(context.Background(), "test", name, fileName, minio.PutObjectOptions{ContentType: "application/text"})
	fmt.Println(b, err)
}

// 监听文件改变
func (c *ServiceContext) watch() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		panic(err)
	}
	defer watcher.Close()

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Has(fsnotify.Create) {
					fmt.Println("modified file:", event.Name)
					time.Sleep(time.Second * 2)
					go c.upload(event.Name)
				}
			case err := <-watcher.Errors:
				fmt.Errorf("watcher.Errors:", err)
			}
		}
	}()

	err = watcher.Add(c.Config.File.Dir)
	if err != nil {
		panic(err)
	}

	<-make(chan bool)
}
