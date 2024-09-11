package pkg

import (
	"fmt"
	"strconv"
	"time"
)

func DateFromTimestamp(timestamp int64) string {
	// Convert timestamp to time.Time
	timeValue := time.Unix(timestamp, 0)

	// Format time in a human-readable layout
	dateString := timeValue.Format("2006-01-02 15:04:05")

	// Print the formatted date
	//fmt.Println("Formatted Date:", dateString)
	return dateString
}

func TimeNowUnixStr() string {
	return strconv.FormatInt(TimeNow().Unix(), 10)
}

// 获取中国时区当前时间
func TimeNow() time.Time {
	var cstSh, _ = time.LoadLocation("Asia/Shanghai") //上海
	return time.Now().In(cstSh)
}

// 10位时间戳转为time
func TimeFromUnix(unix int64) time.Time {
	return time.Unix(unix, 0)
}

// time转为10位时间戳
func UnixFromTime(t time.Time) int64 {
	return t.Unix()
}

func TimeFormatISO(isoTime string) string {
	// 解析时间字符串
	t, err := time.Parse(time.RFC3339, isoTime)
	if err != nil {
		fmt.Println("[TimeFromUnix] Error:", err)
		return ""
	}

	// 格式化时间为所需格式
	formattedTime := t.Format(time.DateTime)
	return formattedTime
}

// 获取n天的整天时间段 offsetDays 为 -1 表示昨天，-30 表示30天前，0 表示今天，1 表示明天
func GetStartAndEndOfDay(offsetDays int) (time.Time, time.Time) {
	now := time.Now().AddDate(0, 0, offsetDays)
	// 获取偏移日期的开始时间 (0点0分0秒)
	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	// 获取偏移日期的结束时间 (23点59分59秒)
	endOfDay := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, now.Location())
	return startOfDay, endOfDay
}

func GetStartAndEndOfDayHour(offsetDays, hour int) time.Time {
	now := time.Now().AddDate(0, 0, offsetDays)
	// 获取偏移日期的开始时间 (0点0分0秒)
	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), hour, 0, 0, 0, now.Location())
	return startOfDay
}

func FormatMSTTime(dateTimeStr, layout string) string {
	// 定义输入时间字符串的格式
	const inputLayout = "2006-01-02 15:04:05.999999 -0700 MST"
	// 解析时间字符串
	t, err := time.Parse(inputLayout, dateTimeStr)
	if err != nil {
		return ""
	}
	// 格式化时间
	formattedTime := t.Format(layout)
	return formattedTime
}

// FormatHour 将整数值转换为时间格式
func FormatHour(hour int) string {
	t := time.Date(0, 1, 1, hour, 0, 0, 0, time.Now().Location())
	return t.Format(time.TimeOnly)
}

// GetTimeRange 返回从今天往前n天的开始时间和结束时间
func GetTimeRange(n int) (time.Time, time.Time) {
	// 结束时间为当前时间
	endTime := time.Now()
	// 开始时间为n天前的时间
	startTime := endTime.AddDate(0, 0, -n)
	return startTime, endTime
}

// GetTimeRangeZero 返回从今天往前n天的0点开始时间和0点结束时间
func GetTimeRangeZero(n int) (time.Time, time.Time) {
	// 获取当前时间的日期部分，并将时间部分设为零点
	endTime := time.Now().Truncate(24 * time.Hour)
	// 计算开始时间，为n天前的零点
	startTime := endTime.AddDate(0, 0, -n)
	return startTime, endTime
}

//func GetTimeRangeHours(n int) (time.Time, time.Time) {
//	// 获取当前时间的日期部分，并将时间设为当天的12点
//	endTime := time.Now().Truncate(24*time.Hour).AddDate(0, 0, 0).Add(12 * time.Hour)
//
//	// 计算开始时间，为n天前的10点
//	startTime := endTime.AddDate(0, 0, -n).Add(-2 * time.Hour)
//
//	return startTime, endTime
//}

func GetTimeRangeHours(n int) (time.Time, time.Time) {
	// 获取当前日期的零点，并减去n天，得到n天前的零点
	baseTime := time.Now().Truncate(24*time.Hour).AddDate(0, 0, -n)

	// 设置开始时间为n天前的10点
	startTime := baseTime.Add(10 * time.Hour)

	// 设置结束时间为n天前的12点
	endTime := baseTime.Add(12 * time.Hour)

	return startTime, endTime
}
