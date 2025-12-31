package bpjs

import (
	"encoding/json"
	"errors"
)

func AntreanPerTanggalService(
	req AntrolPerTanggalRequest,
) (any, error) {

	cfg, err := LoadConfig()
	if err != nil {
		return nil, err
	}

	res, ts, err := DoRequest(
		"GET",
		"/antrean/pendaftaran/tanggal/"+req.Tanggal,
		nil,
		cfg,
	)
	if err != nil {
		return nil, err
	}

	if res.Metadata.Code != 200 && res.Metadata.Code != 1 {
		return nil, errors.New(res.Metadata.Message)
	}

	plain, err := DecryptResponse(
		res.Response,
		cfg.ConsID,
		cfg.SecretKey,
		ts,
		true,
	)
	if err != nil {
		return nil, err
	}

	var data any
	if err := json.Unmarshal([]byte(plain), &data); err != nil {
		return nil, err
	}

	return data, nil
}