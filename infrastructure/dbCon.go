package infrastructure

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func GetDBConn() *gorm.DB {
	db, err := gorm.Open(GetDBConfig())
	if err != nil {
		panic(err)
	}

	db.LogMode(true)
	return db
}

func GetDBConfig() (string, string) {
	DBMS := "XXXXXX"
	USER := "XXXXXXX"
	PASS := "XXXXXX"
	PROTOCOL := "tcp(XXX.XXX.XXX.XXX:XXX)"
	DBNAME := "XXXXXXXXXXX"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME // + "?" + OPTION

	return DBMS, CONNECT
}
