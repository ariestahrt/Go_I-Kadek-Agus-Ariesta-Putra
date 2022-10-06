package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Env(name string) string {
    err := godotenv.Load("C:\\code\\Alta\\Repo-New\\20_Unit Testing\\praktikum\\problem2\\.env")
    if err != nil {
        log.Fatal(err)
    }
  
    return os.Getenv(os.Getenv("ENV")+"_"+name)
}