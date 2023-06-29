package controllers

import (
	conf "demoproject/Configs"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Connection, ConnectionError = gorm.Open(mysql.Open(conf.GetDsn()), &gorm.Config{})
