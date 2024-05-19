package mysql

import "errors"

func CheckTokenIfInvalid(token string) (err error) {
	sqlStr := `SELECT count(token) FROM tokenInvalid where token = ?`
	var count int8
	err = db.Get(&count, sqlStr, token)
	if count > 0 {
		return errors.New(tokenIsInvalid)
	}
	return nil
}
