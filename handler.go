package main

import (
	"context"
	"github.com/li1553770945/sheepim-online-service/biz/infra/container"
	online "github.com/li1553770945/sheepim-online-service/kitex_gen/online"
)

// OnlineServiceImpl implements the last service interface defined in the IDL.
type OnlineServiceImpl struct{}

// SetClientStatus implements the OnlineServiceImpl interface.
func (s *OnlineServiceImpl) SetClientStatus(ctx context.Context, req *online.SetClientStatusReq) (resp *online.SetClientStatusResp, err error) {
	App := container.GetGlobalContainer()
	resp, err = App.OnlineService.SetClientStatus(ctx, req)
	return
}

func (s *OnlineServiceImpl) GetOnlineMemberEndpoint(ctx context.Context, req *online.GetOnlineMemberEndpointReq) (resp *online.GetOnlineMemberEndpointResp, err error) {
	App := container.GetGlobalContainer()
	resp, err = App.OnlineService.GetOnlineMemberEndpoint(ctx, req)
	return
}
