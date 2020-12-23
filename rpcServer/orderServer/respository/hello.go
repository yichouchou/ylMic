package respository

import (
	"github.com/go-xorm/xorm"
	"ylMic/rpcServer/orderServer/model"
)

type Responsitory interface {
	Create(order *model.Order)
	Find(orderId string) (*model.Order, error)
	Update(model.Order, int64) (model.Order, error)
}

func Create(order *model.Order, client *xorm.Engine) error {
	//这里填写数据库操作
	_, err := client.Insert(order)
	return err
}
