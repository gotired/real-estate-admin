package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gotired/real-estate-admin/server/internal/domain/model"
)

func Init() *model.CONFIG {

	appConfig := model.APP{
		Port: getEnv("PORT", false, "3000"),
	}

	dbPort, err := strconv.Atoi(getEnv("POSTGRES_PORT", true))
	if err != nil {
		panic(err)
	}

	dbConfig := model.DATABASE{
		USERNAME: getEnv("POSTGRES_USERNAME", true),
		PASSWORD: getEnv("POSTGRES_PASSWORD", true),
		DB_NAME:  getEnv("POSTGRES_DATABASE", true),
		HOST:     getEnv("POSTGRES_HOST", true),
		PORT:     dbPort,
	}

	return &model.CONFIG{
		APP: appConfig,
		DB:  dbConfig,
	}
}

func getEnv(key string, strict bool, defaultValue ...string) string {
	defVal := ""
	if len(defaultValue) > 0 {
		defVal = defaultValue[0]
	}

	value := os.Getenv(key)
	if value != "" {
		return value
	}

	if !strict {
		return defVal
	}
	panic(fmt.Sprintf("Error loading %s value", key))
}
