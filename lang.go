// Package dbs
// @Description: 临时使用，用来存储一些翻译用的常量
//
package dbs

const (
	ErrLikeInvalidStranymap = `
Like方法的第一个参数是map[string]any类型时，不能有第二个参数。
调用例子1：Like(map[string]any{"name": "张三", "age": 18})
调用例子2：Like(&kit.StrAnyMap{"name": "张三", "age": 18})
调用例子3：Like(kit.StrAnyMap{"name": "张三", "age": 18})
 `
	ErrLikeFirstParamMustBeStranymapOrString = `Like方法的第一个参数必须是map[string]any类型或string类型。`
	ErrLikeEmptySlice                        = `Like方法的第二个参数是切片类型时，slice不能为空。`
	ErrLikeInvalidParamWhenMoreThanTwo       = `Like方法的参数个数大于2个时，所有参数类型必须都是string类型。`
)
