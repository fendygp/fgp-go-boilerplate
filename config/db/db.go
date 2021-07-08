package db

import (
	"errors"
	"fmt"
	"os"

	"bitbucket.org/bridce/fgp-go-boilerplate/config"
	"github.com/jinzhu/gorm"
	"github.com/qor/validations"
)

// DB Global DB connection
var DB *gorm.DB

func init() {
	if DB != nil {
		return
	}
	var err error

	dbConfig := config.Config.DB

	if dbConfig.Adapter == "mysql" {
		DB, err = gorm.Open("mysql", fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=True&loc=Local", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name))
	} else if dbConfig.Adapter == "cockroach" {

		if dbConfig.SslMode == "disable" {
			DB, err = gorm.Open("postgres", fmt.Sprintf("postgresql://%v%v@%v:%v/%v?application_name=cockroach&sslmode=disable", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name))
		} else if dbConfig.SslMode == "verify-ca" {
			DB, err = gorm.Open("postgres", fmt.Sprintf("postgresql://%v%v@%v:%v/%v?application_name=cockroach&sslmode=verify-ca&sslrootcert=certs/ca.crt", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name))
		} else {
			panic(fmt.Errorf("Unsupported SSL Mode: ", dbConfig.SslMode))
		}

	} else if dbConfig.Adapter == "postgres" {

		if dbConfig.SslMode == "disable" {
			DB, err = gorm.Open("postgres", fmt.Sprintf("postgres://%v:%v@%v/%v?sslmode=disable", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Name))
		} else {
			panic(fmt.Errorf("Unsupported SSL Mode: ", dbConfig.SslMode))
		}

	} else if dbConfig.Adapter == "sqlite" {
		DB, err = gorm.Open("sqlite3", fmt.Sprintf("%v/%v", os.TempDir(), dbConfig.Name))
	} else {
		panic(errors.New("not supported database adapter"))
	}

	if err == nil {
		if os.Getenv("DEBUG") == "TRUE" {
			DB.LogMode(true)
		}

		validations.RegisterCallbacks(DB)
	} else {
		panic(err)
	}

}
