package jkn

import (
	"crypto/aes"
	"crypto/cipher"
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

	lzstring "github.com/daku10/go-lz-string"
	"github.com/gin-gonic/gin"
)

// BPJSConfig menyimpan semua konfigurasi BPJS
type BPJSConfig struct {
	ConsID    string
	SecretKey string
	AntrolURL string
	UserKey   string
}

// PKCS7 unpad helper
func pkcs7Unpad(data []byte) ([]byte, error) {
	length := len(data)
	if length == 0 {
		return nil, fmt.Errorf("invalid padding size")
	}
	pad := int(data[length-1])
	if pad > length || pad == 0 {
		return nil, fmt.Errorf("invalid padding")
	}
	return data[:length-pad], nil
}

// RefPoli handler
func RefPoli(c *gin.Context) {
	// --- 1. Connect database ---
	db := config.DBConnect()

	// --- 2. Ambil semua setting BPJS ---
	var settings []models.IoSetting
	if err := db.Where("`group` = ?", "bpjs_kesehatan").Find(&settings).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": err.Error(), "data": nil})
		return
	}
	if len(settings) == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "BPJS settings not found", "data": nil})
		return
	}

	conf := make(map[string]string)
	for _, s := range settings {
		conf[s.SettingOption] = s.Value
	}

	cfg := &BPJSConfig{
		ConsID:    conf["bpjs_cons_id"],
		SecretKey: conf["bpjs_secret_key"],
		AntrolURL: conf["bpjs_antrol_jbase_url"],
		UserKey:   conf["mjkn_user_key"],
	}

	// --- 3. Timestamp & signature ---
	tStamp := fmt.Sprintf("%d", time.Now().Unix())
	h := hmac.New(sha256.New, []byte(cfg.SecretKey))
	h.Write([]byte(cfg.ConsID + "&" + tStamp))
	signature := h.Sum(nil)
	encodedSignature := base64.StdEncoding.EncodeToString(signature)

	// --- 4. Request GET ke BPJS ---
	client := &http.Client{Timeout: 30 * time.Second}
	req, _ := http.NewRequest("GET", cfg.AntrolURL+"/ref/poli", nil)
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-cons-id", cfg.ConsID)
	req.Header.Set("X-timestamp", tStamp)
	req.Header.Set("X-signature", encodedSignature)
	req.Header.Set("user_key", cfg.UserKey)

	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": err.Error(), "data": nil})
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": err.Error(), "data": nil})
		return
	}

	var decode struct {
		Metadata struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		} `json:"metadata"`
		Response string `json:"response"`
	}

	if err := json.Unmarshal(body, &decode); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": err.Error(), "data": nil})
		return
	}

	if decode.Metadata.Code != 1 {
		c.JSON(http.StatusOK, gin.H{
			"code":    decode.Metadata.Code,
			"message": decode.Metadata.Message,
			"data":    nil,
		})
		return
	}

	// --- 5. Base64 decode ---
	ciphertext, err := base64.StdEncoding.DecodeString(decode.Response)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "Invalid base64 response", "data": nil})
		return
	}

	// --- 6. AES-256-CBC decrypt ---
	key := cfg.ConsID + cfg.SecretKey + tStamp
	keyHash := sha256.Sum256([]byte(key))
	iv := keyHash[:16]

	block, err := aes.NewCipher(keyHash[:])
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "AES error: " + err.Error(), "data": nil})
		return
	}

	if len(ciphertext)%aes.BlockSize != 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "Ciphertext is not multiple of block size", "data": nil})
		return
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	plaintext := make([]byte, len(ciphertext))
	mode.CryptBlocks(plaintext, ciphertext)

	// --- 7. PKCS7 unpad ---
	unpad, err := pkcs7Unpad(plaintext)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "Unpad error: " + err.Error(), "data": nil})
		return
	}

	// --- 8. LZString decompress ---
	decompressed, err := lzstring.DecompressFromEncodedURIComponent(string(unpad))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "Decompress error: " + err.Error(), "data": nil})
		return
	}

	// --- 9. Decode JSON hasil decompress ---
	var data interface{}
	if err := json.Unmarshal([]byte(decompressed), &data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "JSON decode error: " + err.Error(), "data": nil})
		return
	}

	// --- 10. Return response ---
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "Ok",
		"data":    data,
	})
}