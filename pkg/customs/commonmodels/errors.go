package commonmodels

import "fmt"

var (
	ErrMessageWithoutEnvelopeInfo = fmt.Errorf("报文没有EnvelopeInfo")
	ErrParseInv101                = fmt.Errorf("解析INV101失败")
	ErrParseSas121                = fmt.Errorf("解析SAS121失败")
	ErrParseIcp101                = fmt.Errorf("解析ICP101s失败")
)
