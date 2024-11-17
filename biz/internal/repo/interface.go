package repo

import (
	"github.com/li1553770945/sheepim-online-service/biz/infra/config"
	"github.com/li1553770945/sheepim-online-service/kitex_gen/online"
	"github.com/redis/go-redis/v9"
)

type IRepository interface {
	GetOnlineMemberEndpoint(clientList []string) (*online.GetOnlineMemberEndpointResp, error)
	SetClientEndpoint(clientId string, endpoint string) error
	RemoveClient(clientId string) error
}

type Repository struct {
	Cache              *redis.Client
	CacheExpireSeconds int64
}

func NewRepository(cache *redis.Client, cfg *config.Config) IRepository {
	return &Repository{
		Cache:              cache,
		CacheExpireSeconds: cfg.CacheConfig.ExpireSeconds,
	}
}
