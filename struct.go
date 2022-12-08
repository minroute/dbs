package dbs

type (
	// TableNameFunc 通过转换得到一个表名，可以是struct里的一个列名，或者其他
	TableNameFunc func(a any) string

	// FieldMapFunc converts a struct field name into a DB column name.
	FieldMapFunc func(string) string
)

// For hook
type (
	// BeforeInsertFunc will be  called before insert execute
	BeforeInsertFunc interface {
		beforeAdd()
	}

	// AfterInsertFunc will be  called after insert execute
	AfterInsertFunc interface {
		afterAdd()
	}

	// BeforeDeleteFunc will be  called before delete execute
	BeforeDeleteFunc interface {
		beforeDel()
	}

	// AfterDeleteFunc will be  called after delete execute
	AfterDeleteFunc interface {
		afterDel()
	}

	// BeforeUpdateFunc will be  called before update execute
	BeforeUpdateFunc interface {
		beforeUpdate()
	}

	// AfterUpdateFunc will be  called after update execute
	AfterUpdateFunc interface {
		afterUpdate()
	}

	// BeforeReplaceFunc  will be  called before replace execute
	BeforeReplaceFunc interface {
		beforeReplace()
	}

	// AfterReplaceFunc  will be  called after replace execute
	AfterReplaceFunc interface {
		afterReplace()
	}

	// BeforeSelectFunc  will be  called before query
	BeforeSelectFunc interface {
		beforeSelect()
	}

	// AfterSelectFunc  will be  called after query
	AfterSelectFunc interface {
		afterSelect()
	}
)

type HookFunc struct {
	// sql 操作的各种hook
	BeforeInsert  BeforeInsertFunc  `-`
	AfterInsert   AfterInsertFunc   `-`
	BeforeDelete  BeforeDeleteFunc  `-`
	AfterDelete   AfterDeleteFunc   `-`
	BeforeUpdate  BeforeUpdateFunc  `-`
	AfterUpdate   AfterUpdateFunc   `-`
	BeforeReplace BeforeReplaceFunc `-`
	AfterReplace  AfterReplaceFunc  `-`
}
