package tools

import "time"

func IsToday(t time.Time) uint {
	return 0
}

func Today() time.Time {
	// 获取当前时间
	now := time.Now()
	// 设置时分秒为0
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
}
