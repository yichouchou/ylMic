package respository

import (
	"fmt"
	"github.com/xormplus/xorm"
	"ylMic/common/pojo"
	proto "ylMic/common/proto/ad"
	"ylMic/common/tool"
	"ylMic/rpcServer/adServer/model"
)

type Responsitory interface {
	QueryAdviertisement(adviertisement *model.Adviertisement, pageInfo *pojo.PageInfo) (pageResult *pojo.Result)
	QueryAdviertisementList(req *proto.QueryByExampleRequest) (pageResult *proto.QueryByExampleResponse)
}

type Adviertisement struct {
}

func (a *Adviertisement) QueryAdviertisement(adviertisement *model.Adviertisement, pageInfo *pojo.PageInfo) (pageResult *pojo.Result) {
	m := make([]model.Adviertisement, 0)
	err := tool.GetDbClient().SQL("select * from adviertisement limit 20").Find(&m)
	if err != nil {
		return nil
	}
	p := new(pojo.Result)
	p.Data = m
	p.Code = 200
	p.Message = "成功"
	return p
}

func (a *Adviertisement) QueryAdviertisementList(req *proto.QueryByExampleRequest, pageResult *proto.QueryByExampleResponse) {
	ad := &proto.Adviertisement{
		Id:    (req.Adviertisement.Id),
		Name:  req.Adviertisement.Name,
		Title: req.Adviertisement.Title,
		Tel:   req.Adviertisement.Tel,
		Type:  (req.Adviertisement.Type),
	}
	p := &proto.PageInfo{
		PageSize: req.PageInfo.PageSize,
		PageNum:  req.PageInfo.PageNum,
	}
	fmt.Println(p)
	fmt.Println(ad)
	adviertisements := make([]*(proto.Adviertisement), 0)
	client := tool.GetDbClient()
	err := client.RegisterSqlMap(xorm.Xml("./sqlMapper", "adSqlMapper.xml"))
	if err != nil {
		fmt.Println(err)
	}

	selectADByPages := "selectADByPages2"
	paramMap_4_3 := map[string]interface{}{"title1": "测试标题", "title2": "测试标题2"}
	t := new(tool.PageUtils)
	t.PageSize = int(req.PageInfo.PageSize)
	t.PageNum = 0
	t.Client = client
	t.SqlName = selectADByPages
	t.SqlExample = paramMap_4_3
	t.BeanList = &adviertisements

	//1.首先创建一个可供查询语句接收的结构体： 比如 	adviertisements := make([]*(proto.Adviertisement), 0)
	//2.然后创建 查询条件封装的 map ：比如 paramMap_4_3 := map[string]interface{}{"title1": "测试标题", "title2": "测试标题2"}
	//3.创建分页相关参数，封装进结构体 包含 pagenum pagesize
	//4.准备client 数据库连接
	//5.准备sqlname字符串，对应xml需要执行的sql

	result := tool.GetPageResult(t)
	fmt.Println(result)
	//results1, err := client.SqlMapClient(selectADByPages, &paramMap_4_3).Query().Json()
	//json, err := client.SqlMapClient("getTotal").Query().Json()
	//fmt.Println(json)
	//
	//sql := client.GetSql("selectADByPages")
	//fmt.Println(sql)
	//
	//fmt.Println(results1)

	//var adviertisements []*ad.Adviertisement
	//err := tool.GetDbClient().SQL("select * from adviertisement ").Limit(int(p.PageSize), int(p.PageNum)).Find(&adviertisements)

	if err != nil {
		fmt.Println(err)
	}
	pageResult.Adviertisement = adviertisements
	//pageResult.Adviertisement = result.Result.([]*proto.Adviertisement)  这一行和上面一行是一致的，adviertisements已经实例化了

	pageResult.PageInfo = &proto.PageInfo{
		PageNum:  result.PageNum,
		PageSize: result.PageSize,
		Total:    result.PageTotal,
	}
	/*m := make([]model.Adviertisement, 0)
	err := tool.GetDbClient().SQL("select * from adviertisement limit 20").Find(&m)
	if err != nil {
		return nil
	}
	p := new(pojo.Result)
	p.Data = m
	p.Code = 200
	p.Message = "成功"
	return p*/
}
