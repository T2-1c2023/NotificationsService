package utilities

import "time"

func GetCurrentDate() string {
	currentTime := time.Now()
	formattedTime := currentTime.Format("2023-05-24 20:25:10")
	return formattedTime
}
