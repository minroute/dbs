package dbs

import (
	"fmt"
	"testing"

	"github.com/minroute/dbs/helper"
)

func TestExpression_Like(t *testing.T) {
	var c Expression

	// 单字段，多模糊搜索
	// c.Like("name", "Joe", "Han", "Eric")              // name like '%Joe%' or name like '%Han' or name like 'Eric%'
	c.Like("name", []string{"Joe1", "%Han", "Eric%"})  // name like '%Joe%' or name like '%Han' or name like 'Eric%'
	c.Like("name", &[]string{"Joe2", "%Han", "Eric%"}) // name like '%Joe%' or name like '%Han' or name like 'Eric%'
	// c.Like("name", &[]int{1, 2, 3})                    // name like '%1%' or name like '%2%' or name like '%3%'
	// c.Like("name", []int{1, 2, 3})                     // name like '%1%' or name like '%2%' or name like '%3%'

	println(c.Build())
}

// 多字段模糊查询，应该传入一个map结构，也只能有一个参数
func TestExpression_Like_Multi_Columns(t *testing.T) {
	var c Expression

	// 正确传参
	if errs := c.Like(&helper.StrAnyMap{"name": "a", "title": "b%", "c": &[]string{"1", "2"}}).Error(); errs != nil {
		t.Errorf(fmt.Sprintf("%v", errs))
	}
	if errs := c.Like(helper.StrAnyMap{"name": "a", "title": "b%", "c": &[]string{"1", "2"}}).Error(); errs != nil {
		t.Errorf(fmt.Sprintf("%v", errs))
	}
	if errs := c.Like(map[string]interface{}{"name": "a", "title": "b%", "c": &[]string{"1", "2"}}).Error(); errs != nil {
		t.Errorf(fmt.Sprintf("%v", errs))
	}

	// 错误传参
	if errs := c.Like(&helper.StrAnyMap{"name": "a", "title": "b%", "c": &[]string{"1", "2"}}, 2).Error(); errs == nil {
		t.Errorf(fmt.Sprintf("%v", errs))
	}
	if errs := c.Like(helper.StrAnyMap{"name": "a", "title": "b%", "c": &[]string{"1", "2"}}, 2).Error(); errs == nil {
		t.Errorf(fmt.Sprintf("%v", errs))
	}
	if errs := c.Like(map[string]interface{}{"name": "a", "title": "b%", "c": &[]string{"1", "2"}}, 2).Error(); errs == nil {
		t.Errorf(fmt.Sprintf("%v", errs))
	}

}

//
// func TestCondition(t *testing.T) {
// 	var c Expression
//
// 	// 原生表达式
// 	c.Exp("contact('a','b') as c", "id=1")
// 	c.OrExp("id=2", "id=3")
//
// 	// Like用法-单个字段
// 	c.Like("name", "a")                               // name like '%a%'
// 	c.Like("name", "b%")                              // name like 'a%'
// 	c.Like("name", "%c")                              // name like '%c'
// 	c.Like("name", &[]string{"Joe", "%Han", "Eric%"}) // name like '%Joe%' or name like '%Han' or name like 'Eric%'
//
// 	// Like用法-多个字段
//
// 	// name like '%a%' and title like 'b%' and ( c like '%1%' or c like '%2%')
// 	c.Like(&data_sctructure.StrAnyMap{"name": "a", "title": "b%", "c": &[]string{"1", "2"}})
//
// 	println(c.Build())
// }
//
// // cmd:go test -v expression_like_test.go const.go expression_like.go  struct.go -test.run TestCondition_Like
// func TestCondition_Like(t *testing.T) {
//
// 	var c Expression
//
// 	// Like用法-单个字段
// 	c.Like("name", "a")                               // name like '%a%'
// 	c.Like("name", "b%")                              // name like 'a%'
// 	c.Like("name", "%c")                              // name like '%c'
// 	c.Like("name", "Joe", "Han", "Eric")              // name like '%Joe%' or name like '%Han' or name like 'Eric%'
// 	c.Like("name", &[]string{"Joe", "%Han", "Eric%"}) // name like '%Joe%' or name like '%Han' or name like 'Eric%'
//
// 	// Like用法-多个字段
// 	c.Like(&data_sctructure.StrAnyMap{"name": "a", "title": "b%", "c": &[]string{"1", "2"}}) // name like '%a%' and title like 'b%' and ( c like '%1%' or c like '%2%')
//
// 	println(c.Build())
// }
