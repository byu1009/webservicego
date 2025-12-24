package models

type IoSetting struct {
	SettingOption string `gorm:"column:setting_option;primaryKey" json:"setting_option"`
	Group         string `gorm:"column:group" json:"group"`
	Value         string `gorm:"column:value" json:"value"`
}

// TableName menentukan nama tabel secara eksplisit
func (IoSetting) TableName() string {
	return "io_settings"
}