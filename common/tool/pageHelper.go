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

func GetPageResult(pageUtil *PageUtils) *pojo.PageInfo {
	pageUtil.SqlExample["pageNum"] = pageUtil.PageNum
	pageUtil.SqlExample["pageSize"] = pageUtil.PageSize
	p := new(pojo.PageInfo)
	engine := pageUtil.Client
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
	_, err2 := engine.SqlMapClient("getTotal").Get(&p.PageTotal)
	if err2 != nil {
		fmt.Println(err2)
	}
	p.PageNum = uint32(pageUtil.PageNum + 1)
	p.PageSize = uint32(pageUtil.PageSize)
	p.Result = pageUtil.BeanList
	return p
}
