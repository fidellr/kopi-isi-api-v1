package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/kopi-isi-api-v1/config"
	httpIngredientDeliver "github.com/kopi-isi-api-v1/recipe/ingredients/delivery/http"
	ingredientRepos "github.com/kopi-isi-api-v1/recipe/ingredients/repository"
	_ingredientUsecases "github.com/kopi-isi-api-v1/recipe/ingredients/usecase"

	httpCoffeeDeliver "github.com/kopi-isi-api-v1/recipe/coffee/delivery/http"
	coffeeRepos "github.com/kopi-isi-api-v1/recipe/coffee/repository"
	_coffeeUsecases "github.com/kopi-isi-api-v1/recipe/coffee/usecase"
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

	ingredientRepo := ingredientRepos.NewIngredientsMongo(db, "ingredients")
	ingredientUsecase := _ingredientUsecases.NewIngredientUsecase(ingredientRepo)
	httpIngredientDeliver.NewIngredientHttpHandler(e, ingredientUsecase)

	coffeeRepo := coffeeRepos.NewCoffeeMongo(db, "coffee")
	coffeeUsecase := _coffeeUsecases.NewCoffeeUsecase(coffeeRepo)
	httpCoffeeDeliver.NewCoffeeHttpHandler(e, coffeeUsecase)

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
