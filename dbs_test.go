package dbs

import (
	"encoding/json"
	"fmt"
	"testing"
)

func Where(a ...any) {
	t, e := json.Marshal(a)
	fmt.Printf("%+v ,e:%v \n", string(t), e)
}

func TestWhere(t *testing.T) {

	// @todo:多次调用where，默认是等于多个And条件。

	// 单个列的设置形式，操作符跟列名组合在一起
	Where("sql")     // where sql
	Where("a", "b")  // where a = b
	Where("a >", 1)  // where a > 1
	Where("a !=", 1) // where a != 1

	// 操作符的形式：  Where("字段", "操作符", 值) // where 字段 操作符 值
	Where("a", "<", 1)
	Where("a", "<=", 1)
	Where("a", ">", 1)
	Where("a", ">=", 1)
	Where("a", "=", 1)
	Where("a", "!=", 1)
	Where("a", "like", "%1%")
	Where("a", "not like", "%1%")
	Where("a", "in", []string{"1", "2"})
	Where("a", "not in", []string{"1", "2"})
	Where("a", "between", []string{"1", "2"})
	Where("a", "not between", []string{"1", "2"})
	Where("a", "is", nil)
	Where("a", "is not", nil)
	Where("a", "is null")
	Where("a", SQL_IsNull)
	Where("a", SQL_IsNull)

	// // 内置便捷方法
	// Where(
	// 	Like("a", "b"),
	// 	Like("c", "d%"),
	// 	OrLike("e", "%f"),
	// ) // where a like '%b%' and c like 'd%' or e like '%f'
	//
	// // 参数方法和非参数方法混合使用
	// Where(
	// 	"a", "<", 1,
	// 	Like("a", "b"),
	// ) // where a < 1 and a like '%b%'

	// 多个列的设置形式
	// Where(&map[string]interface{}{"a": "a'value", "b": []int{1, 2}})             // where a=a.value and b in (b1,b2)
	// Where(&data_sctructure.StrAnyMap{"a": "a'value", "b": []string{"b1", "b2"}}) // where a=a.value and b in (b1,b2)

	// as := kit.AnySet().Add("a").Add(SQL_Like).Add("%1%")
	// Where(as)
	// Where(as)
	// Where(as, as, as, as, as)

}
