// Package dbs
// @Description: 扩展事务
// @Feature:
// 1.通过计数器解决事务嵌套的问题
// 2.通过savepoint解决事务嵌套的问题 @todo:需要进一步测试
package dbs

import "strconv"

func (d *DB) Begin() (*DB, error) {
	d.isRollback = true
	if d.txModeIsSavePoint {
		d.txNext = d.cloneForTx()
		d.txNext.txCounter += 1
		d.tx.Exec("SAVEPOINT SP " + strconv.Itoa(d.txNext.txCounter))
		return d.txNext, nil
	} else {
		if d.tx == nil {
			tx, err := d.db.Begin()
			if err != nil {
				return nil, err
			}
			d.tx = tx
			d.txCounter = txCommitNum
			return d, nil
		}
		d.txCounter++
	}

	return d, nil
}

func (d *DB) Commit() error {

	if d.txModeIsSavePoint {
		if d.txNext != nil && !d.txNext.txResolved {
			err := d.txNext.tx.Commit()
			if err != nil {
				return err
			}
		}

		d.txResolved = true

		if d.txCounter > 0 {
			_, err := d.tx.Exec("RELEASE SAVEPOINT SP" + strconv.Itoa(d.txCounter))
			return err
		}
		return d.tx.Commit()

	} else {
		d.isRollback = false
		if d.tx != nil {
			if d.txCounter == txCommitNum {
				err := d.tx.Commit()
				if err != nil {
					return err
				}
				d.resetTx()
				return nil
			}
			d.txCounter--
		}
	}

	return nil
}

func (d *DB) Rollback() error {
	if d.txModeIsSavePoint {
		d.txResolved = true
		if d.txCounter > 0 {
			_, err := d.tx.Exec("ROLLBACK TO SAVEPOINT SP" + strconv.Itoa(d.txCounter))
			return err
		}
		return d.tx.Rollback()
	} else {
		if d.tx != nil && d.isRollback {
			err := d.tx.Rollback()
			if err != nil {
				return err
			}
			d.resetTx()
			return nil
		}
	}
	return nil
}

func (d *DB) resetTx() {
	d.tx = nil
	d.txCounter = 0
	d.isRollback = false
}

func (d *DB) cloneForTx() *DB {
	return &DB{
		TableName:         d.TableName,
		FieldMapper:       d.FieldMapper,
		LogFunc:           d.LogFunc,
		driver:            d.driver,
		ctx:               d.ctx,
		db:                d.db,
		tx:                d.tx,
		txCounter:         d.txCounter,
		txModeIsSavePoint: d.txModeIsSavePoint,
		isRollback:        d.isRollback,
		txNext:            nil,
		txResolved:        false,
	}
}
