package bpjs

import "encoding/json"

func BatalAntreanService(
	req AntrolBatalRequest,
) (any, int, string, error) {

	cfg, err := LoadConfig()
	if err != nil {
		return nil, 500, "config error", err
	}

	res, ts, err := DoRequest(
		"POST",
		"/antrean/batal",
		req,
		cfg,
	)
	if err != nil {
		return nil, 500, "request bpjs error", err
	}

	// ❌ BPJS ERROR (BUKAN ERROR SISTEM)
	if res.Metadata.Code != 200 && res.Metadata.Code != 1 {
		return nil, res.Metadata.Code, res.Metadata.Message, nil
	}

	plain, err := DecryptResponse(
		res.Response,
		cfg.ConsID,
		cfg.SecretKey,
		ts,
		false,
	)
	if err != nil {
		return nil, 500, "decrypt error", err
	}

	var data any
	if err := json.Unmarshal([]byte(plain), &data); err != nil {
		return nil, 500, "json error", err
	}

	return data, res.Metadata.Code, res.Metadata.Message, nil
}
