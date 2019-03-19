package time_util

import "time"

//func GetNowUnix() int64 {
//	return time.Now().Unix()
//}

func GetNowTimeString() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
func GetNowTimeStringFormat(format string) string {
	return time.Now().Format(format)
}

func GetTimeNow() *time.Time {
	timeN := time.Now()
	return &timeN
}
