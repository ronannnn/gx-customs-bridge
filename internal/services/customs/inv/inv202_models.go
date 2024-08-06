package inv

import (
	"encoding/xml"

	"github.com/ronannnn/gx-customs-bridge/internal/services/customs/common"
)

type InvApprResult struct {
	InvPreentNo  *string `json:"invPreentNo"`  // 核注清单数据中心统一编号(同预录入统一编号)
	BusinessId   *string `json:"businessId"`   // 核注清单编号
	EntrySeqNo   *string `json:"entrySeqNo"`   // 报关单统一编号
	ManageResult *string `json:"manageResult"` // 处理结果(1:生成成功 2:生成失败)
	CreateDate   *string `json:"createDate"`   // 生成日期
	Reason       *string `json:"reason"`       // 生成失败原因
}

// Inv202Xml 核注清单自动生成报关单同一编号报文
type Inv202Xml struct {
	XMLName     xml.Name `xml:"Package"`
	EnvelopInfo common.EnvelopInfo
	DataInfo    struct {
		BusinessData struct {
			Inv202 struct {
				InvApprResult InvApprResult
			} `xml:"INV202"`
		}
	}
}
