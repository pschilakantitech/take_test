package main

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"github.com/pschilakantitech/avitar/log"
	"github.com/pschilakantitech/avitar/pidfile"
	"github.com/pschilakantitech/take_test/env"
	"github.com/pschilakantitech/take_test/pg_persist"
)

func doCommonSetUp() {
	initLogging()

	if err := pidfile.Dump(); err != nil {
		log.Error("Unable to create pid file, monitoring may be affected...", err)
	}
	cfg := pg_persist.Config{
		Host:     env.DBHost,
		Port:     env.DBPort,
		User:     env.DBUser,
		Password: env.DBPassword,
		Database: env.DBDatabase,
	}
	fmt.Println("env.DBPassword", env.DBPassword)
	if err := pg_persist.ConnectToPGDB(cfg); err != nil {
		abort(err)
	}
	fmt.Println("\nConnected to DB...", cfg.Database)

	rand.Seed(time.Now().UTC().UnixNano())
}

const (
	Success = iota
	SetupFailed
)

func initLogging() {
	logFile := filepath.Join("log", env.AppName+"_"+env.Varsion+"_"+env.AppEnv+".log")
	logCfg := log.Config{
		LogPrefix: env.AppName,
		LogName:   logFile,
		Debug:     false,
		AppName:   env.AppName,
		AppEnv:    env.AppEnv,
	}
	if err := log.Setup(logCfg); err != nil {
		abort(err)
	}
}

func abort(msg error) {
	log.Error(msg)
	fmt.Println(msg)
	pidfile.Drop()
	os.Exit(SetupFailed)
}
