package internal

import "strings"

func GetFilenamePrefix(filename string) string {
	return filename[:strings.LastIndex(filename, ".")]
}

func GetFilenameSuffix(filename string) string {
	return filename[strings.LastIndex(filename, ".")+1:]
}
