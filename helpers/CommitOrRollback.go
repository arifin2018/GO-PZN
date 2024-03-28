package helpers

import "database/sql"

func CommitOrRollback(sqlTx *sql.Tx, err error) {
	if err != nil {
		errorRollback := sqlTx.Rollback()
		PanicIfError(errorRollback)
	} else {
		errorCommit := sqlTx.Commit()
		PanicIfError(errorCommit)
	}
}
