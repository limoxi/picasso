package db

// do nothing
import ( // make sure this is the forth line, or code gen will fail
	"github.com/limoxi/ghost"
	_ "picasso/db/file"
	_ "picasso/db/user"
)

func init() {
	dbConf := ghost.Config.GetMap("database.default")
	ghost.ConnectDB(
		ghost.NewDbConfig(
			dbConf.GetString("engine", "mysql"),
			dbConf.GetString("host"),
			dbConf.GetString("port", "3306"),
			dbConf.GetString("user"),
			dbConf.GetString("password"),
			dbConf.GetString("dbname"),
		),
	)
}
