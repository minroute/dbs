// Package kit
// @Description: 非类型安全的map[string]interface{}，并提供一个判断是否是StrAnyMap的方法
//
package helper

type (
	StrAnyMap   map[string]any
	StrAnyMapIf interface {
		IsMe()
	}
)

func (s *StrAnyMap) IsMe() {}

// IsStrAnyMapFromPtr 判断是否是StrAnyMap的指针,
// @usage: IsStrAnyMapFromPtr(&StrAnyMap{"a": 1})
func IsStrAnyMapFromPtr(d any) bool {
	if _, ok := any(d).(StrAnyMapIf); ok {
		return true
	}
	return false
}

// IsStrAnyMap 判断是否是StrAnyMap,
// @usage: IsStrAnyMap(StrAnyMap{"a": 1})
func IsStrAnyMap(d any) bool {
	if _, ok := any(d).(*StrAnyMap); ok {
		return true
	}
	return false
}
