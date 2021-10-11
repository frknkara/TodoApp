package data

import (
	"fmt"
	"os"
)

var (
	dbUsername = "root"
	dbPassword = "root"
	dbHost     = getHostNameFromEnv()
	dbPort     = "3306"
	dbName     = "tododb"
	dbDSN      = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUsername, dbPassword, dbHost, dbPort, dbName)
)

func getHostNameFromEnv() string {
	if value, ok := os.LookupEnv("MYSQL_HOST_NAME"); ok {
		return value
	}
	return "localhost"
}
