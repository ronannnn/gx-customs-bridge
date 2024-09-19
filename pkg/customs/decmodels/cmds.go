package decmodels

import "github.com/ronannnn/gx-customs-bridge/pkg/customs/commonmodels"

type DecOperType string

const (
	DecOprTypeA DecOperType = "A" // 报关单上载
	DecOprTypeB DecOperType = "B" // 报关单、转关单上载
	DecOprTypeC DecOperType = "C" // 报关单申报
	DecOprTypeD DecOperType = "D" // 报关单、转关单申报
	DecOprTypeE DecOperType = "E" // 电子手册报关单上载
	DecOprTypeG DecOperType = "G" // 报关单暂存
)

type MessageRequestPayload struct {
	Data        any                      `json:"data" validate:"required"`
	OperType    DecOperType              `json:"operType" validate:"required,oneof=A B C D E G"`
	CompanyType commonmodels.CompanyType `json:"type" validate:"required,oneof=gxhg gxwl gxgyl"`
}

type MessageResponseResult struct {
	ImpexpMarkcd      string            `json:"impexpMarkcd"`
	DecImportResponse DecImportResponse `json:"decImportResponse"`
}
