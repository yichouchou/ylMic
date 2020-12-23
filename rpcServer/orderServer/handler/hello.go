package handler

import (
	"context"
	proto "ylMic/common/proto/greeter"
	"ylMic/common/tool"
	"ylMic/rpcServer/orderServer/model"
	"ylMic/rpcServer/orderServer/respository"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
	m := new(model.Order)
	m.Total = 10
	m.Id = "1"
	m.OrderTime = "test"
	m.Uid = 1
	err := respository.Create(m, tool.GetDbClient())
	rsp.Greeting = "Hello " + req.Name
	if err != nil {
		return err
	}
	return nil
}
