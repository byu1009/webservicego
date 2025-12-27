package models

type IoDashboard struct {
	DashId			string	`gorm:"column:dash_id"`
	DashName		string	`gorm:"column:dash_name"`
	DashType		string	`gorm:"column:dash_type"`
	DashStatus		string	`gorm:"column:dash_status"`
}

func (IoDashboard) TableName() string {
	return "io_dashboard"
}