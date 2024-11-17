// Code generated by Kitex v0.7.2. DO NOT EDIT.

package onlineservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	online "github.com/li1553770945/sheepim-online-service/kitex_gen/online"
)

func serviceInfo() *kitex.ServiceInfo {
	return onlineServiceServiceInfo
}

var onlineServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "OnlineService"
	handlerType := (*online.OnlineService)(nil)
	methods := map[string]kitex.MethodInfo{
		"SetClientStatus":         kitex.NewMethodInfo(setClientStatusHandler, newOnlineServiceSetClientStatusArgs, newOnlineServiceSetClientStatusResult, false),
		"GetOnlineMemberEndpoint": kitex.NewMethodInfo(getOnlineMemberEndpointHandler, newOnlineServiceGetOnlineMemberEndpointArgs, newOnlineServiceGetOnlineMemberEndpointResult, false),
	}
	extra := map[string]interface{}{
		"PackageName":     "online",
		"ServiceFilePath": `idl/online.thrift`,
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.7.2",
		Extra:           extra,
	}
	return svcInfo
}

func setClientStatusHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*online.OnlineServiceSetClientStatusArgs)
	realResult := result.(*online.OnlineServiceSetClientStatusResult)
	success, err := handler.(online.OnlineService).SetClientStatus(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newOnlineServiceSetClientStatusArgs() interface{} {
	return online.NewOnlineServiceSetClientStatusArgs()
}

func newOnlineServiceSetClientStatusResult() interface{} {
	return online.NewOnlineServiceSetClientStatusResult()
}

func getOnlineMemberEndpointHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*online.OnlineServiceGetOnlineMemberEndpointArgs)
	realResult := result.(*online.OnlineServiceGetOnlineMemberEndpointResult)
	success, err := handler.(online.OnlineService).GetOnlineMemberEndpoint(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newOnlineServiceGetOnlineMemberEndpointArgs() interface{} {
	return online.NewOnlineServiceGetOnlineMemberEndpointArgs()
}

func newOnlineServiceGetOnlineMemberEndpointResult() interface{} {
	return online.NewOnlineServiceGetOnlineMemberEndpointResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) SetClientStatus(ctx context.Context, req *online.SetClientStatusReq) (r *online.SetClientStatusResp, err error) {
	var _args online.OnlineServiceSetClientStatusArgs
	_args.Req = req
	var _result online.OnlineServiceSetClientStatusResult
	if err = p.c.Call(ctx, "SetClientStatus", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetOnlineMemberEndpoint(ctx context.Context, req *online.GetOnlineMemberEndpointReq) (r *online.GetOnlineMemberEndpointResp, err error) {
	var _args online.OnlineServiceGetOnlineMemberEndpointArgs
	_args.Req = req
	var _result online.OnlineServiceGetOnlineMemberEndpointResult
	if err = p.c.Call(ctx, "GetOnlineMemberEndpoint", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
