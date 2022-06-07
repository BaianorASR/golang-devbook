package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	StringDatabaseConection = ""
	Port                    = 5000
)

func Init() {
	var err error
	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	Port, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		Port = 5000
	}

	StringDatabaseConection = fmt.Sprintf("")

}

// StringDatabaseConection = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
// os.Getenv("MYSQL_USER"),
// os.Getenv("MYSQL_PASSWORD"),
// os.Getenv("MYSQL_DATABASE"),
// )

// StringDatabaseConection = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
// os.Getenv("MYSQL_USER"),
// os.Getenv("MYSQL_PASSWORD"),
// os.Getenv("MYSQL_HOST"),
// os.Getenv("MYSQL_PORT"),
// os.Getenv("MYSQL_DATABASE"),
// )
