package services

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"net/url"
	"submit/logger"
)

var database *gorm.DB

type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func Database() *gorm.DB {
	return database
}

func NewDatabase() error {
	var err error = nil
	database, err = gorm.Open("postgres", dbURL())
	if err != nil {
		return err
	}

	database.AutoMigrate(User{}, Course{}, Submission{}, Assignment{}, File{})
	if err != nil {
		fmt.Println(err)
	}


	logger.Info("PostgreSQL connection has been established.")
	return nil
}

func TerminateDatabase() {
	_ = database.Close()
	logger.Info("PostgreSQL connection has been discontinued.")
}

func dbURL() string {
	dbConfig := DBConfig{
		Host:     "localhost",
		Port:     5432,
		User:     "turnindev",
		DBName:   "turnin",
		Password: "toor",
	}
	u := url.URL{
		User:     url.UserPassword(dbConfig.User, dbConfig.Password),
		Scheme:   "postgres",
		Host:     fmt.Sprintf("%s:%d", dbConfig.Host, dbConfig.Port),
		Path:     dbConfig.DBName,
		RawQuery: (&url.Values{"sslmode": []string{"disable"}}).Encode(),
	}
	return u.String()
}
