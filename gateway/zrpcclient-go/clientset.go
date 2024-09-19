package zrpcclient_go

import (
	"gateway/zrpcclient-go/typed/gateway"
)

type Interface interface {
	Gateway() gateway.Interface
}

type Clientset struct {
	gateway *gateway.Client
}

func (x *Clientset) Gateway() gateway.Interface {
	return x.gateway
}


func NewClientWithOptions(ops ...Opt) *Clientset {
	cs := &Clientset{}

	for _, op := range ops {
		op(cs)
	}

	return cs
}
