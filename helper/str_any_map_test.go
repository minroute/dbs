package helper

import "testing"

func TestStrAnyMap(t *testing.T) {
	s := &StrAnyMap{"a": "b", "c": &[]string{"1", "2"}}
	if !IsStrAnyMap(s) {
		t.Error("IsStrAnyMap error")
	}
}
