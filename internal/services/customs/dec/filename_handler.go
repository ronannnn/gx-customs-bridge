package dec

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/ronannnn/gx-customs-bridge/internal"
)

type SuccessedOrFailedType string

const (
	SuccessedOrFailedTypeSuccessed SuccessedOrFailedType = "Successed"
	SuccessedOrFailedTypeFailed    SuccessedOrFailedType = "Failed"
)

type ImpexpMarkcd string

const (
	ImpexpMarkcdI ImpexpMarkcd = "I"
	ImpexpMarkcdE ImpexpMarkcd = "E"
)

func CheckIfImpexpMarkcdValid(impexpMarkcd string) error {
	if impexpMarkcd != string(ImpexpMarkcdI) &&
		impexpMarkcd != string(ImpexpMarkcdE) {
		return fmt.Errorf("无效impexpMarkcd: %s, 必须是'I'或'E'", impexpMarkcd)
	}
	return nil
}

// DecFilenamePart 用于创建和解析dec文件名
type FilenameParts struct {
	SuccessedOrFailed SuccessedOrFailedType // Successed/Failed(海关)
	ImpexpMarkcd      ImpexpMarkcd          // impexpMarkcd(Sys)
	ClientSeqNo       string                // 客户端报关单编号(Sys)
	RetryTimes        int                   // 重试次数(Sys) e.g. t-1, t-2
	Timestamp         string                // 海关客户端打上的时间戳(海关)
}

func (s *FilenameParts) GenOutBoxFilenamePrefix() string {
	return fmt.Sprintf("%s_%s_t-%d", s.ImpexpMarkcd, s.ClientSeqNo, s.RetryTimes)
}

func (s *FilenameParts) GenOutBoxFilename(ext string) string {
	return fmt.Sprintf("%s.%s", s.GenOutBoxFilenamePrefix(), ext)
}

func NewDecFilenameParts(impexpMarkcd *string, clientSeqNo *string) (parts FilenameParts, err error) {
	if impexpMarkcd == nil {
		err = fmt.Errorf("impexpMarkcd不能为nil")
		return
	}
	if clientSeqNo == nil {
		err = fmt.Errorf("uuid不能为nil")
		return
	}
	if err = CheckIfImpexpMarkcdValid(*impexpMarkcd); err != nil {
		return
	}
	parts = FilenameParts{
		ImpexpMarkcd: ImpexpMarkcd(*impexpMarkcd),
		ClientSeqNo:  *clientSeqNo,
		RetryTimes:   1,
	}
	return
}

func ParseDecFilename(filename string) (parts FilenameParts, err error) {
	// get filename prefix
	filenameWithoutParentDir := filepath.Base(filename)
	filenamePrefix := internal.GetFilenamePrefix(filenameWithoutParentDir)
	splittedPrefix := strings.Split(filenamePrefix, "_")

	// validate filename
	if len(splittedPrefix) != 5 {
		err = fmt.Errorf("无效文件名: %s, 以下划线分隔，文件名有且仅有5个部分, 该文件名有%d个", filename, len(splittedPrefix))
		return
	}
	if splittedPrefix[0] != string(SuccessedOrFailedTypeSuccessed) &&
		splittedPrefix[0] != string(SuccessedOrFailedTypeFailed) {
		err = fmt.Errorf("无效文件名: %s, 第一部分必须是'Successed'或'Failed'", filename)
		return
	}
	if err = CheckIfImpexpMarkcdValid(splittedPrefix[1]); err != nil {
		return
	}
	splittedRetryTimes := strings.Split(splittedPrefix[3], "-")
	if len(splittedRetryTimes) != 2 || splittedRetryTimes[0] != "t" {
		err = fmt.Errorf("无效文件名: %s, 第五部分必须是't-{n}'格式", filename)
		return
	}
	var retryTimes int
	if retryTimes, err = strconv.Atoi(splittedRetryTimes[1]); err != nil {
		return
	}

	parts = FilenameParts{
		SuccessedOrFailed: SuccessedOrFailedType(splittedPrefix[0]),
		ImpexpMarkcd:      ImpexpMarkcd(splittedPrefix[1]),
		ClientSeqNo:       splittedPrefix[2],
		RetryTimes:        retryTimes,
		Timestamp:         splittedPrefix[4],
	}
	return
}
