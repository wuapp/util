package util

import (
	"fmt"
	"strconv"
)

//ToString convert an arbitrary value to string.
//If the type of the value implements the `fmt.Stringer` interface, the `String()` method that is,
//this function will return as the `String()` returns.
func ToString(a interface{}, quotes string) string {
	switch a := a.(type) {
	case string:
		if quotes != "" {
			return quotes + a + quotes
		}
		return a
	case bool:
		if a {
			return "true"
		} else {
			return "false"
		}
	case int:
		return strconv.Itoa(a)
	case int8:
		return strconv.Itoa(int(a))
	case int16:
		return strconv.Itoa(int(a))
	case int32:
		return strconv.Itoa(int(a))
	case int64:
		return strconv.FormatInt(a, 10)
	case float32:
		return strconv.FormatFloat(float64(a), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(float64(a), 'f', -1, 64)
	case []byte:
		if quotes != "" {
			return quotes + string(a) + quotes
		}
		return string(a)
	case complex64:
		//todo
	case complex128:
		//todo
	case fmt.Stringer:
		return a.String()
	case struct{}:
		return "[sruct]"
	case interface{}:
		return "[interface]"

	}
	fmt.Println()
	return ""
}
