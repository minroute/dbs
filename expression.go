// Package dbs
// @Description: 表达式页面
// @todo：string 拼接，不是最优性能，后期考虑使用bytes.Buffer
package dbs

import (
	"errors"
	"strings"
)

type Expression struct {
	finalSql  string
	sql       []string // 生成的sql子句对应的参数
	error     []error
	panic     bool              // 是否抛出异常,不抛出则记录错误
	bindNames map[string]string // 绑定的参数名
}

func (e *Expression) Error() []error {
	return e.error
}

// HandleError 处理表达式遇到错误的操作，抛出异常或者记录错误
func (e *Expression) HandleError(str string) {
	if e.panic {
		panic(e.error)
	} else {
		e.error = append(e.error, errors.New(str))
	}
}

// Build 生成所有条件字句
func (c *Expression) Build() string {

	return c.finalSql
}

// Exp 表达式，原样输出,每一个参数就是一个表达式
func (c *Expression) Exp(v ...string) *Expression {
	if len(v) == 0 {
		return c
	}
	c.sql = append(c.sql, strings.Join(v, Space(SQL_And)))
	if c.finalSql != "" {
		c.finalSql += Space(SQL_And) + strings.Join(v, Space(SQL_And))
	} else {
		c.finalSql = strings.Join(v, Space(SQL_And))
	}

	return c
}

// OrExp 生成一个or的表达式
func (c *Expression) OrExp(v ...string) *Expression {
	if len(v) == 0 {
		return c
	}
	c.sql = append(c.sql, strings.Join(v, " "+SQL_Or+" "))
	if c.finalSql != "" {
		c.finalSql += Space(SQL_Or) + strings.Join(v, Space(SQL_Or))
	} else {
		c.finalSql = strings.Join(v, Space(SQL_Or))
	}
	return c
}

func (c *Expression) In(v ...any) string {

	return ""
}

func (c *Expression) OrIn(v ...any) string {
	return ""
}

func (c *Expression) NotIn(v ...any) string {
	return ""
}

func (c *Expression) OrNotIn(v ...any) string {
	return ""
}

func (c *Expression) Between(v ...any) string {
	return ""
}

func (c *Expression) OrBetween(v ...any) string {
	return ""
}

func (c *Expression) NotBetween(v ...any) string {
	return ""
}

func (c *Expression) OrNotBetween(v ...any) string {
	return ""
}

func (c *Expression) IsNull(v ...any) string {
	return ""
}

func (c *Expression) OrIsNull(v ...any) string {
	return ""
}

func (c *Expression) Is(v ...any) string {
	return ""
}

func (c *Expression) OrIs(v ...any) string {
	return ""
}

func (c *Expression) IsNotNull(v ...any) string {
	return ""
}

func (c *Expression) OrIsNotNull(v ...any) string {
	return ""
}

func (c *Expression) IsExists(v ...any) string {
	return ""
}

func (c *Expression) OrIsExists(v ...any) string {
	return ""
}

func (c *Expression) IsNotExists(v ...any) string {
	return ""
}

func (c *Expression) OrIsNotExists(v ...any) string {
	return ""
}
