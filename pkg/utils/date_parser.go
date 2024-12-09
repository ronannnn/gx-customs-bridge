package utils

import (
	"fmt"
	"time"
)

// ParseHdeApprResultTime 解析回执信息中的处理日期
// 日期格式: 2024-07-26 16:39:17
func ParseHdeApprResultTime(timeStr string) time.Time {
	return ParseCstTimeWithLocation("2006-01-02 15:04:05", timeStr)
}

// ParseClientGeneratedTime 解析客户端生成的时间
// 日期格式: 2024101908250905786339
// 日期格式: 20241019164023018641322
// 长度可能不一致，取前14位20241019164023
func ParseClientGeneratedTime(timeStr string) time.Time {
	if len(timeStr) < 14 {
		return time.Now()
	}
	return ParseCstTimeWithLocation("20060102150405", timeStr[:14])
}

// ParseCstTimeWithLocation 把东八区时间字符串转换为本地时间
func ParseCstTimeWithLocation(fmtStr string, timeStr string) time.Time {
	eightZone := time.FixedZone("CST", 8*3600)
	currentLocation, err := time.LoadLocation("Local")
	if err != nil {
		fmt.Println("Error loading local location:", err)
		currentLocation = time.UTC
	}
	parsedTime, err := time.ParseInLocation(fmtStr, timeStr, eightZone)
	if err != nil {
		return time.Now()
	}
	return parsedTime.In(currentLocation)
}
