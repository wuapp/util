package util

import (
	"bytes"
	"strconv"
)

func Join(array interface{}, sep string) string {
	return JoinEx(array, "", sep, "", "")
}

func JoinEx(array interface{}, start, sep, end, quotes string) string {
	buf := new(bytes.Buffer)
	if start != "" {
		buf.WriteString(start)
	}

	arrInt, ok := array.([]int)
	if ok {
		for i, l := 0, len(arrInt); i < l; i++ {
			buf.WriteString(strconv.Itoa(arrInt[i]))
			if i != l-1 {
				buf.WriteString(sep)
			}
		}
	}

	arrStr, ok := array.([]string)
	if ok {
		for i, l := 0, len(arrStr); i < l; i++ {
			buf.WriteString(arrStr[i])
			if i != l-1 {
				buf.WriteString(sep)
			}
		}
	}

	arrInterface, ok := array.([]interface{})
	if ok {
		for i, l := 0, len(arrInterface); i < l; i++ {
			buf.WriteString(ToString(arrInterface[i], quotes))
			if i != l-1 {
				buf.WriteString(sep)
			}
		}
		for _, a := range arrInterface {
			buf.WriteString(ToString(a, quotes))
		}
	}

	//todo: other types

	if end != "" {
		buf.WriteString(end)
	}

	return buf.String()
}
