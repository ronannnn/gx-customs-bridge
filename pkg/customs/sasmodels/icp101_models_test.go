package sasmodels_test

import (
	"context"
	"encoding/json"
	"os"
	"testing"

	"github.com/ronannnn/gx-customs-bridge/pkg/customs/sasmodels"
	"github.com/ronannnn/infra/i18n"
	"github.com/ronannnn/infra/validator"
	"github.com/stretchr/testify/require"
)

func TestIcp101Head(t *testing.T) {
	translator, err := i18n.New(&i18n.Cfg{BundleDir: "./testdata/"})
	require.NoError(t, err)
	srv := validator.New(translator)

	icp101Head := &sasmodels.Icp101Head{}
	// 读取包含 JSON 数据的文件
	jsonData, err := os.ReadFile("./testdata/icp101_head.json")
	require.NoError(t, err)
	// 解析 JSON 数据到结构体
	err = json.Unmarshal(jsonData, &icp101Head)
	require.NoError(t, err)

	errFields, _ := srv.Check(context.Background(), i18n.LanguageChinese, icp101Head)
	require.Equal(t, 0, len(errFields))
}

func TestIcp101List(t *testing.T) {
	translator, err := i18n.New(&i18n.Cfg{BundleDir: "./testdata/"})
	require.NoError(t, err)
	srv := validator.New(translator)

	icp101RltList := &sasmodels.Icp101RltList{}
	// 读取包含 JSON 数据的文件
	jsonData, err := os.ReadFile("./testdata/icp101_rlt_list.json")
	require.NoError(t, err)
	// 解析 JSON 数据到结构体
	err = json.Unmarshal(jsonData, &icp101RltList)
	require.NoError(t, err)

	errFields, _ := srv.Check(context.Background(), i18n.LanguageChinese, icp101RltList)
	require.Equal(t, 0, len(errFields))
}
