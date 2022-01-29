package config

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort   string
	DbConfig  DatabaseConfig
	RdsConfig RedisConfig
	JwtConfig
	IsAdmin     bool
	AdminRoleID int
	UserRoleID  int
}

type DatabaseConfig struct {
	DbDriver   string
	DbHost     string
	DbPort     string
	DbUsername string
	DbPassword string
	DbSchema   string
}

type RedisConfig struct {
	Host string
	Port string
}

type JwtConfig struct {
	RefreshSecret string
	AccessSecret  string
}

var (
	_, b, _, _ = runtime.Caller(0)
	// root folder of this project
	root     = filepath.Join(filepath.Dir(b), "../..")
	filename = ".env"
)

//LoadEnv exported
func LoadEnv() {
	p := root + "/" + filename
	e := godotenv.Load(p)
	if e != nil {
		panic(e.Error())
	}
}

func PrepareConfigurations() *Config {
	LoadEnv()
	return &Config{
		AppPort: os.Getenv("APP_PORT"),
		DbConfig: DatabaseConfig{
			DbDriver:   os.Getenv("DB_DRIVER"),
			DbHost:     os.Getenv("DB_HOST"),
			DbPort:     os.Getenv("DB_PORT"),
			DbUsername: os.Getenv("DB_USER"),
			DbPassword: os.Getenv("DB_PASSWORD"),
			DbSchema:   os.Getenv("DB_DATABASE"),
		},
		RdsConfig: RedisConfig{
			Host: os.Getenv("REDIS_HOST"),
			Port: os.Getenv("REDIS_PORT"),
		},
		JwtConfig: JwtConfig{
			RefreshSecret: os.Getenv("REFRESH_SECRET"),
			AccessSecret:  os.Getenv("ACCESS_SECRET"),
		},
	}
}
