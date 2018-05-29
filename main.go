package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/kopi-isi-api-v1/config"
	httpIngridientDeliver "github.com/kopi-isi-api-v1/recipe/ingridients/delivery/http"
	ingridientRepos "github.com/kopi-isi-api-v1/recipe/ingridients/repository"
	_ingridientUsecases "github.com/kopi-isi-api-v1/recipe/ingridients/usecase"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sirupsen/logrus"
)

func main() {
	db, session, err := config.MongoConfig()
	if err != nil {
		fmt.Println(err)
	}
	defer session.Close()

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.POST, echo.DELETE},
	}))

	ingridientRepo := ingridientRepos.NewIngridientsMongo(db, "ingridients")
	ingridientUsecase := _ingridientUsecases.NewIngridientUsecase(ingridientRepo)
	httpIngridientDeliver.NewIngridientHttpHandler(e, ingridientUsecase)

	port := os.Getenv("PORT")
	if port == "" {
		logrus.Fatal("no port provided")
	}
	e.GET("/", ping)
	e.Start(":" + port)

}

func ping(c echo.Context) error {
	return c.JSON(http.StatusOK, "pong")
}
