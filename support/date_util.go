package support

/**
* @author zhangsheng
* @date 2019/6/13
 */

import (
	"fmt"
	"strings"
	"time"
)

type DateStyle string

const (
	MM_DD                   = "MM-dd"
	YYYYMM                  = "yyyyMM"
	YYYY_MM                 = "yyyy-MM"
	YYYY_MM_DD              = "yyyy-MM-dd"
	YYYYMMDD                = "yyyyMMdd"
	YYYYMMDDHHMMSS          = "yyyyMMddHHmmss"
	YYYYMMDDHHMM            = "yyyyMMddHHmm"
	YYYYMMDDHH              = "yyyyMMddHH"
	YYMMDDHHMM              = "yyMMddHHmm"
	MM_DD_HH_MM             = "MM-dd HH:mm"
	MM_DD_HH_MM_SS          = "MM-dd HH:mm:ss"
	YYYY_MM_DD_HH_MM        = "yyyy-MM-dd HH:mm"
	YYYY_MM_DD_HH_MM_SS     = "yyyy-MM-dd HH:mm:ss"
	YYYY_MM_DD_HH_MM_SS_SSS = "yyyy-MM-dd HH:mm:ss.SSS"

	MM_DD_EN                   = "MM/dd"
	YYYY_MM_EN                 = "yyyy/MM"
	YYYY_MM_DD_EN              = "yyyy/MM/dd"
	MM_DD_HH_MM_EN             = "MM/dd HH:mm"
	MM_DD_HH_MM_SS_EN          = "MM/dd HH:mm:ss"
	YYYY_MM_DD_HH_MM_EN        = "yyyy/MM/dd HH:mm"
	YYYY_MM_DD_HH_MM_SS_EN     = "yyyy/MM/dd HH:mm:ss"
	YYYY_MM_DD_HH_MM_SS_SSS_EN = "yyyy/MM/dd HH:mm:ss.SSS"

	MM_DD_CN               = "MM月dd日"
	YYYY_MM_CN             = "yyyy年MM月"
	YYYY_MM_DD_CN          = "yyyy年MM月dd日"
	MM_DD_HH_MM_CN         = "MM月dd日 HH:mm"
	MM_DD_HH_MM_SS_CN      = "MM月dd日 HH:mm:ss"
	YYYY_MM_DD_HH_MM_CN    = "yyyy年MM月dd日 HH:mm"
	YYYY_MM_DD_HH_MM_SS_CN = "yyyy年MM月dd日 HH:mm:ss"

	HH_MM       = "HH:mm"
	HH_MM_SS    = "HH:mm:ss"
	HH_MM_SS_MS = "HH:mm:ss.SSS"
)

func FormatDate(date time.Time, dateStyle DateStyle) string {
	return date.Format(DateLayout(dateStyle))
}

func DateLayout(style DateStyle) string {
	layout := string(style)
	layout = strings.Replace(layout, "yyyy", "2006", 1)
	layout = strings.Replace(layout, "yy", "06", 1)
	layout = strings.Replace(layout, "MM", "01", 1)
	layout = strings.Replace(layout, "dd", "02", 1)
	layout = strings.Replace(layout, "HH", "15", 1)
	layout = strings.Replace(layout, "mm", "04", 1)
	layout = strings.Replace(layout, "ss", "05", 1)
	layout = strings.Replace(layout, "SSS", "000", -1)
	return layout
}

type UsTime time.Time

func (usTime UsTime) MarshalJSON() ([]byte, error) {
	t := fmt.Sprintf(`"%s"`, FormatDate(time.Time(usTime), YYYY_MM_DD_HH_MM_SS))
	return []byte(t), nil
}

func (usTime *UsTime) UnmarshalJSON(b []byte) (err error) {
	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic(err)
	}
	sv := string(b)
	if len(sv) == 10 {
		sv += " 00:00:00"
	} else if len(sv) == 16 {
		sv += ":00"
	}
	now, err := time.ParseInLocation(DateLayout(YYYY_MM_DD_HH_MM_SS), string(b), loc)
	if err != nil {
		if now, err = time.ParseInLocation(DateLayout(YYYY_MM_DD_HH_MM), string(b), loc); err != nil {
			now, err = time.ParseInLocation(DateLayout(YYYY_MM_DD), string(b), loc)
		}
	}
	*usTime = UsTime(now)
	return
}
