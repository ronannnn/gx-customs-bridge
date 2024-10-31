package commonmodels

import "encoding/xml"

type CommonResponeMessage struct {
	XMLName      xml.Name `xml:"CommonResponeMessage"`
	SeqNo        string   `json:"seqNo"`        // 预录入统一编号
	EtpsPreentNo string   `json:"etpsPreentNo"` // 企业内部编号
	CheckInfo    string   `json:"checkInfo"`    // 响应信息
	DealFlag     string   `json:"dealFlag"`     // 响应代码(0-导入成功 其他值-导入失败)
}

// 信封内容
type EnvelopInfo struct {
	Version     string `json:"version" xml:"version"`          // 版本号
	MessageId   string `json:"messageId" xml:"message_id"`     // 报文编号
	Filename    string `json:"filename" xml:"file_name"`       // 文件名
	MessageType string `json:"messageType" xml:"message_type"` // 报文类型
	SenderId    string `json:"senderId" xml:"sender_id"`       // 发送方
	ReceiverId  string `json:"receiverId" xml:"receiver_id"`   // 接收方
	SendTime    string `json:"sendTime" xml:"send_time"`       // 发送时间
}

type PocketInfo struct {
	PocketId       string `json:"pocketId" xml:"pocket_id"`              // 分组编号
	TotalPocketQty int    `json:"totalPocketQty" xml:"total_pocket_qty"` // 总分组数
	CurPocketNo    int    `json:"curPocketNo" xml:"cur_pocket_no"`       // 当前分组号
	IsUnstructured string `json:"isUnstructured" xml:"is_unstructured"`  // 是否无结构
}

// 审核回执信息
type HdeApprResult struct {
	EtpsPreentNo string `json:"etpsPreentNo" xml:"etpsPreentNo"` // 企业预录入编号(内部编号)
	BusinessId   string `json:"businessId" xml:"businessId"`     // 业务编号
	TmsCnt       string `json:"tmsCnt" xml:"tmsCnt"`             // 变更/报核次数
	Typecd       string `json:"typecd" xml:"typecd"`             // 业务类型
	ManageResult string `json:"manageResult" xml:"manageResult"` // 处理结果
	ManageDate   string `json:"manageDate" xml:"manageDate"`     // 处理日期
	Rmk          string `json:"rmk" xml:"rmk"`                   // 备注
}

// 回执共有的头部，主要用于获取EnvelopeInfo中的MessageType
type ReceiptMessageHeader struct {
	XMLName     xml.Name `xml:"Package"`
	EnvelopInfo EnvelopInfo
}

// 检查信息
type CheckInfo struct {
	XMLName xml.Name `xml:"CheckInfo"`
	Note    string   `json:"note" xml:"note"` // 检查信息
}

// 推送到mq数据的公共头
type MqDataCommonPayload struct {
	Tried uint `json:"tried"` // 重试次数
}

func (m *MqDataCommonPayload) TryAgain() {
	m.Tried++
}

type MqDataCommonInterface interface {
	TryAgain()
}
