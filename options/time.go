package options

import "time"

func GetStartTime() (startTime string) {
	startTime = time.Now().Format("2006-01-02 15:04:05")
	return startTime
}

func GetEndTime() (endTime string) {
	endTime = time.Now().Format("2006-01-02 15:04:05")
	return endTime
}
