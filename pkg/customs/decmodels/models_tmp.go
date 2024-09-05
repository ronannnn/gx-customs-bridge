package decmodels

import "encoding/xml"

type DecHeadTmp struct {
	// 申报单位
	AgentCode string `json:"agentCode"` // 编码
	AgentName string `json:"agentName"` // 名称
	// 录入单位
	CopCode string `json:"copCode"` // 编码
	CopName string `json:"copName"` // 名称
	// 境内收发货人
	TradeCode string `json:"tradeCode"` // 编码
	TradeName string `json:"tradeName"` // 名称

	CustomMaster  string `json:"customMaster"`  // 主管海关代码(申报地海关)
	IEPort        string `json:"iEPort"`        // 进出境关别
	Type          string `json:"type"`          // 单据类型
	IEFlag        string `json:"iEFlag"`        // 进出口标志(I进口 E出口)
	TradeMode     string `json:"tradeMode"`     // 监管方式
	TrafMode      string `json:"trafMode"`      // 运输方式
	TrafName      string `json:"trafName"`      // 运输工具名称
	TradeAreaCode string `json:"tradeAreaCode"` // 贸易国别代码

	DeclTrnRel   string `json:"declTrnRel"`   // 报关/转关关系标志(0：一般报关单 1：转关提前报关单)
	EntryType    string `json:"entryType"`    // 报关单类型(0普通报关单，L为带报关单清单的报关单，W无纸报关类型，D既是清单又是无纸报关的情况，M：无纸化通关)
	PromiseItmes string `json:"promiseItmes"` // 价格说明

	GrossWet string `json:"grossWet"` // 毛重

	InputerName string `json:"inputerName"` // 录入员姓名
	DeclareName string `json:"declareName"` // 申报人员姓名
	TypistNo    string `json:"typistNo"`    // 录入员IC卡号
}

type DecListTmp struct {
	CodeTS             string `json:"codeTS"`             // 商品编码
	GNo                string `json:"gNo"`                // 商品序号
	GName              string `json:"gName"`              // 商品名称
	GQty               string `json:"gQty"`               // 成交数量
	GUnit              string `json:"gUnit"`              // 成交计量单位
	TradeCurr          string `json:"tradeCurr"`          // 成交币制
	DeclTotal          string `json:"declTotal"`          // 申报总价
	OriginCountry      string `json:"originCountry"`      // 原产国
	DestinationCountry string `json:"destinationCountry"` // 最终目的国
}

type DecContainerTmp struct {
	ContainerId string `json:"containerId"` // 集装箱号
	ContainerMd string `json:"containerMd"` // 集装箱规格
	GoodsNo     string `json:"goodsNo"`     // 商品项号
}

type DecFreeTxtTmp struct {
	VoyNo string `json:"voyNo"` // 航次号
}

type DecSignTmp struct {
	OperType    string `json:"operType"`                       // 操作类型
	ClientSeqNo string `json:"clientSeqNo" gorm:"uniqueIndex"` // 客户端自行编制的编号，唯一识别一票报关单
	ICCode      string `json:"iCCode"`                         // 操作员IC卡号
	OperName    string `json:"operName"`                       // 操作员姓名
}

type DecLicenseDocuTmp struct {
}

type DecEdocRealationTmp struct {
	EdocID        string `json:"edocID"`        // 文件名、随附单据编号
	EdocCode      string `json:"edocCode"`      // 随附单证类别
	EdocFomatType string `json:"edocFomatType"` // 随附单证格式类型
	OpNote        string `json:"opNote"`        // 操作说明
	EdocCopId     string `json:"edocCopId"`     // 电子单证企业内部编号
	EdocOwnerCode string `json:"edocOwnerCode"` // 电子单证所属企业代码
	SignUnit      string `json:"signUnit"`      // 签发单位
	SignTime      string `json:"signTime"`      // 签发时间
}

type DecTmp struct {
	DecHead           DecHeadTmp
	DecLists          []DecListTmp
	DecContainers     []DecContainerTmp
	DecFreeTxt        DecFreeTxtTmp
	DecSign           DecSignTmp
	DecLicenseDocus   []DecLicenseDocuTmp
	DecEdocRealations []DecEdocRealationTmp
}

type DecTmpXml struct {
	XMLName  xml.Name `xml:"DecMessage"`
	Version  string   `xml:"Version,attr"`
	Xmlns    string   `xml:"xmlns,attr"`
	DecHead  DecHeadTmp
	DecLists struct {
		DecLists []DecListTmp `xml:"DecList"`
	}
	DecContainers struct {
		DecContainers []DecContainerTmp `xml:"Container"`
	}
	// DecLicenseDocus struct {
	// 	DecLicenseDocus []DecLicenseDocuTmp `xml:"LicenseDocu"`
	// }
	// EdocRealations []DecEdocRealationTmp `xml:"EdocRealation"`
	DecFreeTxt DecFreeTxtTmp
	DecSign    DecSignTmp
}
