package config

import (
	"fmt"
	"os"
)

type Config struct {
	DBHOST string
	DBUSER string
	DBNAME string
	DBPASSWORD string
	DBPORT string
	SSLMODE string
}

func NewDsn()string{
	config := Config{
		DBHOST: os.Getenv("DBHOST"),
		DBUSER: os.Getenv("DBUSER"),
		DBNAME: os.Getenv("DBNAME"),
		DBPASSWORD: os.Getenv("DBPASSWORD"),
		DBPORT: os.Getenv("DBPORT"),
		SSLMODE: os.Getenv("SSLMODE"),
	}
	dsn := fmt.Sprintf("host=%s user=%s port=%s dbname=%s sslmode=%s",config.DBHOST,config.DBUSER,config.DBPORT,config.DBNAME,config.SSLMODE)
	return dsn
}