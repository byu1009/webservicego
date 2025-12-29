package models

type IoRadiologyDicomInstance struct {
	InstanceId		string	`gorm:"column:instance_id;primaryKey"`
	Acsn			string	`gorm:"column:acsn"`
}

func (IoRadiologyDicomInstance) TableName() string {
	return "io_radiology_dicom_instance"
}