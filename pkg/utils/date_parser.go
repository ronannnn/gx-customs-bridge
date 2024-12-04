package utils

import "time"

// ParseHdeApprResultTime 解析回执信息中的处理日期
// 日期格式: 2024-07-26 16:39:17
func ParseHdeApprResultTime(timeStr string) time.Time {
	parsedTime, err := time.Parse("2006-01-02 15:04:05", timeStr)
	if err != nil {
		return time.Now()
	}
	return parsedTime
}

// ParseClientGeneratedTime 解析客户端生成的时间
// 日期格式: 2024101908250905786339
// 日期格式: 20241019164023018641322
// 长度可能不一致，取前14位20241019164023
func ParseClientGeneratedTime(timeStr string) time.Time {
	if len(timeStr) < 14 {
		return time.Now()
	}
	parsedTime, err := time.Parse("20060102150405", timeStr[:14])
	if err != nil {
		return time.Now()
	}
	return parsedTime
}
