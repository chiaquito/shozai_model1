package time

import "time"

func Str2time(t string) (*time.Time, error) {
	// YYYY-MM-DDTHH:MM:SSZZZZの形式で渡される文字列tをtime.Time型に変換して返す
	layout := "2006-01-02T15:04:05-0700"
	parsedTime, err := time.Parse(layout, t)
	if err != nil {
		return nil, err
	}
	return &parsedTime, nil
}
