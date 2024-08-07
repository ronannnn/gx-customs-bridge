package inv

import (
	"encoding/xml"

	"github.com/ronannnn/gx-customs-bridge/internal/services/customs/common"
)

// Inv211Head 核注清单表头
type Inv211Head struct {
	// 非业务相关字段
	ChgTmsCnt string `json:"chgTmsCnt" xml:"chgTmsCnt"` // 变更次数

	// 编号
	InvtPreentNo    string `json:"invtPreentNo" xml:"invtPreentNo"`       // 预录入统一编号(同INV101中的SeqNo)
	EtpsInnerInvtNo string `json:"etpsInnerInvtNo" xml:"etpsInnerInvtNo"` // 企业内部清单编号
	BondInvtNo      string `json:"bondInvtNo" xml:"bondInvtNo"`           // 保税清单编号
	PutrecNo        string `json:"putrecNo" xml:"putrecNo"`               // 备案编号，手(账)册编号

	// 清单信息
	DclTypecd      string `json:"dclTypecd" xml:"dclTypecd"`           // 申报类型(1-备案申请 2-变更申请 3-删除申请)
	BondInvtTypecd string `json:"bondInvtTypecd" xml:"bondInvtTypecd"` // 清单类型代码 同INV101中的InvtType
	InvtStucd      string `json:"invtStucd" xml:"invtStucd"`           // 清单状态代码(1-申报 2-退单 3删单 0-审核通过) 同INV101中的ListStat

	// 核放相关
	InvtIochkptStucd   string `json:"invtIochkptStucd" xml:"invtIochkptStucd"`     // 清单进出卡口状态代码
	PassportUsedTypeCd string `json:"passportUsedTypeCd" xml:"passportUsedTypeCd"` // 核放单生成标志代码
	VrfdedMarkcd       string `json:"vrfdedMarkcd" xml:"vrfdedMarkcd"`             // 核扣标记代码(0-未核扣 1-预核扣 2-已核扣 3-已核销)
	VrfdedModecd       string `json:"vrfdedModecd" xml:"vrfdedModecd"`             // 核扣方式代码(D:正扣、F:反扣、N：不扣 NF：不扣但检查余量、E：保税仓扣减)
	DuCode             string `json:"duCode" xml:"duCode"`                         // 核算代码(两位代码)

	// 经营企业
	BizopEtpsSccd string `json:"bizopEtpsSccd" xml:"bizopEtpsSccd"` // 社会信用代码
	BizopEtpsno   string `json:"bizopEtpsno" xml:"bizopEtpsno"`     // 编号
	BizopEtpsNm   string `json:"bizopEtpsNm" xml:"bizopEtpsNm"`     // 名称
	// 申报企业
	DclEtpsSccd string `json:"dclEtpsSccd" xml:"dclEtpsSccd"` // 社会信用代码
	DclEtpsno   string `json:"dclEtpsno" xml:"dclEtpsno"`     // 编号
	DclEtpsNm   string `json:"dclEtpsNm" xml:"dclEtpsNm"`     // 名称
	// 收发货企业
	RvsngdEtpsSccd string `json:"rvsngdEtpsSccd" xml:"rvsngdEtpsSccd"` // 社会信用代码
	RcvgdEtpsno    string `json:"rcvgdEtpsno" xml:"rcvgdEtpsno"`       // 编号
	RcvgdEtpsNm    string `json:"rcvgdEtpsNm" xml:"rcvgdEtpsNm"`       // 名称

	// 日期
	InvtDclTime      string `json:"invtDclTime" xml:"invtDclTime"`           // 清单申报时间(格式：20240101)
	EntryDclTime     string `json:"entryDclTime" xml:"entryDclTime"`         // 报关单申报时间(格式：20240101)
	PrevdTime        string `json:"prevdTime" xml:"prevdTime"`               // 预核扣时间
	FormalVrfdedTime string `json:"formalVrfdedTime" xml:"formalVrfdedTime"` // 正式核扣时间

	// 代码
	ImpexpMarkcd      string `json:"impexpMarkcd" xml:"impexpMarkcd"`           // 进出口标记代码(I-进口 E-出口)
	MtpckEndprdMarkcd string `json:"mtpckEndprdMarkcd" xml:"mtpckEndprdMarkcd"` // 料件成品标记代码(I-料件 E-成品)
	SupvModecd        string `json:"supvModecd" xml:"supvModecd"`               // 监管方式代码
	TrspModecd        string `json:"trspModecd" xml:"trspModecd"`               // 运输方式代码
	ImpexpPortcd      string `json:"impexpPortcd" xml:"impexpPortcd"`           // 进出境关别代码(进出口口岸代码)
	DclPlcCuscd       string `json:"dclPlcCuscd" xml:"dclPlcCuscd"`             // 主管海关(申报地关区代码)
	StshipTrsarvNatcd string `json:"stshipTrsarvNatcd" xml:"stshipTrsarvNatcd"` // 起运/运抵国(地区)

	// 报关相关
	DclcusFlag        string `json:"dclcusFlag" xml:"dclcusFlag"`               // 是否报关标志(1-报关 2-非报关)
	DclcusTypecd      string `json:"dclcusTypecd" xml:"dclcusTypecd"`           // 报关类型代码
	NeedEntryModified string `json:"needEntryModified" xml:"needEntryModified"` // 报关单同步修改标志
	EntryStucd        string `json:"entryStucd" xml:"entryStucd"`               // 报关状态(0：未放行，1：已放行) 该类型清单满足两个条件才能核扣：报关单被放行+货物全部过卡

	// 对应报关相关
	EntryNo string `json:"entryNo" xml:"entryNo"` // 对应报关单编号

	// 关联报关相关
	RltInvtNo   string `json:"rltInvtNo" xml:"rltInvtNo"`     // 关联清单编号
	RltPutrecNo string `json:"rltPutrecNo" xml:"rltPutrecNo"` // 关联备案编号，关联手(账)册编号
	RltEntryNo  string `json:"rltEntryNo" xml:"rltEntryNo"`   // 关联报关单编号
	// 关联报关单境内收发货人
	RltEntryBizopEtpsSccd string `json:"rltEntryBizopEtpsSccd" xml:"rltEntryBizopEtpsSccd"` // 社会信用代码
	RltEntryBizopEtpsno   string `json:"rltEntryBizopEtpsno" xml:"rltEntryBizopEtpsno"`     // 编号
	RltEntryBizopEtpsNm   string `json:"rltEntryBizopEtpsNm" xml:"rltEntryBizopEtpsNm"`     // 名称

	// others
	ApplyNo   string `json:"applyNo" xml:"applyNo"`     // 申请编号
	Rmk       string `json:"rmk" xml:"rmk"`             // 备注
	LevyBlAmt string `json:"levyBlAmt" xml:"levyBlAmt"` // 计征金额

	Param1 string `json:"param1" xml:"param1"` // 参数1
	Param2 string `json:"param2" xml:"param2"` // 参数2
	Param3 string `json:"param3" xml:"param3"` // 参数3
	Param4 string `json:"param4" xml:"param4"` // 参数4
}

// Inv211List 核注清单表体
type Inv211List struct {
	// 非业务相关字段
	ChgTmsCnt string `json:"chgTmsCnt" xml:"chgTmsCnt"` // 变更次数

	// 编号
	BwsNo   string `json:"bwsNo" xml:"bwsNo"`     // 帐册编号(和表头的PutrecNo一致)
	InvtNo  string `json:"invtNo" xml:"invtNo"`   // 记账清单编号(同QD号)
	InvtGNo string `json:"invtGNo" xml:"invtGNo"` // 记账清单商品序号(同商品序号)

	GdsSeqno         string `json:"gdsSeqno" xml:"gdsSeqno"`                 // 商品序号(同备案序号)
	Gdecd            string `json:"gdecd" xml:"gdecd"`                       // 商品编码
	GdsMtno          string `json:"gdsMtno" xml:"gdsMtno"`                   // 商品料号
	GdsNm            string `json:"gdsNm" xml:"gdsNm"`                       // 商品名称
	GdsSpcfModelDesc string `json:"gdsSpcfModelDesc" xml:"gdsSpcfModelDesc"` // 商品规格型号

	DclUnitcd      string `json:"dclUnitcd" xml:"dclUnitcd"`           // 申报计量单位代码
	InQty          string `json:"inQty" xml:"inQty"`                   // 入仓数量
	LawfUnitcd     string `json:"lawfUnitcd" xml:"lawfUnitcd"`         // 法定计量单位代码
	InLawfQty      string `json:"inLawfQty" xml:"inLawfQty"`           // 入仓法定数量
	SecdLawfUnitcd string `json:"secdLawfUnitcd" xml:"secdLawfUnitcd"` // 第二法定计量单位代码
	InSecdLawfQty  string `json:"inSecdLawfQty" xml:"inSecdLawfQty"`   // 第二入仓法定数量

	ActlIncQty   string `json:"actlIncQty" xml:"actlIncQty"`     // 实增数量
	ActlRedcQty  string `json:"actlRedcQty" xml:"actlRedcQty"`   // 实减数量
	PrevdIncQty  string `json:"prevdIncQty" xml:"prevdIncQty"`   // 预增数量
	PrevdRedcQty string `json:"prevdRedcQty" xml:"prevdRedcQty"` // 预减数量

	DclCurrcd  string `json:"dclCurrcd" xml:"dclCurrcd"`   // 申报币制代码
	DclUprcAmt string `json:"dclUprcAmt" xml:"dclUprcAmt"` // 企业申报单价
	AvgPrice   string `json:"avgPrice" xml:"avgPrice"`     // 平均美元单价 每次总价变化时自动计算：总价/库存
	TotalAmt   string `json:"totalAmt" xml:"totalAmt"`     // 库存美元总价 每次核注(增减）时，同时核注(增减）金额

	InDate    string `json:"inDate" xml:"inDate"`       // 最近入仓(核增)日期
	OutDate   string `json:"outDate" xml:"outDate"`     // 最近出仓日期
	LimitDate string `json:"limitDate" xml:"limitDate"` // 存储(监管)期限

	Natcd         string `json:"natcd" xml:"natcd"`                 // 原产国（地区）代码
	InType        string `json:"inType" xml:"inType"`               // 设备入区方式代码(1:一线入区、2:二线入区、3:结转入区)
	CusmExeMarkcd string `json:"cusmExeMarkcd" xml:"cusmExeMarkcd"` // 海关执行标志(1-正常执行 4-暂停进出口 5-暂停进口 6-暂停出口，默认为)

	ModfMarkcd string `json:"modfMarkcd" xml:"modfMarkcd"` // 修改标志
	Rmk        string `json:"rmk" xml:"rmk"`               // 备注
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
		} `xml:"BussinessData"`
	}
}
