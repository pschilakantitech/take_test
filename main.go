package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/pschilakantitech/avitar/pidfile"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/pschilakantitech/avitar/log"
	"github.com/pschilakantitech/take_test/env"
)

var e *echo.Echo

func main() {
	sig, quit := make(chan os.Signal), make(chan bool)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	e = echo.New()
	e.Use(middleware.Secure())
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	assignHandlers()

	fmt.Println("Starting echo.labstack.com server...")
	go func() {
		if err := e.Start(env.ServiceOnPort); err != nil {
			log.Info("got error,shutting down the server", err)
		}
	}()

	fmt.Println("Ready to serve the requests on the port", env.ServiceOnPort)
	fmt.Println("Setup OK.\nRunning... ")

	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		log.Info(err)
	}
	fmt.Println("shutting down the server... Done")

	pidfile.Drop()

	fmt.Println("Stopping goroutines... Done")
	fmt.Println("OK")
	log.Info("Shutdown OK")
	fmt.Println("All done. Bye Bye...")
	os.Exit(0)
}

func init() {
	doCommonSetUp()
}
