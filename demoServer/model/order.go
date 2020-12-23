package model

type Order struct {
	Id        string `xorm:"not null pk autoincr "`
	OrderTime string `xorm:" 'ordertime' comment('时间') VARCHAR(255)" json:"ordertime"`
	Total     uint32 `xorm:"default '' comment('合计') VARCHAR(50)"`
	Uid       uint32 `xorm:"default '' comment('uid') VARCHAR(50)"`
}
