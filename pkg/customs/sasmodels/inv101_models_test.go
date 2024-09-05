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

func TestInv101Head(t *testing.T) {
	translator, err := i18n.New(&i18n.Cfg{BundleDir: "./testdata/"})
	require.NoError(t, err)
	srv := validator.New(translator)

	inv101Head := &sasmodels.Inv101Head{}
	// 读取包含 JSON 数据的文件
	jsonData, err := os.ReadFile("./testdata/inv101_head.json")
	require.NoError(t, err)
	// 解析 JSON 数据到结构体
	err = json.Unmarshal(jsonData, &inv101Head)
	require.NoError(t, err)

	errFields, _ := srv.Check(context.Background(), i18n.LanguageChinese, inv101Head)
	for _, errField := range errFields {
		t.Logf("Error: %s", errField.ErrorMsg)
	}
	require.Equal(t, 1, len(errFields))
	firstNameErrField := errFields[0]
	require.Equal(t, "inv101Head.rltEntryBizopEtpsno", firstNameErrField.ErrorWithNamespace)
	require.Equal(t, "rltEntryBizopEtpsno", firstNameErrField.ErrorField)
	require.Equal(t, "关联报关单境内收发货人编号长度必须是10个字符", firstNameErrField.ErrorMsg)
}

func TestInv101List(t *testing.T) {
	translator, err := i18n.New(&i18n.Cfg{BundleDir: "./testdata/"})
	require.NoError(t, err)
	srv := validator.New(translator)

	inv101List := &sasmodels.Inv101List{}
	// 读取包含 JSON 数据的文件
	jsonData, err := os.ReadFile("./testdata/inv101_list.json")
	require.NoError(t, err)
	// 解析 JSON 数据到结构体
	err = json.Unmarshal(jsonData, &inv101List)
	require.NoError(t, err)

	errFields, _ := srv.Check(context.Background(), i18n.LanguageChinese, inv101List)
	for _, errField := range errFields {
		t.Logf("Error: %s", errField.ErrorMsg)
	}
	require.Equal(t, 0, len(errFields))
}
