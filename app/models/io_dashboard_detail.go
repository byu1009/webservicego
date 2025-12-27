package models

type IoDashboardDetail struct {
	DDashId				string	`gorm:"column:ddash_id"`
	DDashParent			string	`gorm:"column:ddash_parent"`
	DDashPoli			string	`gorm:"column:ddash_poli"`
	DDashDokter			string	`gorm:"column:ddash_dokter"`
	DDashStatus			string	`gorm:"column:ddash_status"`
}

func (IoDashboardDetail) TableName() string {
	return "io_dashboard_detail"
}