package sas

import (
	"encoding/xml"

	"github.com/ronannnn/gx-customs-bridge/internal/services/customs/common"
	"github.com/shopspring/decimal"
)

// Inv101Head 核注清单表头
type Inv101Head struct {
	// 非业务相关字段
	ChgTmsCnt *string `json:"chgTmsCnt"` // 变更次数

	// 编号
	SeqNo           *string `json:"seqNo"`           // 预录入统一编号
	EtpsInnerInvtNo *string `json:"etpsInnerInvtNo"` // 企业内部清单编号
	BondInvtNo      *string `json:"bondInvtNo"`      // 保税清单编号
	PutrecNo        *string `json:"putrecNo"`        // 备案编号，手(账)册编号

	// 清单信息
	DclTypecd      *string `json:"dclTypecd"`      // 申报类型(1-备案申请 2-变更申请 3-删除申请)
	InvtType       *string `json:"invtType"`       // 清单类型代码
	BondInvtTypecd *string `json:"bondInvtTypecd"` // 清单类型代码
	ListStat       *string `json:"listStat"`       // 清单状态
	ListType       *string `json:"listType"`       // 流转类型

	// 核放相关
	InvtIochkptStucd   *string `json:"invtIochkptStucd"`   // 清单进出卡口状态代码
	PassportUsedTypeCd *string `json:"passportUsedTypeCd"` // 核放单生成标志代码
	VrfdedMarkcd       *string `json:"vrfdedMarkcd"`       // 核扣标记代码(0-未核扣 1-预核扣 2-已核扣 3-已核销)

	// 经营企业
	BizopEtpsSccd *string `json:"bizopEtpsSccd"` // 社会信用代码
	BizopEtpsno   *string `json:"bizopEtpsno"`   // 编号
	BizopEtpsNm   *string `json:"bizopEtpsNm"`   // 名称
	// 申报企业
	DclEtpsSccd *string `json:"dclEtpsSccd"` // 社会信用代码
	DclEtpsno   *string `json:"dclEtpsno"`   // 编号
	DclEtpsNm   *string `json:"dclEtpsNm"`   // 名称
	// 收发货企业
	RvsngdEtpsSccd *string `json:"rvsngdEtpsSccd"` // 社会信用代码
	RcvgdEtpsno    *string `json:"rcvgdEtpsno"`    // 编号
	RcvgdEtpsNm    *string `json:"rcvgdEtpsNm"`    // 名称
	// 录入单位
	InputCreditCode *string `json:"inputCreditCode"` // 社会信用代码
	InputCode       *string `json:"inputCode"`       // 编号
	InputName       *string `json:"inputName"`       // 名称

	// 日期
	InputTime        *string `json:"inputTime"`        // 录入日期(格式：20240101)
	InvtDclTime      *string `json:"invtDclTime"`      // 清单申报时间(格式：20240101)
	EntryDclTime     *string `json:"entryDclTime"`     // 报关单申报时间(格式：20240101)
	PrevdTime        *string `json:"prevdTime"`        // 预核扣时间
	FormalVrfdedTime *string `json:"formalVrfdedTime"` // 正式核扣时间

	// 代码
	ImpexpMarkcd      *string `json:"impexpMarkcd"`      // 进出口标记代码(I-进口 E-出口)
	MtpckEndprdMarkcd *string `json:"mtpckEndprdMarkcd"` // 料件成品标记代码(I-料件 E-成品)
	SupvModecd        *string `json:"supvModecd"`        // 监管方式代码
	TrspModecd        *string `json:"trspModecd"`        // 运输方式代码
	ImpexpPortcd      *string `json:"impexpPortcd"`      // 进出境关别代码(进出口口岸代码)
	DclPlcCuscd       *string `json:"dclPlcCuscd"`       // 主管海关(申报地关区代码)
	StshipTrsarvNatcd *string `json:"stshipTrsarvNatcd"` // 起运/运抵国(地区)

	// 报关相关
	DclcusFlag        *string `json:"dclcusFlag"`        // 是否报关标志(1-报关 2-非报关)
	GenDecFlag        *string `json:"genDecFlag"`        // 是否生成报关单:1-生成 2-不生成
	DecType           *string `json:"decType"`           // 报关单类型
	DclcusTypecd      *string `json:"dclcusTypecd"`      // 报关类型代码
	NeedEntryModified *string `json:"needEntryModified"` // 报关单同步修改标志
	EntryStucd        *string `json:"entryStucd"`        // 报关状态(0：未放行，1：已放行) 该类型清单满足两个条件才能核扣：报关单被放行+货物全部过卡

	// 对应报关相关
	EntryNo *string `json:"entryNo"` // 对应报关单编号
	// 对应报关单申报单位
	CorrEntryDclEtpsSccd *string `json:"corrEntryDclEtpsSccd"` // 社会信用代码
	CorrEntryDclEtpsNo   *string `json:"corrEntryDclEtpsNo"`   // 编号
	CorrEntryDclEtpsNm   *string `json:"corrEntryDclEtpsNm"`   // 名称

	// 关联报关相关
	RltInvtNo   *string `json:"rltInvtNo"`   // 关联清单编号
	RltPutrecNo *string `json:"rltPutrecNo"` // 关联备案编号，关联手(账)册编号
	RltEntryNo  *string `json:"rltEntryNo"`  // 关联报关单编号
	// 关联报关单境内收发货人
	RltEntryBizopEtpsSccd *string `json:"rltEntryBizopEtpsSccd"` // 社会信用代码
	RltEntryBizopEtpsno   *string `json:"rltEntryBizopEtpsno"`   // 编号
	RltEntryBizopEtpsNm   *string `json:"rltEntryBizopEtpsNm"`   // 名称
	// 关联报关单生产销售(消费使用)单位
	RltEntryRvsngdEtpsSccd *string `json:"rltEntryRvsngdEtpsSccd"` // 社会信用代码
	RltEntryRcvgdEtpsno    *string `json:"rltEntryRcvgdEtpsno"`    // 编号
	RltEntryRcvgdEtpsNm    *string `json:"rltEntryRcvgdEtpsNm"`    // 名称
	// 关联报关单申报单位
	RltEntryDclEtpsSccd *string `json:"rltEntryDclEtpsSccd"` // 社会信用代码
	RltEntryDclEtpsno   *string `json:"rltEntryDclEtpsno"`   // 编号
	RltEntryDclEtpsNm   *string `json:"rltEntryDclEtpsNm"`   // 名称

	// others
	IcCardNo  *string `json:"icCardNo"`  // 申报人IC卡号
	ApplyNo   *string `json:"applyNo"`   // 申请编号
	Rmk       *string `json:"rmk"`       // 备注
	LevyBlAmt *string `json:"levyBlAmt"` // 计征金额
}

// Inv101List 核注清单表体
type Inv101List struct {
	SeqNo         *string `json:"seqNo"`         // 中心统一编号
	GdsSeqno      *string `json:"gdsSeqno"`      // 商品序号
	EntryGdsSeqno *string `json:"entryGdsSeqno"` // 报关单商品序号
	PutrecSeqno   *string `json:"putrecSeqno"`   // 备案序号(对应底账序号）

	Gdecd            *string `json:"gdecd"`            // 商品编码
	GdsMtno          *string `json:"gdsMtno"`          // 商品料号
	GdsNm            *string `json:"gdsNm"`            // 商品名称
	GdsSpcfModelDesc *string `json:"gdsSpcfModelDesc"` // 商品规格型号

	LawfUnitcd     *string          `json:"lawfUnitcd"`     // 法定计量单位代码
	LawfQty        *decimal.Decimal `json:"lawfQty"`        // 法定数量
	SecdLawfUnitcd *string          `json:"secdLawfUnitcd"` // 第二法定计量单位代码
	SecdLawfQty    *decimal.Decimal `json:"secdLawfQty"`    // 第二法定数量

	DclUnitcd       *string          `json:"dclUnitcd"`       // 申报计量单位代码
	DclQty          *decimal.Decimal `json:"dclQty"`          // 申报数量
	DclCurrcd       *string          `json:"dclCurrcd"`       // 申报币制代码
	DclUprcAmt      *decimal.Decimal `json:"dclUprcAmt"`      // 企业申报单价
	DclTotalAmt     *decimal.Decimal `json:"dclTotalAmt"`     // 企业申报总价
	UsdStatTotalAmt *decimal.Decimal `json:"usdStatTotalAmt"` // 美元统计总金额

	Natcd            *string `json:"natcd"`            // 原产国（地区）代码
	DestinationNatcd *string `json:"destinationNatcd"` // 最终目的国（地区）代码
	LvyrlfModecd     *string `json:"lvyrlfModecd"`     // 征免方式代码
	ClyMarkcd        *string `json:"clyMarkcd"`        // 危化品标志

	WtSfVal      *decimal.Decimal `json:"wtSfVal"`      // 重量比例因子值
	FstSfVal     *decimal.Decimal `json:"fstSfVal"`     // 第一比例因子值
	SecdSfVal    *decimal.Decimal `json:"secdSfVal"`    // 第二比例因子值
	GrossWt      *decimal.Decimal `json:"grossWt"`      // 毛重
	NetWt        *decimal.Decimal `json:"netWt"`        // 净重
	UcnsVerno    *string          `json:"ucnsVerno"`    // 单耗版本号
	UseCd        *string          `json:"useCd"`        // 重点商品标识：0-非重点商品、1-目录重点商品、2-连带重点商品、3-过渡期重点商品
	ApplyTbSeqno *string          `json:"applyTbSeqno"` // 流转申报表序号

	ModfMarkcd *string `json:"modfMarkcd"` // 修改标志
	Rmk        *string `json:"rmk"`        // 备注
}

type Inv101 struct {
	Head Inv101Head   `json:"head"`
	List []Inv101List `json:"list"`
}

type Inv101Xml struct {
	XMLName xml.Name `xml:"Signature"`
	Object  struct {
		Package struct {
			EnvelopInfo common.EnvelopInfo
			DataInfo    struct {
				BusinessData struct {
					DeclareFlag string `xml:"DelcareFlag"` // 这里xml的key和struct的key不一致，海关那个应该是拼写错误
					InvtMessage struct {
						SysId          string
						OperCusRegCode string
						InvtHeadType   Inv101Head
						InvtListType   []Inv101List
					}
				} `xml:"BussinessData"`
			}
		}
	}
}
