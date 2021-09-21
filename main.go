package main

import (
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Resp struct {
	Message string `json:"message"`
}

type ErrorResp struct {
	Error string `json:"error"`
}

var (
	stayTime = 0
)

func GetWait(c echo.Context) error {
	if stayTime >= 0 {
		time.Sleep(time.Duration(stayTime) * time.Second)
	}

	resp := Resp{
		Message: "bow-wow!",
	}
	return c.JSON(http.StatusOK, resp)
}

func GetWaitTime(c echo.Context) error {
	waitTime := c.Param("time")

	i, err := strconv.ParseInt(waitTime, 10, 64)
	if err != nil {
		return err
	}
	time.Sleep(time.Duration(i) * time.Second)

	resp := Resp{
		Message: "bow-wow!",
	}
	return c.JSON(http.StatusOK, resp)
}

func GetWaitRandom(c echo.Context) error {
	rand.Seed(time.Now().UnixNano())

	ra := rand.Int63n(10) // 1 ~ 10
	time.Sleep(time.Duration(ra) * time.Second)

	resp := Resp{
		Message: "bow-wow!",
	}
	return c.JSON(http.StatusOK, resp)
}

type ParamPostSay struct {
	T int `json:"time"`
}

func PostSay(c echo.Context) error {
	t := c.Param("time")
	i, err := strconv.Atoi(t)
	if err != nil {
		return err
	}

	stayTime = i
	resp := Resp{
		Message: "bow?",
	}
	return c.JSON(http.StatusOK, resp)
}

func index(c echo.Context) error {
	h := struct {
		Health string `json:"health"`
	}{
		Health: "ok",
	}
	return c.JSON(http.StatusOK, h)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	e := echo.New()

	e.Use(middleware.Logger())

	e.GET("/", index)
	e.GET("/wait", GetWait)
	e.GET("/wait/random", GetWaitRandom)
	e.GET("/wait/:time", GetWaitTime)
	e.POST("/say/:time", PostSay)

	e.Logger.Fatal(e.Start(":" + port))
}
