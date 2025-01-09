package variables

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	ADDR         = "0.0.0.0"
	PORT         = ""
	DATABASE_URL = ""
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Print("Error loading .env file")
	}
	ADDR = os.Getenv("ADDR")
	PORT = os.Getenv("PORT")
}
