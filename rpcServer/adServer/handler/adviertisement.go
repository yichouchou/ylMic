package handler

import (
	"context"
	proto "ylMic/common/proto/ad"
	"ylMic/rpcServer/adServer/respository"
)

type Adviertisement struct {
	ad *respository.Adviertisement
}

func (g *Adviertisement) QueryAdviertisement(ctx context.Context, req *proto.QueryByExampleRequest, rsp *proto.QueryByExampleResponse) error {
	g.ad.QueryAdviertisementList(req, rsp)
	return nil
}
