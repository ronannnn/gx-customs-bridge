package dec_test

import (
	"encoding/json"
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
		SysId: "Z7",
		IcCards: []internal.CustomsIcCard{
			{
				Name:           "gxwl",
				OperCusRegCode: "330261A004",
				IcCardNo:       "JJ3G900420543",
				DclErConc:      "贺婷婷",
			},
		},
	})
	var decModel decmodels.Dec
	// 读取包含 JSON 数据的文件
	jsonData, err := os.ReadFile("./test/dec.json")
	require.NoError(t, err)
	// 解析 JSON 数据到结构体
	err = json.Unmarshal(jsonData, &decModel)
	require.NoError(t, err)

	xmlBytes, err := service.GenDecTmpXml(decModel, "G", "gxwl")
	require.NoError(t, err)
	os.WriteFile("dec_tmp.xml", xmlBytes, 0644)
	fmt.Printf("xmlBytes: %s\n", string(xmlBytes))
}
