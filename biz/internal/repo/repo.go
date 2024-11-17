package repo

import (
	"context"
	"github.com/li1553770945/sheepim-online-service/kitex_gen/online"
)

func (r *Repository) GetOnlineMemberEndpoint(clientList []string) (*online.GetOnlineMemberEndpointResp, error) {

	ctx := context.Background()

	var onlineClients online.GetOnlineMemberEndpointResp
	for _, clientId := range clientList {
		// 检查客户端是否在线
		onlineKey := "online:" + clientId
		exist, err := r.Cache.Exists(ctx, onlineKey).Result()
		if err != nil {
			return nil, err
		}
		if exist == 0 {
			continue
		}
		endPoint, err := r.Cache.Get(ctx, onlineKey).Result()
		if err != nil {
			return nil, err
		}
		onlineClients.Status = append(onlineClients.Status, &online.ClientStatusData{
			ClientId:       clientId,
			ServerEndpoint: endPoint,
			IsOnline:       true,
		})
	}

	return &onlineClients, nil
}
func (r *Repository) SetClientEndpoint(clientId string, endpoint string) error {
	ctx := context.Background()

	// Redis key 格式
	onlineKey := "online:" + clientId

	// 将客户端的 endpoint 存储到 Redis
	_, err := r.Cache.Set(ctx, onlineKey, endpoint, 0).Result() // 不设置过期时间，可以根据需求设置

	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) RemoveClient(clientId string) error {
	ctx := context.Background()

	// Redis key 格式
	onlineKey := "online:" + clientId

	// 删除 Redis 中的客户端信息
	_, err := r.Cache.Del(ctx, onlineKey).Result()

	if err != nil {
		return err
	}

	return nil
}
