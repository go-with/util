package util

import (
	"database/sql"
)

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func IsErrNoRows(err error) bool {
	return sql.ErrNoRows == err
}
