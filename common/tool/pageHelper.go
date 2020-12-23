package tool

import (
	"bytes"
	"fmt"
	"github.com/xormplus/xorm"
	"strconv"
	"ylMic/common/pojo"
)

type PageUtils struct {
	PageNum    int
	PageSize   int
	Client     *xorm.Engine
	SqlName    string
	SqlExample map[string]interface{}
	BeanList   interface{}
}

var PageUtil *PageUtils = nil

//func (*PageUtils) getPageResult() *pojo.PageInfo {
//
//}

func GetPageResult(pageUtil *PageUtils) *pojo.PageInfo {
	pageUtil.SqlExample["pageNum"] = pageUtil.PageNum
	pageUtil.SqlExample["pageSize"] = pageUtil.PageSize
	p := new(pojo.PageInfo)
	engine := pageUtil.Client
	//var bean []*interface{}
	//bean := make([]*interface{}, 0)

	//err := engine.SqlMapClient(pageUtil.SqlName, &pageUtil.SqlExample).Find(pageUtil.BeanList)
	//if err != nil {
	//	fmt.Println(err)
	//}
	sql := engine.GetSql(pageUtil.SqlName)

	s2 := "limit " + strconv.Itoa(pageUtil.PageNum) + "," + strconv.Itoa(pageUtil.PageSize)
	var bt bytes.Buffer
	bt.WriteString(sql)
	bt.WriteString(s2)
	s := bt.String()

	engine.SqlMap.Sql[pageUtil.SqlName] = s
	getSql := engine.GetSql(pageUtil.SqlName)
	fmt.Println(getSql)
	err := engine.SqlMapClient(pageUtil.SqlName, &pageUtil.SqlExample).Find(pageUtil.BeanList)
	if err != nil {
		fmt.Println(err)
	}

	//p.Result = adviertisements
	var total int
	err2 := engine.SqlMapClient("getTotal").Find(total)
	if err2 != nil {
		fmt.Println(err2)
	}
	p.PageTotal = uint32(total)
	p.PageNum = uint32(pageUtil.PageNum + 1)
	p.PageSize = uint32(pageUtil.PageSize)
	p.Result = pageUtil.BeanList
	return p
}
