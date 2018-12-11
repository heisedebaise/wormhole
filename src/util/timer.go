package util

import (
	"strings"
)

// ParseTime 解析时间长度。
func ParseTime(str string) int {
	var time int
	if indexOf := strings.Index(str, "d"); indexOf > -1 {
		time = ToInt(str[:indexOf], 0) * 24 * 60 * 60
		str = str[indexOf+1:]
	}
	if indexOf := strings.Index(str, "h"); indexOf > -1 {
		time += ToInt(str[:indexOf], 0) * 60 * 60
		str = str[indexOf+1:]
	}
	if indexOf := strings.Index(str, "m"); indexOf > -1 {
		time += ToInt(str[:indexOf], 0) * 60
		str = str[indexOf+1:]
	}
	if len(str) > 0 {
		time += ToInt(str, 0)
	}

	return time
}
