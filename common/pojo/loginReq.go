package pojo

type LoginReq struct {
	UserName string `json:"userName" uri:"userName" form:"userName"`
	PassWord string `json:"passWord" uri:"passWord" form:"passWord"`
	ErWeiMa  string `json:"erWeiMa" uri:"erWeiMa" form:"erWeiMa"`
}
