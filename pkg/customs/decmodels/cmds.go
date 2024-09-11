package decmodels

type MessageResponseResult struct {
	ImpexpMarkcd      string            `json:"impexpMarkcd"`
	DecImportResponse DecImportResponse `json:"decImportResponse"`
}
