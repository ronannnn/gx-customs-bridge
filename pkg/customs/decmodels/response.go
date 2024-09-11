package decmodels

import "encoding/xml"

type DecImportResponse struct {
	XMLName      xml.Name `xml:"DecImportResponse"`
	Xmlns        string   `xml:"xmlns,attr"`
	SeqNo        string   `json:"seqNo"`        // 数据中心生成的对报关单的唯一标识
	ResponseCode string   `json:"responsCode"`  // 0：导入成功。其它值：导入失败，失败信息在ErrorMessage字段中。
	ErrorMessage string   `json:"errorMessage"` // 报关单服务错误信息
	ClientSeqNo  string   `json:"clientSeqNo"`  // 客户端对该票报关单的唯一标识
	TrnPreId     string   `json:"trnPreId"`     // 转关单统一编号(普通报关单为空)
}

type DecResult struct {
	XMLName      xml.Name `xml:"DEC_RESULT"`
	CusCiqNo     string   `json:"cusCiqNo" xml:"CUS_CIQ_NO"`        // 数据中心统一编号
	EntryId      string   `json:"entryId" xml:"ENTRY_ID"`           // 报关单编号
	NoticeDate   string   `json:"noticeDate" xml:"NOTICE_DATE"`     // 通知时间
	Channel      string   `json:"channel" xml:"CHANNEL"`            // 回执代码
	Note         string   `json:"note" xml:"NOTE"`                  // 回执说明
	CustomMaster string   `json:"customMaster" xml:"CUSTOM_MASTER"` // 申报地海关
	IEDate       string   `json:"iEDate" xml:"I_E_DATE"`            // 进出口日期
	DDate        string   `json:"dDate" xml:"D_DATE"`               // 申报日期
}
