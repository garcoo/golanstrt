package infrastructure

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type MUser struct {
	User_id         uint64 `gorm:"primary_key"`
	User_name       string
	User_nickname   string
	Chat_account_id uint64
	Activ_flg       string
}

func FindOneByChatAccountId(cw_id int32) []MUser {

	db := GetDBConn()
	db.SingularTable(true)

	var result []MUser

	db.Where("chat_account_id = ? AND activ_flg = ?", cw_id, 1).Find(&result)

	defer db.Close()
	return result
}
