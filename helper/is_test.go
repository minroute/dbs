// @todo：完善单元测试
package helper

import (
	"fmt"
	"testing"
)

func TestPtrSlice(t *testing.T) {
	s := []string{"Joe1", "%Han", "Eric%"}
	s2 := IsSlicePtr(s)
	// s2 := reflect.TypeOf(s).Kind()
	fmt.Printf("%+v", s2)
}

func TestIsString(t *testing.T) {
	if !IsString("a") {
		t.Error("IsString error")
	}

	s := "a"
	if !IsString(s) {
		t.Error("IsString error-a")
	}
}

func TestIsInt(t *testing.T) {
	if !IsInt(1) {
		t.Error("IsInt error")
	}

	i := -1
	if !IsInt(i) {
		t.Error("IsInt error-a")
	}

	i2 := 1.6
	if IsInt(i2) {
		t.Error("IsInt error-b")
	}

	i3 := int64(1)
	if !IsInt(i3) {
		t.Error("IsInt error-c")
	}

	i4 := int32(1)
	if !IsInt(i4) {
		t.Error("IsInt error-d")
	}

	i5 := int16(1)
	if !IsInt(i5) {
		t.Error("IsInt error-e")
	}

	i6 := int8(1)
	if !IsInt(i6) {
		t.Error("IsInt error-f")
	}

	i7 := uint(1)
	if !IsInt(i7) {
		t.Error("IsInt error-g")
	}

	i8 := uint64(1)
	if !IsInt(i8) {
		t.Error("IsInt error-h")
	}

	i9 := uint32(1)
	if !IsInt(i9) {
		t.Error("IsInt error-i")
	}

	i10 := "AAA"
	if IsInt(i10) {
		t.Error("IsInt error-j")
	}

	i11 := &[3]int{1, 2, 3}
	if IsInt(i11) {
		t.Error("IsInt error-k")
	}

	i12 := []int{1, 2, 3}
	if IsInt(i12) {
		t.Error("IsInt error-l")
	}
}

func TestIsBool(t *testing.T) {
	if !IsBool(true) {
		t.Error("IsBool error")
	}

	b := true
	if !IsBool(b) {
		t.Error("IsBool error-a")
	}

	b2 := 1
	if IsBool(b2) {
		t.Error("IsBool error-b")
	}

	b3 := 0
	if IsBool(b3) {
		t.Error("IsBool error-c")
	}

	b4 := "10"
	if IsBool(b4) {
		t.Error("IsBool error-d")
	}
}

func TestIsPtr(t *testing.T) {

	// 传地址，
	if !IsPtr(&[]string{"a", "b"}) {
		t.Error("IsPtr error")
	}

	p := *(&[]string{"a", "b"})
	if IsPtr(p) {
		t.Error("IsPtr error-a")
	}

	// 直接传类型值，应该返回false
	if IsPtr([]string{"a", "b"}) {
		t.Error("IsPtr error-b")
	}

}

//
// func TestIsFloat(t *testing.T) {
// 	if !IsFloat(1.1) {
// 		t.Error("IsFloat error")
// 	}
//
// 	f := 1.1
// 	if !IsFloat(f) {
// 		t.Error("IsFloat error-a")
// 	}
//
// 	f2 := 1
// 	if IsFloat(f2) {
// 		t.Error("IsFloat error-b")
// 	}
//
// 	f3 := "1"
// 	if IsFloat(f3) {
// 		t.Error("IsFloat error-c")
// 	}
// }

func TestIsSlice(t *testing.T) {
	if !IsSlice([]string{"a", "b"}) {
		t.Error("IsSlice error")
	}

	s := []string{"a", "b"}
	if !IsSlice(s) {
		t.Error("IsSlice error-a")
	}

	s2 := []int{1, 2}
	if !IsSlice(s2) {
		t.Error("IsSlice error-b")
	}

	s3 := []interface{}{"a", 1}
	if !IsSlice(s3) {
		t.Error("IsSlice error-c")
	}

	s4 := "a"
	if IsSlice(s4) {
		t.Error("IsSlice error-d")
	}

	s5 := []string{}
	if !IsSlice(s5) {
		t.Error("IsSlice error-e")
	}

	s6 := []string(nil)
	if !IsSlice(s6) {
		t.Error("IsSlice error-f")
	}
}

func TestIsMap(t *testing.T) {
	if !IsMap(map[string]string{"a": "b"}) {
		t.Error("IsMap error")
	}

	m := map[string]string{"a": "b"}
	if !IsMap(m) {
		t.Error("IsMap error-a")
	}

	m2 := map[string]int{"a": 1}
	if !IsMap(m2) {
		t.Error("IsMap error-b")
	}

	m3 := map[string]interface{}{"a": "b"}
	if !IsMap(m3) {
		t.Error("IsMap error-c")
	}

	m4 := map[string]interface{}{"a": 1}
	if !IsMap(m4) {
		t.Error("IsMap error-d")
	}

	m5 := map[string]interface{}{"a": []string{"a", "b"}}
	if !IsMap(m5) {
		t.Error("IsMap error-e")
	}

	m6 := map[string]interface{}{"a": map[string]string{"a": "b"}}
	if !IsMap(m6) {
		t.Error("IsMap error-f")
	}

	m7 := map[string]interface{}{"a": map[string]interface{}{"a": "b"}}
	if !IsMap(m7) {
		t.Error("IsMap error-g")
	}

	m8 := map[string]interface{}{"a": map[string]interface{}{"a": 1}}
	if !IsMap(m8) {
		t.Error("IsMap error-h")
	}

	m9 := map[string]interface{}{"a": map[string]interface{}{"a": []string{"a", "b"}}}
	if !IsMap(m9) {
		t.Error("IsMap error-i")
	}

	m10 := map[string]interface{}{"a": map[string]interface{}{"a": map[string]string{"a": "b"}}}
	if !IsMap(m10) {
		t.Error("IsMap error-j")
	}

	m11 := map[string]interface{}{"a": map[string]interface{}{"a": map[string]interface{}{"a": "b"}}}
	if !IsMap(m11) {
		t.Error("IsMap error-k")
	}

	m12 := map[string]interface{}{"a": map[string]interface{}{"a": map[string]interface{}{"a": 1}}}
	if !IsMap(m12) {
		t.Error("IsMap error-l")
	}
}

func TestIsStruct(t *testing.T) {
	if !IsStruct(struct {
		A string
		B int
	}{"a", 1}) {
		t.Error("IsStruct error")
	}

}

// @todo: 测试不全
