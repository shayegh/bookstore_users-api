package usersdb

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)
const (
	mysql_user = "mysql_user"
	mysql_pass = "mysql_pass"
	mysql_host = "mysql_host"
	mysql_schema = "mysql_schema"
)
var(
	Client *sql.DB
	username = os.Getenv(mysql_user)
	password = os.Getenv(mysql_pass)
	host = os.Getenv(mysql_host)
	scheme = os.Getenv(mysql_schema)
)

func init(){
	log.Println(username)
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",username,password,host,scheme)
	var err error
	Client, err = sql.Open("mysql",dataSourceName)

	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("Database Configured Successfully.")

}