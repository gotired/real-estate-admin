package model

type CONFIG struct {
	APP APP
	DB  DATABASE
}

type APP struct {
	Port string
}

type DATABASE struct {
	USERNAME string
	PASSWORD string
	DB_NAME  string
	HOST     string
	PORT     int
}
