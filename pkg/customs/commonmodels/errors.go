package commonmodels

import "fmt"

var (
	ErrMessageWithoutEnvelopeInfo = fmt.Errorf("报文没有EnvelopeInfo")
	ErrParseInv101                = fmt.Errorf("case INV101失败")
	ErrParseSas121                = fmt.Errorf("case SAS121失败")
)
