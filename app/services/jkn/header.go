package jkn

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"net/http"
	"strconv"
	"time"
)

type BPJSHeader struct {
	Timestamp string
	Signature string
}

func GenerateBPJSHeader(consID, secretKey string) BPJSHeader {
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)

	mac := hmac.New(sha256.New, []byte(secretKey))
	mac.Write([]byte(consID + "&" + timestamp))
	signature := base64.StdEncoding.EncodeToString(mac.Sum(nil))

	return BPJSHeader{
		Timestamp: timestamp,
		Signature: signature,
	}
}

func ApplyBPJSHeader(req *http.Request, cfg *BPJSConfig, h BPJSHeader) {
	req.Header.Set("X-cons-id", cfg.ConsID)
	req.Header.Set("X-timestamp", h.Timestamp)
	req.Header.Set("X-signature", h.Signature)
	req.Header.Set("user_key", cfg.UserKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
}
