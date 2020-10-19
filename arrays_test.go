package util

import "testing"

func TestJoinEx(t *testing.T) {
	arr := []interface{}{"a", 1, true}
	t.Log(JoinEx(arr, "[", ",", "]", `""`))
}

func TestJoin(t *testing.T) {

}
