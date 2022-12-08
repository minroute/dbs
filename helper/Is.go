// Package dbs
// @Description: 对buildin的类型进行判断。包括string,int,bool,struct,bytes,ptr,array,slice,map,chan,func
//
// @note：
//	// 自定义一个类型，可以通过断言来判断是否是这个类型
//	// type (
//	// 	StrAnyMap map[string]any
//	// )
//	//
//	// // IsStrAnyMap 是否是指定的Struct
//	// func IsStrAnyMap(v any) bool {
//	// 	switch v.(type) {
//	// 	case StrAnyMap:
//	// 		return true
//	// 	}
//	// 	return false
//	// }

package helper

import (
	"reflect"
)

// IsString  判断类型是否是字符串
func IsString(v any) bool {
	switch v.(type) {
	case string:
		return true
	}
	return false
}

// IsInt 判断类型是否是数字，包括int,int8,int16,int32,int64,uint,uint8,uint16,uint32,uint64
func IsInt(v any) bool {
	switch v.(type) {
	case int, int8, int16, int32, int64:
		return true
	case uint, uint8, uint16, uint32, uint64:
		return true
	}
	return false
}

// IsBool 判断类型是否是布尔值
func IsBool(v any) bool {
	switch v.(type) {
	case bool:
		return true
	}
	return false
}

// IsStruct 判断类型是否是结构体,
// Struct is a value type that represents a structured data record consisting of a sequence of named elements, called fields, each of which has a name and a type.
// 此处无法通过断言来判断，只能是reflect来辨别
func IsStruct(v any) bool {
	if reflect.TypeOf(v).Kind() == reflect.Struct {
		return true
	}
	return false
}

// IsBytes 判断类型是否是字节切片
func IsBytes(v any) bool {
	switch v.(type) {
	case []byte:
		return true
	}
	return false
}

// IsPtr 判断类型是否是指针
func IsPtr(v any) bool {
	if reflect.TypeOf(v).Kind() == reflect.Ptr {
		return true
	}
	return false
}

// IsArray 判断类型是否是数组
func IsArray(v any) bool {
	if reflect.TypeOf(v).Kind() == reflect.Array {
		return true
	}
	return false
}

// IsSlice 判断类型是否是切片
func IsSlice(v any) bool {
	if reflect.TypeOf(v).Kind() == reflect.Slice {
		return true
	}
	return false
}

// IsSlicePtr 判断类型是否是切片指针
func IsSlicePtr(v any) bool {
	if reflect.ValueOf(v).Elem().Kind() == reflect.Slice && reflect.TypeOf(v).Kind() == reflect.Ptr {
		return true
	}
	return false
}

// IsMap 判断类型是否是Map
func IsMap(v any) bool {
	if reflect.TypeOf(v).Kind() == reflect.Map {
		return true
	}
	return false
}

// IsChan 判断类型是否是通道
func IsChan(v any) bool {
	if reflect.TypeOf(v).Kind() == reflect.Chan {
		return true
	}
	return false
}

// IsFunc 判断类型是否是函数
func IsFunc(v any) bool {
	if reflect.TypeOf(v).Kind() == reflect.Func {
		return true
	}
	return false
}

// IsInterface 判断类型是否是接口
func IsInterface(v any) bool {
	if reflect.TypeOf(v).Kind() == reflect.Interface {
		return true
	}
	return false
}
