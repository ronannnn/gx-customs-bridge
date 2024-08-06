package inv

import (
	"encoding/xml"
	"time"

	"github.com/ronannnn/gx-customs-bridge/internal/services/customs/common"
	"github.com/shopspring/decimal"
)

// Inv211Head 核注清单表头
type Inv211Head struct {
	// 非业务相关字段
	ChgTmsCnt *string `json:"chgTmsCnt"` // 变更次数

	// 编号
	InvtPreentNo    *string `json:"invtPreentNo"`    // 预录入统一编号(同INV101中的SeqNo)
	EtpsInnerInvtNo *string `json:"etpsInnerInvtNo"` // 企业内部清单编号
	BondInvtNo      *string `json:"bondInvtNo"`      // 保税清单编号
	PutrecNo        *string `json:"putrecNo"`        // 备案编号，手(账)册编号

	// 清单信息
	DclTypecd      *string `json:"dclTypecd"`      // 申报类型(1-备案申请 2-变更申请 3-删除申请)
	BondInvtTypecd *string `json:"bondInvtTypecd"` // 清单类型代码 同INV101中的InvtType
	InvtStucd      *string `json:"invtStucd"`      // 清单状态代码(1-申报 2-退单 3删单 0-审核通过) 同INV101中的ListStat

	// 核放相关
	InvtIochkptStucd   *string `json:"invtIochkptStucd"`   // 清单进出卡口状态代码
	PassportUsedTypeCd *string `json:"passportUsedTypeCd"` // 核放单生成标志代码
	VrfdedMarkcd       *string `json:"vrfdedMarkcd"`       // 核扣标记代码(0-未核扣 1-预核扣 2-已核扣 3-已核销)
	VrfdedModecd       *string `json:"vrfdedModecd"`       // 核扣方式代码(D:正扣、F:反扣、N：不扣 NF：不扣但检查余量、E：保税仓扣减)
	DuCode             *string `json:"duCode"`             // 核算代码(两位代码)

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
	// 录入单位 201没有
	// InputCreditCode *string `json:"inputCreditCode"` // 社会信用代码
	// InputCode       *string `json:"inputCode"`       // 编号
	// InputName       *string `json:"inputName"`       // 名称

	// 日期
	// InputTime        *string    `json:"inputTime"`        // 录入日期(格式：20240101)
	InvtDclTime      *string    `json:"invtDclTime"`      // 清单申报时间(格式：20240101)
	EntryDclTime     *string    `json:"entryDclTime"`     // 报关单申报时间(格式：20240101)
	PrevdTime        *time.Time `json:"prevdTime"`        // 预核扣时间
	FormalVrfdedTime *time.Time `json:"formalVrfdedTime"` // 正式核扣时间

	// 代码
	ImpexpMarkcd      *string `json:"impexpMarkcd"`      // 进出口标记代码(I-进口 E-出口)
	MtpckEndprdMarkcd *string `json:"mtpckEndprdMarkcd"` // 料件成品标记代码(I-料件 E-成品)
	SupvModecd        *string `json:"supvModecd"`        // 监管方式代码
	TrspModecd        *string `json:"trspModecd"`        // 运输方式代码
	ImpexpPortcd      *string `json:"impexpPortcd"`      // 进出境关别代码(进出口口岸代码)
	DclPlcCuscd       *string `json:"dclPlcCuscd"`       // 主管海关(申报地关区代码)
	StshipTrsarvNatcd *string `json:"stshipTrsarvNatcd"` // 起运/运抵国(地区)

	// 报关相关
	DclcusFlag *string `json:"dclcusFlag"` // 是否报关标志(1-报关 2-非报关)
	// GenDecFlag        *string `json:"genDecFlag"`        // 是否生成报关单:1-生成 2-不生成
	// DecType           *string `json:"decType"`           // 报关单类型
	DclcusTypecd      *string `json:"dclcusTypecd"`      // 报关类型代码
	NeedEntryModified *string `json:"needEntryModified"` // 报关单同步修改标志
	EntryStucd        *string `json:"entryStucd"`        // 报关状态(0：未放行，1：已放行) 该类型清单满足两个条件才能核扣：报关单被放行+货物全部过卡

	// 对应报关相关
	EntryNo *string `json:"entryNo"` // 对应报关单编号
	// 对应报关单申报单位 201没有
	// CorrEntryDclEtpsSccd *string `json:"corrEntryDclEtpsSccd"` // 社会信用代码
	// CorrEntryDclEtpsNo   *string `json:"corrEntryDclEtpsNo"`   // 编号
	// CorrEntryDclEtpsNm   *string `json:"corrEntryDclEtpsNm"`   // 名称

	// 关联报关相关
	RltInvtNo   *string `json:"rltInvtNo"`   // 关联清单编号
	RltPutrecNo *string `json:"rltPutrecNo"` // 关联备案编号，关联手(账)册编号
	RltEntryNo  *string `json:"rltEntryNo"`  // 关联报关单编号
	// 关联报关单境内收发货人
	RltEntryBizopEtpsSccd *string `json:"rltEntryBizopEtpsSccd"` // 社会信用代码
	RltEntryBizopEtpsno   *string `json:"rltEntryBizopEtpsno"`   // 编号
	RltEntryBizopEtpsNm   *string `json:"rltEntryBizopEtpsNm"`   // 名称
	// 关联报关单生产销售(消费使用)单位
	// RltEntryRvsngdEtpsSccd *string `json:"rltEntryRvsngdEtpsSccd"` // 社会信用代码
	// RltEntryRcvgdEtpsno    *string `json:"rltEntryRcvgdEtpsno"`    // 编号
	// RltEntryRcvgdEtpsNm    *string `json:"rltEntryRcvgdEtpsNm"`    // 名称
	// 关联报关单申报单位
	// RltEntryDclEtpsSccd *string `json:"rltEntryDclEtpsSccd"` // 社会信用代码
	// RltEntryDclEtpsno   *string `json:"rltEntryDclEtpsno"`   // 编号
	// RltEntryDclEtpsNm   *string `json:"rltEntryDclEtpsNm"`   // 名称

	// others
	// IcCardNo  *string `json:"icCardNo"`  // 申报人IC卡号
	ApplyNo   *string `json:"applyNo"`   // 申请编号
	Rmk       *string `json:"rmk"`       // 备注
	LevyBlAmt *string `json:"levyBlAmt"` // 计征金额

	Param1 *string `json:"param1"` // 参数1
	Param2 *string `json:"param2"` // 参数2
	Param3 *string `json:"param3"` // 参数3
	Param4 *string `json:"param4"` // 参数4
}

// Inv211List 核注清单表体
type Inv211List struct {
	// 非业务相关字段
	ChgTmsCnt *string `json:"chgTmsCnt"` // 变更次数

	// 编号
	BwsNo   *string `json:"bwsNo"`   // 帐册编号(和表头的PutrecNo一致)
	InvtNo  *string `json:"invtNo"`  // 记账清单编号(同QD号)
	InvtGNo *string `json:"invtGNo"` // 记账清单商品序号(同商品序号)

	GdsSeqno         *string `json:"gdsSeqno"`         // 商品序号(同备案序号)
	Gdecd            *string `json:"gdecd"`            // 商品编码
	GdsMtno          *string `json:"gdsMtno"`          // 商品料号
	GdsNm            *string `json:"gdsNm"`            // 商品名称
	GdsSpcfModelDesc *string `json:"gdsSpcfModelDesc"` // 商品规格型号

	DclUnitcd      *string          `json:"dclUnitcd"`      // 申报计量单位代码
	InQty          *decimal.Decimal `json:"inQty"`          // 入仓数量
	LawfUnitcd     *string          `json:"lawfUnitcd"`     // 法定计量单位代码
	InLawfQty      *decimal.Decimal `json:"inLawfQty"`      // 入仓法定数量
	SecdLawfUnitcd *string          `json:"secdLawfUnitcd"` // 第二法定计量单位代码
	InSecdLawfQty  *decimal.Decimal `json:"inSecdLawfQty"`  // 第二入仓法定数量

	ActlIncQty   *decimal.Decimal `json:"actlIncQty"`   // 实增数量
	ActlRedcQty  *decimal.Decimal `json:"actlRedcQty"`  // 实减数量
	PrevdIncQty  *decimal.Decimal `json:"prevdIncQty"`  // 预增数量
	PrevdRedcQty *decimal.Decimal `json:"prevdRedcQty"` // 预减数量

	DclCurrcd  *string          `json:"dclCurrcd"`  // 申报币制代码
	DclUprcAmt *decimal.Decimal `json:"dclUprcAmt"` // 企业申报单价
	AvgPrice   *decimal.Decimal `json:"avgPrice"`   // 平均美元单价 每次总价变化时自动计算：总价/库存
	TotalAmt   *decimal.Decimal `json:"totalAmt"`   // 库存美元总价 每次核注(增减）时，同时核注(增减）金额

	InDate    *string `json:"inDate"`    // 最近入仓(核增)日期
	OutDate   *string `json:"outDate"`   // 最近出仓日期
	LimitDate *string `json:"limitDate"` // 存储(监管)期限

	Natcd         *string `json:"natcd"`         // 原产国（地区）代码
	InType        *string `json:"inType"`        // 设备入区方式代码(1:一线入区、2:二线入区、3:结转入区)
	CusmExeMarkcd *string `json:"cusmExeMarkcd"` // 海关执行标志(1-正常执行 4-暂停进出口 5-暂停进口 6-暂停出口，默认为)

	ModfMarkcd *string `json:"modfMarkcd"` // 修改标志
	Rmk        *string `json:"rmk"`        // 备注
}

type Inv211 struct {
	Head Inv211Head   `json:"head"`
	List []Inv211List `json:"list"`
}

// Inv211Xml 核注清单记账回执
type Inv211Xml struct {
	XMLName     xml.Name `xml:"Package"`
	EnvelopInfo common.EnvelopInfo
	DataInfo    struct {
		PocketInfo   common.PocketInfo
		BusinessData struct {
			Inv211 struct {
				BondInvtBsc Inv211Head
				BwsDt       []Inv211List
			} `xml:"INV211"`
		}
	}
}
