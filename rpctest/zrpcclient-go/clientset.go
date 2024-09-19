package zrpcclient_go

import (
	"rpctest/zrpcclient-go/typed/rpctest"
)

type Interface interface {
	Rpctest() rpctest.Interface
}

type Clientset struct {
	rpctest *rpctest.Client
}

func (x *Clientset) Rpctest() rpctest.Interface {
	return x.rpctest
}


func NewClientWithOptions(ops ...Opt) *Clientset {
	cs := &Clientset{}

	for _, op := range ops {
		op(cs)
	}

	return cs
}
