// Package dbs
// @Description: 表达式页面
// @todo：string 拼接，不是最优性能，后期考虑使用bytes.Buffer
package dbs

import (
	"strings"

	"github.com/xiiapp/be"
	"github.com/xiiapp/dbs/helper"
)

// Like系列抽离出来的公共方法，用于生成Like系列的的表达式
// @todo：1.对列名进行格式化
func (c *Expression) likeExecutor(opType, likeType string, v ...any) *Expression {

	count := len(v)
	if count == 0 {
		return c
	}

	// Checker
	// 检测第一个参数是map类型时,参数总数不得超过1个。
	_, isOriginStrAnyMap := v[0].(map[string]any)
	if helper.IsStrAnyMap(v[0]) || helper.IsStrAnyMapFromPtr(v[0]) || isOriginStrAnyMap {
		if count > 1 {
			c.HandleError(ErrLikeInvalidStranymap)
		}

		// 本节主要匹配方法调用：
		// c.Like(&helper.StrAnyMap{"name": "a", "title": "b%", "c": &[]string{"1", "2"}})  //StrAnyMap指针，
		// c.Like(helper.StrAnyMap{"name": "a", "title": "b%", "c": &[]string{"1", "2"}})   //StrAnyMap值，
		// c.Like(map[string]interface{}{"name": "a", "title": "b%", "c": &[]string{"1", "2"}}) //map[string]interface{}值，
		// 注意：map的key的对应的值可以是 string,int,&[]string,[]string,[]int,&[]int
		// @todo,对这个解析的实现。

		return c
	} else {
		// 检测第一个参数是string,否则报错
		if value, isString := v[0].(string); isString {
			// 当只有一个参数时，即默认原生sql表达式，直接拼接即可。
			if count == 1 {

				c.Exp(value)
				return c

			} else if count == 2 {
				// 有两个参数的的，第二个参数可能是string或者[]string或者[]int两种情况
				// 第一种情况：[]string或者[]int，全部专程[]string后拼接
				// 本节主要匹配一下方法调用，其他*Like方法类似：
				//  c.Like("name", []string{"Joe1", "%Han", "Eric%"})  // name like '%Joe%' or name like '%Han' or name like 'Eric%'
				//	c.Like("name", &[]string{"Joe2", "%Han", "Eric%"}) // name like '%Joe%' or name like '%Han' or name like 'Eric%'
				//	c.Like("name", &[]int{1, 2, 3})                    // name like '%1%' or name like '%2%' or name like '%3%'
				//	c.Like("name", []int{1, 2, 3})                     // name like '%1%' or name like '%2%' or name like '%3%'
				isSlice := helper.IsSlice(v[1])
				isSlicePtr := helper.IsSlicePtr(v[1])
				if isSlice || isSlicePtr { // 是slice，
					var items []string
					if isSlice {
						items = []string((v[1]).([]string))
					} else {
						items = *(*[]string)(v[1].(*[]string))
					}

					if len(items) == 0 {
						c.HandleError(ErrLikeEmptySlice)
						return c
					}

					var tempStr []string
					for _, item := range items {
						temp := be.String(item)
						if !strings.Contains(temp, "%") {
							temp = "%" + temp + "%"
						}

						bindName := ":" + helper.RandStr(10)
						c.bindNames[bindName] = temp
						s := string(v[0].(string)) + Space(likeType) + bindName
						tempStr = append(tempStr, s)
					}
					c.Exp(strings.Join(tempStr, Space(opType)))

					return c
				} else {
					// 第二个参数是string，直接拼接，按设计约定，应该就是模糊搜索的参数值
					// 本节主要匹配方法调用：
					//      c.Like("name", "张三")  // 生成：name like '%张三%'
					//      c.Like("name", "张三%") // 生成：name like '%张三'
					//      c.Like("name", "%张三") // 生成：name like '张三%'
					// 其他*Like系列方法如是。
					bindName := ":" + helper.RandStr(10)
					temp := string(v[1].(string))
					if !strings.Contains(temp, "%") {
						temp = "%" + temp + "%"
					}
					c.bindNames[bindName] = temp
					s := string(v[0].(string)) + Space(likeType) + bindName
					c.Exp(s)
					return c
				}

			} else {
				// 参数大于2个的，遍历所有参数是否都是字符串，否则有错
				// 本节主要为了匹配方法调用：c.Like("name", "Joe", "Han", "Eric") 这种情况
				// 本节最终生成：Name like :name and Name like :name2 and Name like :name3
				var tempStr []string
				for index, item := range v {
					if _, isString := item.(string); !isString {
						c.HandleError(ErrLikeInvalidParamWhenMoreThanTwo)
						return c
					}

					if index > 0 {
						temp := be.String(item)
						if !strings.Contains(temp, "%") {
							temp = "%" + temp + "%"
						}

						bindName := ":" + helper.RandStr(10)
						c.bindNames[bindName] = temp
						s := string(v[0].(string)) + Space(likeType) + bindName
						tempStr = append(tempStr, s)
					}
				}
				c.Exp(strings.Join(tempStr, Space(opType)))
				return c
			}

		} else {
			// 第一个参数不是string，也不是map类型，直接报错
			c.HandleError(ErrLikeFirstParamMustBeStranymapOrString)
			return c
		}
	}

	return c
}

// Like ,
// @todo：通过给参数指定泛型参数来规避掉内部的类型判断？
func (c *Expression) Like(v ...any) *Expression {
	return c.likeExecutor(SQL_And, SQL_Like, v...)
}

// OrLike
// @todo：应该报错，OrLike 在最终生成的sql中，如果只有一个字段like一个值，这是意义不大的，应该建议使用Like
func (c *Expression) OrLike(v ...any) *Expression {
	return c.likeExecutor(SQL_Or, SQL_Like, v...)
}

func (c *Expression) NotLike(v ...any) *Expression {
	return c.likeExecutor(SQL_And, SQL_NOT_LIKE, v...)
}

// OrNotLike
// @todo：应该报错，OrLike 在最终生成的sql中，如果只有一个字段like一个值，这是意义不大的，应该建议使用Like
func (c *Expression) OrNotLike(v ...any) *Expression {
	return c.likeExecutor(SQL_Or, SQL_NOT_LIKE, v...)
}
