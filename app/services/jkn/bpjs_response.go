package jkn

type BPJSResponse struct {
	Metadata struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"metadata"`
	Response string `json:"response"`
}
