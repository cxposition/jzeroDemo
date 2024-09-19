package zrpcclient_go

import (
	"github.com/zeromicro/go-zero/zrpc"

	"gateway/zrpcclient-go/typed/gateway"
)

type Opt func(client *Clientset)


func WithGatewayClient(cli zrpc.Client) Opt {
	return func(client *Clientset) {
		client.gateway = gateway.New(cli)
	}
}


