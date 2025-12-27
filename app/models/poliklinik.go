package models

type Poliklinik struct {
	KdPoli				string	`gorm:"column:kd_poli"`
	NmPoli				string	`gorm:"column:nm_poli"`
	Registrasi			string	`gorm:"column:registrasi"`
	RegistrasiLama		string	`gorm:"column:registrasi_lama"`
	Status				string	`gorm:"column:status"`
}

func (Poliklinik) TableName() string {
	return "poliklinik"
}