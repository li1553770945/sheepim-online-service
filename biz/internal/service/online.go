package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/li1553770945/sheepim-online-service/biz/constant"
	"github.com/li1553770945/sheepim-online-service/kitex_gen/base"
	"github.com/li1553770945/sheepim-online-service/kitex_gen/online"
)

func (s *OnlineService) SetClientStatus(ctx context.Context, req *online.SetClientStatusReq) (resp *online.SetClientStatusResp, err error) {
	klog.CtxInfof(ctx, "客户端%s状态:%v,%s ", req.ClientId, req.IsOnline, req.ServerEndpoint)
	if req.IsOnline {
		err = s.Repo.SetClientEndpoint(req.ClientId, req.ServerEndpoint)
		if err != nil {
			klog.CtxErrorf(ctx, "设置客户端 %s endpoint出错: %v", req.ClientId, err)
			return nil, err
		}
	} else {
		err = s.Repo.RemoveClient(req.ClientId)
		if err != nil {
			klog.CtxErrorf(ctx, "移除客户端 %s 出错", req.ClientId)
			return nil, err
		}
	}
	klog.CtxInfof(ctx, "设置客户端%s状态成功", req.ClientId)
	return &online.SetClientStatusResp{
		BaseResp: &base.BaseResp{
			Code: constant.Success,
		},
	}, nil
}

func (s *OnlineService) GetOnlineMemberEndpoint(ctx context.Context, req *online.GetOnlineMemberEndpointReq) (resp *online.GetOnlineMemberEndpointResp, err error) {
	klog.CtxInfof(ctx, "获取客户端列表状态，客户端ID列表: %v", req.ClientIdList)

	// 调用 repo 层方法获取客户端在线信息
	onlineMembersResp, err := s.Repo.GetOnlineMemberEndpoint(req.ClientIdList)
	if err != nil {
		klog.CtxErrorf(ctx, "获取客户端列表状态出错: %v", err)
		return nil, err
	}

	// 组装响应数据
	resp = &online.GetOnlineMemberEndpointResp{
		BaseResp: &base.BaseResp{
			Code:    constant.Success,
			Message: "获取成功",
		},
		Status: onlineMembersResp.Status,
	}

	klog.CtxInfof(ctx, "获取客户端列表状态成功, 共%v条数据", len(resp.Status))
	return resp, nil
}
