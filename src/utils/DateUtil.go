package utils

import (
	"time"
)

//当前系统秒
func GetCurrentTimestamp() int32 {
	return int32(time.Now().Unix())
}

//当前系统毫秒
func GetCurrentTimeMillis() int64 {
	return time.Now().UnixNano() / 1e6
}

//当前系统毫秒
func GetCurrentTimeNano() int64 {
	return time.Now().UnixNano() / 1e6
}

func GetDatetimeYmd() string {
	now := time.Now()
	return now.Format("20060102")
}

//时间原点：2006-01-02 15:04:05
func GetDatetimeYmdHis() string {
	now := time.Now()
	return now.Format("20060102150405")
}