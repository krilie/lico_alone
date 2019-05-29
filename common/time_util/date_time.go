package time_util

import "time"

// 本月开始时间
func GetBeijingMonthStartTime(time time.Time) time.Time {

	time = time.AddDate(0, 0, -time.Day()+1).In(BeijingZone)
	return GetBeijingZeroTime(time).In(BeijingZone)
}

// 获取传入的时间所在月份的最后一天，即某月最后一天的0点。如传入time.Now(), 返回当前月份的最后一天0点时间。
func GetBeijingLastDateOfMonth(d time.Time) time.Time {
	return GetBeijingMonthStartTime(d).AddDate(0, 1, -1).In(BeijingZone)
}

// 获取零点时间
func GetBeijingZeroTime(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location()).In(BeijingZone)
}
