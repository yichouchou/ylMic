package model

type Adviertisement struct {
	Id    uint8  `xorm:" 'id' not null pk autoincr "`
	Name  string `xorm:" 'name' comment('时间') VARCHAR(255)" json:"Name"`
	Title string `xorm:" 'title' comment('合计') VARCHAR(50)"`
	Tel   string `xorm:" 'tel' comment('uid') "`
	Type  uint8  `xorm:" 'Type' comment('uid') "`
}
