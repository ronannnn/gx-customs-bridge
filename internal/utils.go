package internal

import (
	"archive/zip"
	"bytes"
	"io"
	"path/filepath"
)

func GetFilenamePrefix(filename string) string {
	return filename[:len(filename)-len(filepath.Ext(filename))]
}

func ZipFile(filename string, fileBytes []byte) (zipFileBytes []byte, err error) {
	buf := new(bytes.Buffer)
	writer := zip.NewWriter(buf)

	// 将文件数据写入 ZIP 文件
	var file io.Writer
	if file, err = writer.Create(filename); err != nil {
		return
	}
	if _, err = file.Write(fileBytes); err != nil {
		return
	}

	// 关闭 ZIP 文件写入器
	if err = writer.Close(); err != nil {
		return
	}

	zipFileBytes = buf.Bytes()
	return
}
