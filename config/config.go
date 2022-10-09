package config

import "os"

func Config() DBConfig {
	return DBConfig{
		"host":     os.Getenv("DB_HOST"),
		"user":     os.Getenv("DB_USER"),
		"password": os.Getenv("DB_PASSWORD"),
		"dbName":   os.Getenv("DB_DATABASE"),
		"sslMode":  os.Getenv("DB_SSLMODE"),
	}
}
