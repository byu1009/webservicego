package models

import "time"

type IoAntrianTaskid struct {
	Nobooking		string		`gorm:"column:nobooking;primaryKey"`
	Taskid3			*time.Time	`gorm:"column:taskid_3"`
	Taskid3Send		*time.Time	`gorm:"column:taskid_3_send"`
	Taskid4			*time.Time	`gorm:"column:taskid_4"`
	Taskid4Send		*time.Time	`gorm:"column:taskid_4_send"`
	Taskid5			*time.Time	`gorm:"column:taskid_5"`
	Taskid5Send		*time.Time	`gorm:"column:taskid_5_send"`
	Taskid6			*time.Time	`gorm:"column:taskid_6"`
	Taskid6Send		*time.Time	`gorm:"column:taskid_6_send"`
	Taskid7			*time.Time	`gorm:"column:taskid_7"`
	Taskid7Send		*time.Time	`gorm:"column:taskid_7_send"`
	Taskid99		*time.Time	`gorm:"column:taskid_99"`
	Taskid99Send	*time.Time	`gorm:"column:taskid_99_send"`
}

func (IoAntrianTaskid) TableName() string {
	return "io_antrian_taskid"
}