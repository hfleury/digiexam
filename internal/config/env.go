package config

import (
	"os"

	"github.com/joho/godotenv"
)

func SetEnv() (map[string]string, error) {
	runEnv := os.Getenv("DIGI_ENVIRONMENT")
	// if local env, get the env vars from the .env file
	if runEnv == "" || runEnv == "LOCAL" {
		err := godotenv.Load(".env")
		if err != nil {
			return nil, err
		}
	} else if runEnv == "TEST" {
		err := godotenv.Load("../../.env")
		if err != nil {
			return nil, err
		}
	}

	m := GetEnv()
	return m, nil
}

func GetEnv() map[string]string {
	m := make(map[string]string)
	// HOST env vars
	m["DIGI_HOST_PORT"] = os.Getenv("DIGI_HOST_PORT")
	m["DIGI_HOST_ADDRESS"] = os.Getenv("DIGI_HOST_ADDRESS")

	// DB env vars
	m["DIGI_POSTGRESQL_PORT"] = os.Getenv("DIGI_POSTGRESQL_PORT")
	m["DIGI_POSTGRESQL_DB"] = os.Getenv("DIGI_POSTGRESQL_DB")
	m["DIGI_POSTGRESQL_PASS"] = os.Getenv("DIGI_POSTGRESQL_PASS")
	m["DIGI_POSTGRESQL_HOST"] = os.Getenv("DIGI_POSTGRESQL_HOST")
	m["DIGI_POSTGRESQL_USER"] = os.Getenv("DIGI_POSTGRESQL_USER")

	return m
}
