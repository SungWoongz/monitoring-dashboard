package variables

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	ADDR     = "0.0.0.0"
	PORT     = "8080"
	DATABASE = "./db/serverstatus.db"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Print("Error loading .env file")
	}
	ADDR = getEnv("ADDR", ADDR)
	PORT = getEnv("PORT", PORT)
	DATABASE = getEnv("DATABASE", DATABASE)
}

func getEnv(k string, org string) string {
	v, isExist := os.LookupEnv(k)
	if !isExist {
		return org
	}
	return v
}
