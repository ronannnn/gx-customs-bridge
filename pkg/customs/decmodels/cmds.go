package decmodels

type MessageRequestPayload struct {
	Data     any    `json:"data"`
	OperType string `json:"operType"`
}

type MessageResponseResult struct {
	ImpexpMarkcd      string            `json:"impexpMarkcd"`
	DecImportResponse DecImportResponse `json:"decImportResponse"`
}
