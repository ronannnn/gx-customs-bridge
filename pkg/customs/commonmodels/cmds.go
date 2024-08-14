package commonmodels

type MessageRequestPayload struct {
	Data        any    `json:"data"`
	UploadType  string `json:"uploadType"`
	DeclareFlag string `json:"declareFlag"`
}

type MessageResponseResult struct {
	ImpexpMarkcd         string               `json:"impexpMarkcd"`
	UploadType           string               `json:"uploadType"`
	CommonResponeMessage CommonResponeMessage `json:"commonResponeMessage"`
}

type ReceiptResult struct {
	ReceiptType string `json:"receiptType"`
	Data        any    `json:"data"`
}
