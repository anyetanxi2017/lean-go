package utils

import (
	"fmt"
	"strconv"
	"time"
)

const (
	SecondsOfHour  = time.Second * 60 * 60
	SecondsOfDay   = SecondsOfHour * 24
	SecondsOfWeek  = SecondsOfDay * 7
	SecondsOfMonth = SecondsOfDay * 30
)

type TimeUtil struct {
}

type Datetime time.Time

// 日期格式化
func (t Datetime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(t).Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}

// 获取某天的开始结束时间
func (TimeUtil) GetTimeBegin(date time.Time) time.Time {
	timeStr := date.Format("2006-01-02")
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr+" 00:00:00", time.Local)
	return t
}

// 获取某天的最后结束时间
func (TimeUtil) GetTimeEnd(date time.Time) time.Time {
	timeStr := date.Format("2006-01-02")
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr+" 23:59:59", time.Local)
	return t
}

//获取传入的时间所在月份的第一天，即某月第一天的0点。如传入time.Now(), 返回当前月份的第一天0点时间。
func (t TimeUtil) GetBeginDateOfMonth(d time.Time) time.Time {
	d = d.AddDate(0, 0, -d.Day()+1)
	return t.GetZeroTime(d)
}

//获取传入的时间所在月份的最后一天，即某月最后一天的0点。如传入time.Now(), 返回当前月份的最后一天0点时间。
func (t TimeUtil) GetEndDateOfMonth(d time.Time) time.Time {
	return t.GetBeginDateOfMonth(d).AddDate(0, 1, -1)
}

//获取某一天的0点时间
func (TimeUtil) GetZeroTime(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
}

// 毫秒转时间
func (TimeUtil) MsToTime(ms string) (time.Time, error) {
	msInt, err := strconv.ParseInt(ms, 10, 64)
	if err != nil {
		return time.Time{}, err
	}

	tm := time.Unix(0, msInt*int64(time.Millisecond))

	fmt.Println(tm.Format("2006-02-01 15:04:05.000"))

	return tm, nil
}

//当前时间是否在指定范围内
//参数为时间字符串，格式为"时:分:秒"
func (TimeUtil) IsNowInTimeRange(startTimeStr, endTimeStr string) bool {
	//当前时间
	now := time.Now()
	//统一日期
	format := now.Format("2006-01-02")
	//转换为time类型需要的格式
	layout := "2006-01-02 15:04:05"
	//将开始时间拼接“年-月-日 ”转换为time类型
	timeStart, _ := time.ParseInLocation(layout, format+" "+startTimeStr, time.Local)
	//将结束时间拼接“年-月-日 ”转换为time类型
	timeEnd, _ := time.ParseInLocation(layout, format+" "+endTimeStr, time.Local)
	//使用time的Before和After方法，判断当前时间是否在参数的时间范围
	return now.Before(timeEnd) && now.After(timeStart)
}

// 时间格式为：202101040001
func (t TimeUtil) IsToday(timeStr string) (ok bool) {
	s := timeStr[0 : len(timeStr)-4]
	if parse, err := time.ParseInLocation("20060102", s, time.Local); err != nil {
		panic(err)
	} else {
		today := t.GetTimeBegin(time.Now())
		return parse == today
	}
}
