package users_db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	mysqlUsersUsername = "root"
	mysqlUsersPassword = "13700508S@m"
	mysqlUsersHost     = "127.0.0.1:3306"
	mysqlUsersSchema   = "users_db"
)

var (
	Client *sql.DB
	// username = os.Getenv(mysqlUsersUsername)
	// password = os.Getenv(mysqlUsersPassword)
	// host     = os.Getenv(mysqlUsersHost)
	// schema   = os.Getenv(mysqlUsersSchema)
)

func init() {

	datasourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb3",
		mysqlUsersUsername, // "root",           //username
		mysqlUsersPassword, // "13700508S@m",    //password
		mysqlUsersHost,     // "localhost:3306", //localhost
		mysqlUsersSchema,   // "users_db",       //Schema-name or db-name
	)
	fmt.Println("DSN:", datasourceName) // Print DSN for debugging

	// log.Println(fmt.Sprintf("about to connect to %s", datasourceName))
	var err error
	Client, err = sql.Open("mysql", datasourceName)
	if err != nil {
		panic(err)
	}

	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("database successfully configured")
}
