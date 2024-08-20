package common

import (
	"path/filepath"
)

const OutBoxDirName = "OutBox"
const SentBoxDirName = "SentBox"
const FailBoxDirName = "FailBox"
const InBoxDirName = "InBox"

const HandledFilesDirName = "Gx"
const FilesCannotParseDirName = "FilesCannotParse"
const FilesCannotUploadDirName = "FilesCannotUpload"

type FilepathHandler struct {
	ImpPath string
	BizType string
}

func NewFilepathHandler(impPath, bizType string) *FilepathHandler {
	return &FilepathHandler{
		ImpPath: impPath,
		BizType: bizType,
	}
}

func (fph *FilepathHandler) GenPath(elem ...string) string {
	var elems = []string{fph.ImpPath, fph.BizType}
	elems = append(elems, elem...)
	result := filepath.Join(elems...)
	return result
}

// 海关本身的文件夹结构
func (fph *FilepathHandler) GenOutBoxPath(ele ...string) string {
	var elems = []string{OutBoxDirName}
	elems = append(elems, ele...)
	return fph.GenPath(elems...)
}
func (fph *FilepathHandler) GenInBoxPath(ele ...string) string {
	var elems = []string{InBoxDirName}
	elems = append(elems, ele...)
	return fph.GenPath(elems...)
}
func (fph *FilepathHandler) GenSentBoxPath(ele ...string) string {
	var elems = []string{SentBoxDirName}
	elems = append(elems, ele...)
	return fph.GenPath(elems...)
}
func (fph *FilepathHandler) GenFailBoxPath(ele ...string) string {
	var elems = []string{FailBoxDirName}
	elems = append(elems, ele...)
	return fph.GenPath(elems...)
}

// 文件处理后的移动到的文件夹结构
func (fph *FilepathHandler) GenHandledPath(elem ...string) string {
	var elems = []string{HandledFilesDirName}
	elems = append(elems, elem...)
	return fph.GenPath(elems...)
}
func (fph *FilepathHandler) GenHandledInBoxPath(ele ...string) string {
	var elems = []string{InBoxDirName}
	elems = append(elems, ele...)
	return fph.GenHandledPath(elems...)
}
func (fph *FilepathHandler) GenHandledOutBoxPath(ele ...string) string {
	var elems = []string{OutBoxDirName}
	elems = append(elems, ele...)
	return fph.GenHandledPath(elems...)
}
func (fph *FilepathHandler) GenHandledSentBoxPath(ele ...string) string {
	var elems = []string{SentBoxDirName}
	elems = append(elems, ele...)
	return fph.GenHandledPath(elems...)
}
func (fph *FilepathHandler) GenHandledFailBoxPath(ele ...string) string {
	var elems = []string{FailBoxDirName}
	elems = append(elems, ele...)
	return fph.GenHandledPath(elems...)
}
func (fph *FilepathHandler) GenHandledCannotParsePath(ele ...string) string {
	var elems = []string{FilesCannotParseDirName}
	elems = append(elems, ele...)
	return fph.GenHandledPath(elems...)
}
func (fph *FilepathHandler) GenHandledCannotUploadPath(ele ...string) string {
	var elems = []string{FilesCannotUploadDirName}
	elems = append(elems, ele...)
	return fph.GenHandledPath(elems...)
}
