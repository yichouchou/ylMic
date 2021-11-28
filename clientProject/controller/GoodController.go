package controller

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"ylMic/common/pojo"
)

type GoodController struct {
}

func (gc *GoodController) Router(app *gin.Engine) {
	app.GET("/api/foods", gc.getGoods)
}

//获取某个商户下面所包含的食品
func (gc *GoodController) getGoods(context *gin.Context) {
	shopId, exist := context.GetQuery("shop_id")
	if !exist {
		pojo.Failed(context, "请求参数错误，请重试")
		return
	}

	//实例化一个goodService,并调用对应的service方法
	id, err := strconv.Atoi(shopId)
	if err != nil {
		pojo.Failed(context, "请求参数错误，请重试")
		return
	}

	goodService := service.NewGoodService()
	goods := goodService.GetFoods(int64(id))
	if len(goods) == 0 {
		pojo.Failed(context, "未查询到相关数据")
		return
	}
	//查询到商户中的食品数据
	pojo.Success(context, goods)
}
