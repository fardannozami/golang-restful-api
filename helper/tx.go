package helper

import "database/sql"

func CommitOrRollback(tx *sql.Tx) {
	if tx == nil {
		return
	}

	err := recover()
	if err != nil {
		_ = tx.Rollback()
		panic(err)
	} else {
		_ = tx.Commit()
	}
}
