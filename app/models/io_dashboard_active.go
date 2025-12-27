package models

type IoDashboardActive struct {
	DashacTgl			string	`gorm:"column:dashac_tgl"`
	DashacIdddash		string	`gorm:"column:dashac_idddash"`
	DashacStatus		string	`gorm:"column:dashac_status"`
}

func (IoDashboardActive) TableName() string {
	return "io_dashboard_active"
}