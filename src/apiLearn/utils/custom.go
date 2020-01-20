package utils

import (
	"fmt"
	"time"
)

func GetDateTime() string{
	today := time.Now()
	day := today.Day()
	month :=int(today.Month())
	year :=today.Year()
	hour := today.Hour()
	minutes := today.Minute()
	seconds := today.Second()

	return fmt.Sprintf("%d-%d-%d %d:%d:%d",year,month,day,hour,minutes,seconds)
}
