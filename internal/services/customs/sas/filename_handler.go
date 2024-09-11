package sas

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

type UploadType string

const (
	UploadTypeInv101 UploadType = "INV101"
	UploadTypeSas121 UploadType = "SAS121"
	UploadTypeIcp101 UploadType = "ICP101"
)

func CheckIfUploadTypeValid(uploadType string) error {
	if uploadType != string(UploadTypeInv101) &&
		uploadType != string(UploadTypeSas121) &&
		uploadType != string(UploadTypeIcp101) {
		return fmt.Errorf("无效uploadType: %s, 必须是'INV101'或'SAS121'或'ICP101'", uploadType)
	}
	return nil
}

type ImpexpMarkcd string

const (
	ImpexpMarkcdI ImpexpMarkcd = "I"
	ImpexpMarkcdE ImpexpMarkcd = "E"
	ImpexpMarkcdN ImpexpMarkcd = "N"
)

func CheckIfImpexpMarkcdValid(impexpMarkcd string) error {
	if impexpMarkcd != string(ImpexpMarkcdI) &&
		impexpMarkcd != string(ImpexpMarkcdE) &&
		impexpMarkcd != string(ImpexpMarkcdN) {
		return fmt.Errorf("无效impexpMarkcd: %s, 必须是'I'或'E'或'N'", impexpMarkcd)
	}
	return nil
}

// SasFilenamePart 用于创建和解析Sas文件名
type FilenameParts struct {
	SuccessedOrFailed SuccessedOrFailedType // Successed/Failed(海关)
	UploadType        UploadType            // INV101/SAS121(Sys)
	ImpexpMarkcd      ImpexpMarkcd          // impexpMarkcd(Sys)
	EtpsInnerInvtNo   string                // 企业内部编号(Sys)
	RetryTimes        int                   // 重试次数(Sys) e.g. t-1, t-2
	Timestamp         string                // 海关客户端打上的时间戳(海关)
}

func (s *FilenameParts) GenOutBoxFilenamePrefix() string {
	return fmt.Sprintf("%s_%s_%s_t-%d", s.UploadType, s.ImpexpMarkcd, s.EtpsInnerInvtNo, s.RetryTimes)
}

func (s *FilenameParts) GenOutBoxFilename(ext string) string {
	return fmt.Sprintf("%s.%s", s.GenOutBoxFilenamePrefix(), ext)
}

func NewSasFilenameParts(uploadType UploadType, impexpMarkcd *string, etpsInnerInvtNo *string) (parts FilenameParts, err error) {
	if impexpMarkcd == nil {
		err = fmt.Errorf("impexpMarkcd不能为nil")
		return
	}
	if etpsInnerInvtNo == nil {
		err = fmt.Errorf("etpsInnerInvtNo不能为nil")
		return
	}
	if err = CheckIfUploadTypeValid(string(uploadType)); err != nil {
		return
	}
	if err = CheckIfImpexpMarkcdValid(*impexpMarkcd); err != nil {
		return
	}
	parts = FilenameParts{
		UploadType:      UploadType(uploadType),
		ImpexpMarkcd:    ImpexpMarkcd(*impexpMarkcd),
		EtpsInnerInvtNo: *etpsInnerInvtNo,
		RetryTimes:      1,
	}
	return
}

func ParseSasFilename(filename string) (parts FilenameParts, err error) {
	// get filename prefix
	filenameWithoutParentDir := filepath.Base(filename)
	filenamePrefix := internal.GetFilenamePrefix(filenameWithoutParentDir)
	splittedPrefix := strings.Split(filenamePrefix, "_")

	// validate filename
	if len(splittedPrefix) != 6 {
		err = fmt.Errorf("无效文件名: %s, 以下划线分隔，文件名有且仅有六个部分, 该文件名有%d个", filename, len(splittedPrefix))
		return
	}
	if splittedPrefix[0] != string(SuccessedOrFailedTypeSuccessed) &&
		splittedPrefix[0] != string(SuccessedOrFailedTypeFailed) {
		err = fmt.Errorf("无效文件名: %s, 第一部分必须是'Successed'或'Failed'", filename)
		return
	}
	if err = CheckIfUploadTypeValid(string(splittedPrefix[1])); err != nil {
		return
	}
	if err = CheckIfImpexpMarkcdValid(splittedPrefix[2]); err != nil {
		return
	}
	splittedRetryTimes := strings.Split(splittedPrefix[4], "-")
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
		UploadType:        UploadType(splittedPrefix[1]),
		ImpexpMarkcd:      ImpexpMarkcd(splittedPrefix[2]),
		EtpsInnerInvtNo:   splittedPrefix[3],
		RetryTimes:        retryTimes,
		Timestamp:         splittedPrefix[5],
	}
	return
}
