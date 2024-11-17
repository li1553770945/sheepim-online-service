namespace go online
include "base.thrift"

struct SetClientStatusReq{
    1: required string clientId
    2: required string serverEndpoint
    3: required bool isOnline
}
struct SetClientStatusResp{
    1: required base.BaseResp baseResp
}

struct GetOnlineMemberEndpointReq{
    1: required list<string> clientIdList
}

struct ClientStatusData{
    1: required string clientId
    2: required string serverEndpoint
    3: required bool isOnline
}
struct GetOnlineMemberEndpointResp{
    1: required base.BaseResp baseResp
    2: optional list<ClientStatusData> status
}


service OnlineService {
    SetClientStatusResp SetClientStatus(SetClientStatusReq req)
   GetOnlineMemberEndpointResp GetOnlineMemberEndpoint(GetOnlineMemberEndpointReq req)
}
