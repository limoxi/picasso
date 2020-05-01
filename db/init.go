package db

import (
	"github.com/limoxi/ghost"
	_ "picasso/db/space"
	_ "picasso/db/user"
)

func init(){
	dbConf := ghost.Config.GetMap("database.default")
	db := ghost.ConnectDB(
		ghost.NewDbConfig(
			dbConf.GetString("engine", "mysql"),
			dbConf.GetString("host"),
			dbConf.GetString("port", "3306"),
			dbConf.GetString("user"),
			dbConf.GetString("password"),
			dbConf.GetString("dbname"),
		),
	)
	db.LogMode(true)
	db.SingularTable(true)
}