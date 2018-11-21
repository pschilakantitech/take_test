package env

import (
	"time"
)

// Application related variables...
var (
	AppName       string
	Varsion       string
	AppEnv        string
	ServiceOnPort string
)

// pg DB variables...
var (
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBDatabase string
)

// test variables...
var (
	TestTime time.Duration
)
