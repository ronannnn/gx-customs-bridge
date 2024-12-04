package utils

import "time"

// ParseHdeApprResultDate 解析回执信息中的处理日期
// 日期格式: 2024-07-26 16:39:17
func ParseHdeApprResultDate(date string) time.Time {
	parsedTime, err := time.Parse("2006-01-02 15:04:05", date)
	if err != nil {
		return time.Now()
	}
	return parsedTime
}
