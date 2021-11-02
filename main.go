package main

import (
	"fmt"
	"log"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var err error

var mySigningKey = []byte("mysupersecretkey")

func generateJWT() (string error) {
	token := jwt.New(jwt.SigningMethodES256)

	claims := token.Claims(jwt.MapClaims)

	claims["authorized"] = "true"
	claims["user"] = "test"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("Something went wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}

func main() {
	db, err = gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/golang-crud?charset=utf8mb4&parseTime=True&loc=Local")

	if err != nil {
		log.Println("Connection failed", err)
	} else {
		log.Println("Connected")
	}
}
