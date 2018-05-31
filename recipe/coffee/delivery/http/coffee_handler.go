package http

import (
	"net/http"

	"github.com/kopi-isi-api-v1/model"
	cofModel "github.com/kopi-isi-api-v1/model/recipe"
	cofUsecase "github.com/kopi-isi-api-v1/recipe"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	validator "gopkg.in/go-playground/validator.v9"
)

type ResponseError struct {
	Message string `json:"message" bson:"message"`
}

type HttpCoffeeHandler struct {
	CoffeeUsecase cofUsecase.CoffeeUsecase
}

func (hCof *HttpCoffeeHandler) Save(c echo.Context) error {
	coffee := new(cofModel.CoffeeEntity)
	coffee.ID = ""

	if err := c.Bind(coffee); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error)
	}

	if ok, err := isRequestValid(coffee); !ok || err != nil {
		return c.JSON(http.StatusBadRequest, ResponseError{Message: err.Error()})
	}

	cof, existedCof, err := hCof.CoffeeUsecase.Save(coffee)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	if existedCof != nil {
		return c.JSON(http.StatusConflict, existedCof)
	}

	return c.JSON(http.StatusCreated, cof)

}

func (hCof *HttpCoffeeHandler) FindByID(c echo.Context) error {
	qID := c.Param("id")
	cof, err := hCof.CoffeeUsecase.FindByID(qID)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, cof)
}

func (hCof *HttpCoffeeHandler) FindAll(c echo.Context) error {
	listOfCof, err := hCof.CoffeeUsecase.FindAll()
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, listOfCof)
}

func (hCof *HttpCoffeeHandler) Update(c echo.Context) error {
	coffeePayload := new(cofModel.CoffeeEntity)
	qID := c.Param("id")

	if err := c.Bind(coffeePayload); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	updatedCof, err := hCof.CoffeeUsecase.Update(qID, coffeePayload)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, updatedCof)
}

func (hCof *HttpCoffeeHandler) Delete(c echo.Context) error {
	qID := c.Param("id")
	isDeleted, err := hCof.CoffeeUsecase.Delete(qID)
	if err != nil || !isDeleted {
		return c.JSON(http.StatusUnprocessableEntity, "can't process your coffee")
	}

	return c.JSON(http.StatusOK, "your coffee has been deleted")
}

func NewCoffeeHttpHandler(e *echo.Echo, cofUsecase cofUsecase.CoffeeUsecase) {
	handler := &HttpCoffeeHandler{
		CoffeeUsecase: cofUsecase,
	}
	e.POST("/recipe/coffee", handler.Save)
	e.GET("/recipe/coffee/:id", handler.FindByID)
	e.GET("/recipe/coffee", handler.FindAll)
	e.PUT("/recipe/coffee/:id", handler.Update)
	e.DELETE("/recipe/coffee/:id", handler.Delete)
}

func isRequestValid(cofStruct *cofModel.CoffeeEntity) (bool, error) {

	validate := validator.New()

	err := validate.Struct(cofStruct)
	if err != nil {
		return false, err
	}
	return true, nil
}

func getStatusCode(err error) int {

	if err == nil {
		return http.StatusOK
	}

	logrus.Error(err)

	switch err {
	case model.INTERNAL_SERVER_ERROR:
		return http.StatusInternalServerError
	case model.NOT_FOUND_ERROR:
		return http.StatusNotFound
	case model.CONFLICT_ERROR:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
