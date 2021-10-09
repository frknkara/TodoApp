package dcl

import "fmt"

var (
	dbUsername = "root"
	dbPassword = "root"
	dbHost     = "localhost"
	dbPort     = "3306"
	dbName     = "tododb"
	dbDSN      = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUsername, dbPassword, dbHost, dbPort, dbName)
)
