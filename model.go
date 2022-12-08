// Package dbs
// @Description: Model Support
//
package dbs

type Model interface {
	TableName() string
	Index() string
	Unique() string
}

// FinalTableName  will return the final table name  when using the ORM mode.
// Ignore if  direct execute raw sql or  sql builder
// @todo:Model操作化
func (d *DB) FinalTableName() string {
	if d.TableName != nil {
		return d.TableName(nil)
	}
	return "TODO"
}
