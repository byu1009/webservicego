package jkn

import (
	"errors"
	"webservicego/app/models"

	"gorm.io/gorm"
)

type BPJSConfig struct {
	ConsID     string
	SecretKey string
	AntrolURL string
	UserKey   string
}

func GetBPJSConfig(db *gorm.DB) (*BPJSConfig, error) {
	var settings []models.IoSetting

	// Ambil semua setting bpjs_kesehatan
	if err := db.Where("`group` = ?", "bpjs_kesehatan").Find(&settings).Error; err != nil {
		return nil, err
	}

	if len(settings) == 0 {
		return nil, errors.New("BPJS settings not found")
	}

	conf := make(map[string]string)
	for _, s := range settings {
		conf[s.SettingOption] = s.Value
	}

	// mapping ke struct BPJSConfig
	cfg := &BPJSConfig{
		ConsID:    conf["bpjs_cons_id"],
		SecretKey: conf["bpjs_secret_key"],
		AntrolURL: conf["bpjs_antrol_jbase_url"],
		UserKey:   conf["mjkn_user_key"],
	}

	return cfg, nil
}

