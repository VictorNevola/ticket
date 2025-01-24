package time_location

import "time"

var location *time.Location

func init() {
	var err error
	location, err = time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		panic(err)
	}
}

func Now() time.Time {
	return time.Now().In(location)
}

func StringToTime(value string) (time.Time, error) {
	return time.ParseInLocation("2006-01-02 15:04:05", value, location)
}
