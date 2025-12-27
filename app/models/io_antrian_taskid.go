package models

type IoAntrianTaskid struct {
	Nobooking		string	`gorm:"column:nobooking;primaryKey"`
	Taskid3			string	`gorm:"column:taskid_3"`
	Taskid3Send		string	`gorm:"column:taskid_3_send"`
	Taskid4			string	`gorm:"column:taskid_4"`
	Taskid4Send		string	`gorm:"column:taskid_4_send"`
	Taskid5			string	`gorm:"column:taskid_5"`
	Taskid5Send		string	`gorm:"column:taskid_5_send"`
	Taskid6			string	`gorm:"column:taskid_6"`
	Taskid6Send		string	`gorm:"column:taskid_6_send"`
	Taskid7			string	`gorm:"column:taskid_7"`
	Taskid7Send		string	`gorm:"column:taskid_7_send"`
	Taskid99		string	`gorm:"column:taskid_99"`
	Taskid99Send	string	`gorm:"column:taskid_99_send"`
}

func (IoAntrianTaskid) TableName() string {
	return "io_antrian_taskid"
}