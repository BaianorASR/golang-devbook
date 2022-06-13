package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// StringDatabaseConection is a string that contains the database conection.
	StringDatabaseConection = ""
	// Port is the port of the server.
	Port = 5000
	// SecretKey is a secret key that is used to sign the JWT.
	SecretKey = ""
)

// Init is a function that initializes the config.
func Init() {
	var err error
	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	Port, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		Port = 5000
	}

	SecretKey = os.Getenv("SECRET_KEY")

	StringDatabaseConection = fmt.Sprintf("")

}
