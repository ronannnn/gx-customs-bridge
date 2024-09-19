package common

import (
	"path/filepath"

	"github.com/ronannnn/gx-customs-bridge/internal"
)

const OutBoxDirName = "OutBox"
const SentBoxDirName = "SentBox"
const FailBoxDirName = "FailBox"
const InBoxDirName = "InBox"

const HandledFilesDirName = "Gx"
const FilesCannotParseDirName = "FilesCannotParse"
const FilesCannotUploadDirName = "FilesCannotUpload"

type FilepathHandler struct {
	icCardMap map[string]*internal.CustomsIcCard
	BizType   string
}

func NewFilepathHandler(icCardMap map[string]*internal.CustomsIcCard, bizType string) *FilepathHandler {
	return &FilepathHandler{
		icCardMap: icCardMap,
		BizType:   bizType,
	}
}

func (fph *FilepathHandler) GenPath(companyType string, elem ...string) string {
	icCard, ok := fph.icCardMap[companyType]
	if !ok {
		return ""
	}
	var elems = []string{icCard.ImpPath, fph.BizType}
	elems = append(elems, elem...)
	result := filepath.Join(elems...)
	return result
}

// 海关本身的文件夹结构
func (fph *FilepathHandler) GenOutBoxPath(companyType string, ele ...string) string {
	var elems = []string{OutBoxDirName}
	elems = append(elems, ele...)
	return fph.GenPath(companyType, elems...)
}
func (fph *FilepathHandler) GenInBoxPath(companyType string, ele ...string) string {
	var elems = []string{InBoxDirName}
	elems = append(elems, ele...)
	return fph.GenPath(companyType, elems...)
}
func (fph *FilepathHandler) GenSentBoxPath(companyType string, ele ...string) string {
	var elems = []string{SentBoxDirName}
	elems = append(elems, ele...)
	return fph.GenPath(companyType, elems...)
}
func (fph *FilepathHandler) GenFailBoxPath(companyType string, ele ...string) string {
	var elems = []string{FailBoxDirName}
	elems = append(elems, ele...)
	return fph.GenPath(companyType, elems...)
}

// 文件处理后的移动到的文件夹结构
func (fph *FilepathHandler) GenHandledPath(companyType string, elem ...string) string {
	var elems = []string{HandledFilesDirName}
	elems = append(elems, elem...)
	return fph.GenPath(companyType, elems...)
}
func (fph *FilepathHandler) GenHandledInBoxPath(companyType string, ele ...string) string {
	var elems = []string{InBoxDirName}
	elems = append(elems, ele...)
	return fph.GenHandledPath(companyType, elems...)
}
func (fph *FilepathHandler) GenHandledOutBoxPath(companyType string, ele ...string) string {
	var elems = []string{OutBoxDirName}
	elems = append(elems, ele...)
	return fph.GenHandledPath(companyType, elems...)
}
func (fph *FilepathHandler) GenHandledSentBoxPath(companyType string, ele ...string) string {
	var elems = []string{SentBoxDirName}
	elems = append(elems, ele...)
	return fph.GenHandledPath(companyType, elems...)
}
func (fph *FilepathHandler) GenHandledFailBoxPath(companyType string, ele ...string) string {
	var elems = []string{FailBoxDirName}
	elems = append(elems, ele...)
	return fph.GenHandledPath(companyType, elems...)
}
func (fph *FilepathHandler) GenHandledCannotParsePath(companyType string, ele ...string) string {
	var elems = []string{FilesCannotParseDirName}
	elems = append(elems, ele...)
	return fph.GenHandledPath(companyType, elems...)
}
func (fph *FilepathHandler) GenHandledCannotUploadPath(companyType string, ele ...string) string {
	var elems = []string{FilesCannotUploadDirName}
	elems = append(elems, ele...)
	return fph.GenHandledPath(companyType, elems...)
}
