package convert

import (
	"strings"
	"time"
)

const (
	oneHour int64 = 3600
	dayUnit       = 24 * oneHour
	GMT8          = 8 * oneHour
)

// support SG only
var (
	// time difference of countries vs UTC
	timeDiffMap = map[string]int64{
		"SG": GMT8,
	}

	countryTimezone = map[string]string{
		"SG": "Etc/GMT-8",
	}
)

// TimeLeftAlign help align timestamp to nearby border according to a custom time unit
// for example. time unit is 0.5 hour, 13:23 will be align to 13:00
// the time unit should be less than or equal to day
func TimeLeftAlign(country string, t int64, timeUnit int64) int64 {
	if timeUnit == 0 {
		return t
	}
	var (
		timeDiff int64
		ok       bool
	)
	if timeUnit > oneHour {
		timeDiff, ok = timeDiffMap[strings.ToUpper(country)]
		if !ok {
			timeDiff = GMT8
		}
	}
	if timeUnit > dayUnit {
		timeUnit = dayUnit
	}
	return t - ((t + timeDiff) % timeUnit)
}

// CleanTime("SG", t, time.Minute*30)
func CleanTime(country string, t int64, timeDuration time.Duration) int64 {
	return TimeLeftAlign(country, t, int64(timeDuration.Seconds()))
}

//func main() {
//	t := time.Now().Unix()
//	log.Println(CleanTime("SG", t, time.Minute*30))
//}
