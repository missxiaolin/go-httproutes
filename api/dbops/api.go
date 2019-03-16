package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func openConn() * sql.DB  {
	dbConn, err := sql.Open("mysql", "root:!@#@tcp(localhost:3306)/go_video?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
	return  dbConn
}

func AddUserCredential(loginName string, pwd string) error  {
	
}

func GetUserCredential(loginName string) (string, error) {

}
