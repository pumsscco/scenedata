package data
import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)
var (
	Db *sql.DB
	Layout string="2006-01-02 15:04:05"
)
func init() {
	Db, _ = sql.Open("mysql", "golang:ktnpu9hxm27sr26g@/repo")
	truncGob:="truncate GameObject"
	Db.Exec(truncGob)
	truncMst:="truncate MstInfo"
	Db.Exec(truncMst)
	truncNPC:="truncate NPCInfo"
	Db.Exec(truncNPC)
	truncPwy:="truncate PathWay"
	Db.Exec(truncPwy)
}
