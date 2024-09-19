package main

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/zrpc"
	rpctest "rpctest/zrpcclient-go"
	"rpctest/zrpcclient-go/model/rpctest/pb/hellopb"
)

func main() {
	target, err := zrpc.NewClientWithTarget("localhost:8000")
	if err != nil {
		panic(err)
	}
	var cs rpctest.Interface
	cs = rpctest.NewClientWithOptions(rpctest.WithRpctestClient(target))

	hello, err := cs.Rpctest().Hello().SayHello(context.Background(), &hellopb.SayHelloRequest{Message: "hello"})
	if err != nil {
		panic(err)
	}

	fmt.Println(*hello)
}
