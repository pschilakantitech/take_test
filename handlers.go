package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func assignHandlers() {
	e.GET("/", welcomes)
}

func welcomes(c echo.Context) error {

	return c.String(http.StatusOK, "praveen chbshcksdbckjsdcnk")
}
