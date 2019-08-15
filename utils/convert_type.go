package utils

import (
	"strconv"
)

func Intf2String(v interface{}) (str string) {
	if v2, ok := v.(int); ok {
		str = strconv.Itoa(v2)
	} else if v3, ok := v.(string); ok {
		str = v3
	} else if v4, ok := v.(int64); ok {
		str = strconv.FormatInt(v4, 10)
	}
	return str
}

func Intf2Int(v interface{}) (n int) {
	switch e := v.(type) {
	case int:
		n = e
	case string:
		n, _ = strconv.Atoi(e)
	case int64:
		n = int(e)
	case float32:
		n = int(e)
	case float64:
		n = int(e)
	}
	return n
}
