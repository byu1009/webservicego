package models

type IoAssessmentPoliAccess struct {
	KdPoli				string	`gorm:"column:kd_poli;primaryKey"`
	Asessment			string	`gorm:"column:asessment"`
	AssessmentNurse		string	`gorm:"column:assessment_nurse"`
}

func (IoAssessmentPoliAccess) TableName() string {
	return "io_assessment_poli_access"
}