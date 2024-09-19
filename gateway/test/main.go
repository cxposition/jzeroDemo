package main

import (
	"context"
	"fmt"
	gateway "gateway/gateway-go"
	"gateway/gateway-go/model/gateway/pb/hellopb"
	"github.com/jzero-io/restc"
)

func main() {
	headers := make(map[string][]string)
	headers["Content-Type"] = []string{"application/json"}
	clientset, err := gateway.NewClientWithOptions(
		restc.WithAddr("127.0.0.1"),
		restc.WithPort("8001"),
		restc.WithProtocol("http"),
		restc.WithHeaders(headers),
	)
	if err != nil {
		panic(err)
	}

	result, err := clientset.Hello().SayHello(context.Background(), &hellopb.SayHelloRequest{Message: "hello"})
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}
