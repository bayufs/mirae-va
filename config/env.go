package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var SERVICE_HOST string
var SERVICE_PORT string
var SERVICE_ROOT_PATH string

var DB_USER string
var DB_PASS string
var DB_NAME string
var DB_HOST string
var DB_PORT string

func init() {
	LoadEnv()
}

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
		log.Fatal("Error loading .env file")
		return
	}
	InitEnvVariables()
	log.Println("Success load .env file")
}

func InitEnvVariables() {
	setEnv(&SERVICE_HOST, "SERVICE_HOST", "0.0.0.0")
	setEnv(&SERVICE_PORT, "SERVICE_PORT", "8000")
	setEnv(&SERVICE_ROOT_PATH, "SERVICE_ROOT_PATH", "va")

	setEnv(&DB_USER, "DB_USER")
	setEnv(&DB_PASS, "DB_PASS")
	setEnv(&DB_NAME, "DB_NAME")
	setEnv(&DB_HOST, "DB_HOST")
	setEnv(&DB_PORT, "DB_PORT")
}

func setEnv(envVariable interface{}, envName string, defaultValue ...interface{}) {
	var envValue = os.Getenv(envName)
	switch envVariable.(type) {
	case *string:
		*envVariable.(*string) = envValue
		if *envVariable.(*string) == "" {
			if len(defaultValue) > 0 {
				*envVariable.(*string) = defaultValue[0].(string)
			}
		}
	case *[]byte:
		if envValue == "" {
			if len(defaultValue) > 0 {
				*envVariable.(*[]byte) = defaultValue[0].([]byte)
			} else {
				*envVariable.(*[]byte) = []byte("")
			}
		} else {
			*envVariable.(*[]byte) = []byte(envValue)
		}
	case *int:
		if envValue == "" {
			if len(defaultValue) > 0 {
				*envVariable.(*int) = defaultValue[0].(int)
			} else {
				*envVariable.(*int) = 0
			}
		} else {
			newInt, _ := strconv.Atoi(os.Getenv(envName))
			*envVariable.(*int) = newInt
		}
	case *bool:
		if envValue == "" {
			if len(defaultValue) > 0 {
				*envVariable.(*bool) = defaultValue[0].(bool)
			} else {
				*envVariable.(*bool) = false
			}
		} else {
			if envValue == "1" {
				*envVariable.(*bool) = true
			} else {
				*envVariable.(*bool) = false
			}
		}
	}
}
