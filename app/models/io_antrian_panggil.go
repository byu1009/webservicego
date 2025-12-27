package models

type IoAntrianPanggil struct {
	NoReferensi		string	`gorm:"column:no_referensi"`
	DashboardId		string	`gorm:"column:dashboard_id"`
	Type			string	`gorm:"column:type"`
	Counter			string	`gorm:"column:counter"`
}

func (IoAntrianPanggil) TableName() string {
	return "io_antrian_panggil"
}