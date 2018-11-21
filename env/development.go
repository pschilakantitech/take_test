package env

import (
	"time"
)

func init() {

	AppName = "take_test"
	Varsion = "v0.0.1"
	AppEnv = "development"
	ServiceOnPort = ":3000"

	DBHost = "localhost"
	DBPort = "5432"
	DBUser = "postgres"
	DBPassword = "praveensc48"
	DBDatabase = "take_test_db"

	TestTime = time.Minute * 30
}
