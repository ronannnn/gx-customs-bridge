package dec_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/ronannnn/gx-customs-bridge/internal"
	"github.com/ronannnn/gx-customs-bridge/internal/services/customs/dec"
	"github.com/ronannnn/gx-customs-bridge/pkg/customs/decmodels"
	"github.com/stretchr/testify/require"
)

func TestGenDecTmpXml(t *testing.T) {
	service := dec.ProvideDecXmlService(&internal.CustomsCfg{
		SysId:          "Z7",
		OperCusRegCode: "330261A004",
		IcCardNo:       "JJ3G900420543",
		DclErConc:      "贺婷婷",
	})
	xmlBytes, err := service.GenDecTmpXml(decmodels.DecTmp{
		DecHead: decmodels.DecHeadTmp{
			IEFlag:        "I",
			CustomMaster:  "3111",
			IEPort:        "3104",
			Type:          "ML1000",
			AgentCode:     "330261A004",
			AgentName:     "宁波高新物流有限公司",
			CopCode:       "330261A004",
			CopName:       "宁波高新物流有限公司",
			TradeCode:     "330261A004",
			TradeName:     "宁波高新物流有限公司",
			TrafMode:      "2",
			TrafName:      "YM WELLHEAD",
			TradeMode:     "5034",
			GrossWet:      "49744.23",
			DeclTrnRel:    "0",
			TradeAreaCode: "CHN",
			PromiseItmes:  "111999",
			EntryType:     "M",
		},
		DecLists: []decmodels.DecListTmp{
			{
				GNo:                "1",
				CodeTS:             "390230",
				GName:              "丙烯共聚物",
				GQty:               "49.5",
				GUnit:              "070",
				DeclTotal:          "1231",
				TradeCurr:          "CNY",
				OriginCountry:      "CHN",
				DestinationCountry: "CHN",
			},
		},
		DecContainers: []decmodels.DecContainerTmp{
			{
				ContainerId: "ZIMU1234567",
				ContainerMd: "22G1",
				GoodsNo:     "1",
			},
			{
				ContainerId: "ZIMU1234568",
				ContainerMd: "22G1",
				GoodsNo:     "1",
			},
			{
				ContainerId: "ZIMU1234569",
				ContainerMd: "22G1",
				GoodsNo:     "1",
			},
			{
				ContainerId: "ZIMU1234561",
				ContainerMd: "22G1",
				GoodsNo:     "1",
			},
			{
				ContainerId: "ZIMU1234562",
				ContainerMd: "22G1",
				GoodsNo:     "1",
			},
			{
				ContainerId: "ZIMU1234563",
				ContainerMd: "22G1",
				GoodsNo:     "1",
			},
			{
				ContainerId: "ZIMU1234564",
				ContainerMd: "22G1",
				GoodsNo:     "1",
			},
			{
				ContainerId: "ZIMU1234565",
				ContainerMd: "22G1",
				GoodsNo:     "1",
			},
		},
		DecFreeTxt: decmodels.DecFreeTxtTmp{
			VoyNo: "043E",
		},
		DecSign: decmodels.DecSignTmp{
			OperType:    "G",
			ClientSeqNo: "测试报文11",
		},
	})
	require.NoError(t, err)
	os.WriteFile("dec_tmp.xml", xmlBytes, 0644)
	fmt.Printf("xmlBytes: %s\n", string(xmlBytes))
}
