package util

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
)

func GetConfig(key string) string {
	viper.AddConfigPath(".")
	viper.SetConfigFile("/mnt/c/code/Alta/Repo-New/21_Clean and Hexagonal Architecture/praktikum/.env")
	
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("failed reading config: %s", err)
	}

	return viper.GetString(key)
}

func CreateToken(userId uint) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"expire_date": time.Now().Add(time.Hour * 1).Unix(),
	})

	tokenString, err := token.SignedString([]byte(GetConfig("SECRET_JWT_KEY")))

	if err != nil {
		log.Fatalf("failed creating token: %v", err)
	}

	return tokenString
}
