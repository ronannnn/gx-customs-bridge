package commonmodels

type CompanyType string

const (
	CompanyTypeGxhg  CompanyType = "gxhg"
	CompanyTypeGxwl  CompanyType = "gxwl"
	CompanyTypeGxgyl CompanyType = "gxgyl"
)

type MessageRequestPayload struct {
	CompanyType CompanyType `json:"companyType" validate:"required,oneof=gxhg gxwl gxgyl"`
	Data        any         `json:"data" validate:"required"`
	UploadType  string      `json:"uploadType" validate:"required,oneof=INV101 SAS121 ICP101"`
	DeclareFlag string      `json:"declareFlag" validate:"required,oneof=0 1"`
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
