package tool

import (
	"bytes"
	"fmt"
	"strconv"
	"ylMic/common/pojo"
)

type PageUtils struct {
	PageNum    int
	PageSize   int
	SqlName    string
	SqlExample map[string]interface{}
	BeanList   interface{}
}

func GetPageResult(pageUtil *PageUtils) *pojo.PageInfo {
	pageUtil.SqlExample["pageNum"] = pageUtil.PageNum
	pageUtil.SqlExample["pageSize"] = pageUtil.PageSize
	p := new(pojo.PageInfo)
	client := GetDbClient()
	//从sqlmap获取初始sql语句
	sql := client.GetSql(pageUtil.SqlName)
	sqlCacheMap := pojo.GetSqlCacheMap()
	//拼接 limit
	s2 := "limit " + strconv.Itoa(pageUtil.PageNum) + "," + strconv.Itoa(pageUtil.PageSize)
	var bt bytes.Buffer
	bt.WriteString(sql)
	bt.WriteString(s2)
	s := bt.String()
	client.SqlMap.Sql[pageUtil.SqlName] = s
	getSql := client.GetSql(pageUtil.SqlName)
	fmt.Println(getSql)
	err := client.SqlMapClient(pageUtil.SqlName, &pageUtil.SqlExample).Find(pageUtil.BeanList)
	if err != nil {
		fmt.Println(err)
	}
	_, err2 := client.SqlMapClient("getTotal").Get(&p.PageTotal)
	if err2 != nil {
		fmt.Println(err2)
	}
	p.PageNum = uint32(pageUtil.PageNum + 1)
	p.PageSize = uint32(pageUtil.PageSize)
	p.Result = pageUtil.BeanList
	sqlCacheMap[pageUtil.SqlName] = pageUtil
	return p
}
