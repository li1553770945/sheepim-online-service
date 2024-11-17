package service

import (
	"context"
	"github.com/li1553770945/sheepim-online-service/biz/internal/repo"
	"github.com/li1553770945/sheepim-online-service/kitex_gen/online"
)

type OnlineService struct {
	Repo repo.IRepository
}

type IOnlineService interface {
	SetClientStatus(ctx context.Context, req *online.SetClientStatusReq) (resp *online.SetClientStatusResp, err error)
	GetOnlineMemberEndpoint(ctx context.Context, req *online.GetOnlineMemberEndpointReq) (resp *online.GetOnlineMemberEndpointResp, err error)
}

func NewOnlineService(repo repo.IRepository) IOnlineService {
	return &OnlineService{
		Repo: repo,
	}
}
