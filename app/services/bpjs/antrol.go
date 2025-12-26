package bpjs

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"webservicego/app/models"
	"webservicego/config"
)

type Config struct {
	ConsID    string
	SecretKey string
	BaseURL   string
	UserKey   string
}

type Response struct {
	Metadata struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"metadata"`
	Response string `json:"response"`
}

func LoadConfig() (*Config, error) {
	db := config.DBConnect()

	var settings []models.IoSetting
	if err := db.Where("`group` = ?", "bpjs_kesehatan").Find(&settings).Error; err != nil {
		return nil, err
	}

	conf := map[string]string{}
	for _, s := range settings {
		conf[s.SettingOption] = s.Value
	}

	return &Config{
		ConsID:    conf["bpjs_cons_id"],
		SecretKey: conf["bpjs_secret_key"],
		BaseURL:   conf["bpjs_antrol_jbase_url"],
		UserKey:   conf["mjkn_user_key"],
	}, nil
}

func signature(cfg *Config) (string, string) {
	ts := fmt.Sprintf("%d", time.Now().Unix())

	h := hmac.New(sha256.New, []byte(cfg.SecretKey))
	h.Write([]byte(cfg.ConsID + "&" + ts))
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))

	return ts, signature
}

func DoRequest(
	method, url string,
	body interface{},
	cfg *Config,
) (*Response, string, error) {

	ts, sign := signature(cfg)

	var buf io.Reader
	if body != nil {
		b, _ := json.Marshal(body)
		buf = bytes.NewBuffer(b)
	}

	req, _ := http.NewRequest(method, cfg.BaseURL+url, buf)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-cons-id", cfg.ConsID)
	req.Header.Set("X-timestamp", ts)
	req.Header.Set("X-signature", sign)
	req.Header.Set("user_key", cfg.UserKey)

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, "", err
	}
	defer resp.Body.Close()

	raw, _ := io.ReadAll(resp.Body)

	var result Response
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, "", err
	}

	return &result, ts, nil
}
