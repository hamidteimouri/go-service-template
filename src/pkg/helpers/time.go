package helpers

import (
	"github.com/hamidteimouri/gommon/htenvier"
	"time"
)

var (
	TimeFullFormat = "2006-01-02 15:04:05"
)

func GetTimezone() string {
	return htenvier.EnvOrDefault("TIMEZONE", "Asia/Tehran")
}

func Now() string {
	/* WAY  01 */
	/*
		loc, _ := time.LoadLocation(TIMEZONE)
		time.Local = loc
		return time.Now().Format("2006-01-02 15:04:05")
	*/

	/* WAY 02 */
	loc, _ := time.LoadLocation(GetTimezone())
	return time.Now().In(loc).Format("2006-01-02 15:04:05")
}

func NowInTime() string {
	loc, _ := time.LoadLocation(GetTimezone())
	return time.Now().In(loc).Format("15:04")
}
