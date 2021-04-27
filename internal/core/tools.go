package core

import (
	"strconv"
	"time"
)

func GetTimeStamp() string {
	return strconv.FormatInt(time.Now().Unix(), 10)
}
