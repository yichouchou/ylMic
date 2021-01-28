package pojo

//分页器结构
type PageInfo struct {
	Result    interface{}
	PageSize  uint32 //每页大小
	PageTotal uint32 //总页数
	PageNum   uint32 //当前页数
}

var PageSize = 10 //默认页大小

//获取默认页大小
func GetPageSize() uint32 {
	return uint32(PageSize)
}

//设置默认页大小
func SetPageSize(ps uint32) {
	if ps < 1 {
		ps = 1
	}
	PageSize = int(ps)
}
