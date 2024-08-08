package common

type MessageResponseResult struct {
	Id                   string               `json:"id"`
	CommonResponeMessage CommonResponeMessage `json:"commonResponeMessage"`
}

type ReceiptResult struct {
	ReceiptType string `json:"receiptType"`
	Data        any    `json:"data"`
}
