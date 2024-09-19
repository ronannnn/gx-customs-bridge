package decmodels

import (
	"encoding/xml"

	"github.com/ronannnn/infra/models"
)

type DecHead struct {
	// 编码
	SeqNo   *string `json:"seqNo" validate:"omitempty,len=18"`   // 数据中心生成的对报关单的唯一标识
	EntryId *string `json:"entryId" validate:"omitempty,len=18"` // 报关单号
	BillNo  *string `json:"billNo" validate:"required"`          // 提运单号
	// 申报单位
	AgentCode *string `json:"agentCode" validate:"required,len=10"`    // 编码
	AgentName *string `json:"agentName" validate:"required,not_blank"` // 名称
	// 录入单位
	CopCode *string `json:"copCode" validate:"required,len=10"`    // 编码
	CopName *string `json:"copName" validate:"required,not_blank"` // 名称
	// 境内收发货人
	TradeCode *string `json:"tradeCode" validate:"required,len=10"`    // 编码
	TradeName *string `json:"tradeName" validate:"required,not_blank"` // 名称

	CustomMaster  *string `json:"customMaster" validate:"required,len=4,numeric"` // 主管海关代码(申报地海关)
	IEPort        *string `json:"iEPort" validate:"required,len=4,numeric"`       // 进出境关别
	Type          *string `json:"type" validate:"required,max=6"`                 // 单据类型
	IEFlag        *string `json:"iEFlag" validate:"required,oneof=I E"`           // 进出口标志(I进口 E出口)
	TradeMode     *string `json:"tradeMode" validate:"required,len=4,numeric"`    // 监管方式
	TrafMode      *string `json:"trafMode" validate:"required,len=1,numeric"`     // 运输方式
	TrafName      *string `json:"trafName" validate:"required"`                   // 运输工具名称
	TradeAreaCode *string `json:"tradeAreaCode" validate:"required,len=3,alpha"`  // 贸易国别代码

	DeclTrnRel   *string `json:"declTrnRel" validate:"required,oneof=0 1"`      // 报关/转关关系标志(0：一般报关单 1：转关提前报关单)
	EntryType    *string `json:"entryType" validate:"required,oneof=0 L W D M"` // 报关单类型(0普通报关单，L为带报关单清单的报关单，W无纸报关类型，D既是清单又是无纸报关的情况，M：无纸化通关)
	PromiseItmes *string `json:"promiseItmes" validate:"required"`              // 价格说明

	GrossWet *models.DecimalSafe `json:"grossWet" validate:"required,d_gt=0"` // 毛重

	InputerName *string `json:"inputerName" validate:"omitempty"` // 录入员姓名(配置文件填写)
	DeclareName *string `json:"declareName" validate:"omitempty"` // 申报人员姓名(配置文件填写)
	TypistNo    *string `json:"typistNo" validate:"omitempty"`    // 录入员IC卡号(配置文件填写)
}

type DecList struct {
	CodeTS             *string             `json:"codeTS" validate:"required,numeric"`                 // 商品编码
	GNo                *string             `json:"gNo" validate:"required,numeric"`                    // 商品序号
	GName              *string             `json:"gName" validate:"required"`                          // 商品名称
	GQty               *models.DecimalSafe `json:"gQty" validate:"required,d_gt=0"`                    // 成交数量
	GUnit              *string             `json:"gUnit" validate:"required,len=3,numeric"`            // 成交计量单位
	TradeCurr          *string             `json:"tradeCurr" validate:"required,len=3,numeric"`        // 成交币制
	DeclTotal          *models.DecimalSafe `json:"declTotal" validate:"required,d_gt=0"`               // 申报总价
	OriginCountry      *string             `json:"originCountry" validate:"required,len=3,alpha"`      // 原产国
	DestinationCountry *string             `json:"destinationCountry" validate:"required,len=3,alpha"` // 最终目的国
}

type DecContainer struct {
	ContainerId *string `json:"containerId" validate:"required,len=11,alphanum"` // 集装箱号
	ContainerMd *string `json:"containerMd" validate:"required"`                 // 集装箱规格
	GoodsNo     *string `json:"goodsNo" validate:"required,numeric"`             // 商品项号
}

type DecFreeTxt struct {
	VoyNo *string `json:"voyNo" validate:"required"` // 航次号
}

type DecSign struct {
	OperType    *DecOperType `json:"operType" validate:"omitempty,oneof=A B C D E G"`    // 操作类型
	ClientSeqNo *string      `json:"clientSeqNo" validate:"required" gorm:"uniqueIndex"` // 客户端自行编制的编号，唯一识别一票报关单
	ICCode      *string      `json:"iCCode" validate:"omitempty"`                        // 操作员IC卡号(配置文件填写)
	OperName    *string      `json:"operName" validate:"omitempty"`                      // 操作员姓名(配置文件填写)
}

type Dec struct {
	DecHead       DecHead        `json:"decHead"`
	DecLists      []DecList      `json:"decLists"`
	DecContainers []DecContainer `json:"decContainers"`
	DecFreeTxt    DecFreeTxt     `json:"decFreeTxt"`
	DecSign       DecSign        `json:"decSign"`
}

type DecXml struct {
	XMLName  xml.Name `xml:"DecMessage"`
	Version  string   `xml:"Version,attr"`
	Xmlns    string   `xml:"xmlns,attr"`
	DecHead  DecHead
	DecLists struct {
		DecLists []DecList `xml:"DecList"`
	}
	DecContainers struct {
		DecContainers []DecContainer `xml:"Container"`
	}
	DecFreeTxt DecFreeTxt
	DecSign    DecSign
}
