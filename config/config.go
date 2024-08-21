package config

import "rest-api-echo/utils"

type Configuration struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_DATABASE string
	DB_HOST     string
	DB_PORT     string
}

func GetConfig() Configuration {
	conf := Configuration{}
	conf.DB_USERNAME = utils.GetEnv("DB_USERNAME")
	conf.DB_PASSWORD = utils.GetEnv("DB_PASSWORD")
	conf.DB_DATABASE = utils.GetEnv("DB_DATABASE")
	conf.DB_HOST = utils.GetEnv("DB_HOST")
	conf.DB_PORT = utils.GetEnv("DB_PORT")

	return conf
}
