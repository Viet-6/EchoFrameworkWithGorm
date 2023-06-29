package configs

import "fmt"

var username string = "root"
var password string = "mydb_p@ssw0rd"
var host string = "localhost"
var port string = "3306"
var dbname string = "typing_tournament"

func GetDsn() (dsn string) {
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
