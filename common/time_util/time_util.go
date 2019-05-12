package time_util

import "time"

//func GetNowUnix() int64 {
//	return time.Now().Unix()
//}
var cstZone *time.Location

func init() {
	cstZone = time.FixedZone("CST", 8*3600) // 东八
}

func GetNowTimeString() string {
	return time.Now().In(cstZone).Format("2006-01-02 15:04:05")
}
func GetNowTimeStringFormat(format string) string {
	return time.Now().In(cstZone).Format(format)
}

func GetTimeNow() *time.Time {
	timeN := time.Now().In(cstZone)
	return &timeN
}

func GetTimeString(unix int64) string {
	return time.Unix(unix, 0).In(cstZone).Format("2006-01-02 15:04:05")
}
