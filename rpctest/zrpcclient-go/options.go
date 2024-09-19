package zrpcclient_go

import (
	"github.com/zeromicro/go-zero/zrpc"

	"rpctest/zrpcclient-go/typed/rpctest"
)

type Opt func(client *Clientset)


func WithRpctestClient(cli zrpc.Client) Opt {
	return func(client *Clientset) {
		client.rpctest = rpctest.New(cli)
	}
}


