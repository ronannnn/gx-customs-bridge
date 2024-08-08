package internal

import "strings"

func GetFilenamePrefix(filename string) string {
	return filename[:strings.LastIndex(filename, ".")]
}

func GetFilenameSuffix(filename string) string {
	return filename[strings.LastIndex(filename, ".")+1:]
}

func GetFilenameFromPath(filePath string) string {
	return filePath[strings.LastIndex(filePath, "/")+1:]
}
