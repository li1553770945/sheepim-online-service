// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package container

import (
	"github.com/li1553770945/sheepim-online-service/biz/infra/cache"
	"github.com/li1553770945/sheepim-online-service/biz/infra/config"
	"github.com/li1553770945/sheepim-online-service/biz/infra/log"
	"github.com/li1553770945/sheepim-online-service/biz/infra/trace"
	"github.com/li1553770945/sheepim-online-service/biz/internal/repo"
	"github.com/li1553770945/sheepim-online-service/biz/internal/service"
)

// Injectors from wire.go:

func GetContainer(env string) *Container {
	configConfig := config.GetConfig(env)
	traceLogger := log.InitLog()
	traceStruct := trace.InitTrace(configConfig)
	client := cache.NewCache(configConfig)
	iRepository := repo.NewRepository(client, configConfig)
	iOnlineService := service.NewOnlineService(iRepository)
	container := NewContainer(configConfig, traceLogger, traceStruct, iOnlineService)
	return container
}
