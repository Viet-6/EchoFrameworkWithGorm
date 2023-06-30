package configs

import (
	"fmt"

	"github.com/joho/godotenv"
)

var username string = "root"
var password string = "root"
var host string = "mysql"
var port string = "3306"
var dbname string = "typing_tournament"

func GetDsn() (dsn string) {

	err := godotenv.Load()

	fmt.Print()
	dsn = fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username,
		password,
		host,
		port,
		dbname,
	)

	return
}
