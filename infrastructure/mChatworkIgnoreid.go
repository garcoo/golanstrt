package infrastructure

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type MChatworkIgnoreid struct {
	chat_account_id uint64
}

func XXXXXXXXXXXXXXXXXXXXXXXXXXxx(cw_id int32) []MUser {

	db := GetDBConn()
	db.SingularTable(true)

	var result []MUser

	db.First("chat_account_id = ? AND activ_flg = ?", cw_id, 1).Find(&result)

	return result
}
