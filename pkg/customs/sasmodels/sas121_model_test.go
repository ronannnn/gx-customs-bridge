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

func TestSas121Head(t *testing.T) {
	translator, err := i18n.New(&i18n.Cfg{BundleDir: "./testdata/"})
	require.NoError(t, err)
	srv := validator.New(translator)

	sas121Head := &sasmodels.Sas121Head{}
	// 读取包含 JSON 数据的文件
	jsonData, err := os.ReadFile("./testdata/sas121_head.json")
	require.NoError(t, err)
	// 解析 JSON 数据到结构体
	err = json.Unmarshal(jsonData, &sas121Head)
	require.NoError(t, err)

	errFields, _ := srv.Check(context.Background(), i18n.LanguageChinese, sas121Head)
	for _, errField := range errFields {
		t.Logf("Error: %s", errField.ErrorMsg)
	}
	require.Equal(t, 0, len(errFields))
}

func TestSas121List(t *testing.T) {
	translator, err := i18n.New(&i18n.Cfg{BundleDir: "./testdata/"})
	require.NoError(t, err)
	srv := validator.New(translator)

	sas121List := &sasmodels.Sas121List{}
	// 读取包含 JSON 数据的文件
	jsonData, err := os.ReadFile("./testdata/sas121_list.json")
	require.NoError(t, err)
	// 解析 JSON 数据到结构体
	err = json.Unmarshal(jsonData, &sas121List)
	require.NoError(t, err)

	errFields, _ := srv.Check(context.Background(), i18n.LanguageChinese, sas121List)
	for _, errField := range errFields {
		t.Logf("Error: %s", errField.ErrorMsg)
	}
	require.Equal(t, 0, len(errFields))
}
