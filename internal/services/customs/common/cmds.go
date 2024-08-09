package common

type MessageRequestPayload struct {
	Data        any    `json:"data"`
	UploadType  string `json:"uploadType"`
	DeclareFlag string `json:"declareFlag"`
}

type MessageResponseResult struct {
	Id                   string               `json:"id"`
	UploadType           string               `json:"uploadType"`
	CommonResponeMessage CommonResponeMessage `json:"commonResponeMessage"`
}

type ReceiptResult struct {
	ReceiptType string `json:"receiptType"`
	Data        any    `json:"data"`
}
