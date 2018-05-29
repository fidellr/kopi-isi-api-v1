package http

import (
	"net/http"

	"github.com/kopi-isi-api-v1/model"
	ingModel "github.com/kopi-isi-api-v1/model/recipe"
	ingUsecase "github.com/kopi-isi-api-v1/recipe"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	validator "gopkg.in/go-playground/validator.v9"
)

type ResponseError struct {
	Message string `json:"message" bson:"message"`
}

type HttpIngridientHandler struct {
	IngridientUsecase ingUsecase.IngridientsUsecase
}

func (hIng *HttpIngridientHandler) Save(c echo.Context) error {
	ingridient := new(ingModel.IngridientMaster)
	ingridient.ID = ""

	if err := c.Bind(ingridient); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := isRequestValid(ingridient); !ok || err != nil {
		return c.JSON(http.StatusBadRequest, ResponseError{Message: err.Error()})
	}

	ing, existedIng, err := hIng.IngridientUsecase.Save(ingridient)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	if existedIng != nil {
		return c.JSON(http.StatusConflict, existedIng)
	}

	return c.JSON(http.StatusCreated, ing)
}

func (hIng *HttpIngridientHandler) FindByID(c echo.Context) error {
	qID := c.Param("id")
	ing, err := hIng.IngridientUsecase.FindByID(qID)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, ing)
}

func (hIng *HttpIngridientHandler) FindAll(c echo.Context) error {
	listOfIng, err := hIng.IngridientUsecase.FindAll()
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, listOfIng)
}

func (hIng *HttpIngridientHandler) Update(c echo.Context) error {
	ingridientPayload := new(ingModel.IngridientUpdate)
	qID := c.Param("id")
	_, err := hIng.IngridientUsecase.FindByID(qID)
	if err != nil {
		return c.JSON(http.StatusNotFound, ResponseError{Message: err.Error()})
	}

	if err := c.Bind(ingridientPayload); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error)
	}

	updatedIng, err := hIng.IngridientUsecase.Update(qID, ingridientPayload)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, updatedIng)
}

func (hIng *HttpIngridientHandler) Delete(c echo.Context) error {
	qID := c.Param("id")
	_, err := hIng.IngridientUsecase.FindByID(qID)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	isDeleted, err := hIng.IngridientUsecase.Delete(qID)
	if err != nil || !isDeleted {
		return c.JSON(http.StatusUnprocessableEntity, "can't process your ingridient")
	}

	return c.JSON(http.StatusOK, "your ingridient deleted")

}

func NewIngridientHttpHandler(e *echo.Echo, ingUsecase ingUsecase.IngridientsUsecase) {
	handler := &HttpIngridientHandler{
		IngridientUsecase: ingUsecase,
	}
	e.POST("/recipe/ingridient", handler.Save)
	e.GET("/recipe/ingridient/:id", handler.FindByID)
	e.GET("/recipe/ingridient", handler.FindAll)
	e.PUT("/recipe/ingridient/:id", handler.Update)
	e.DELETE("/recipe/ingridient/:id", handler.Delete)
}

func isRequestValid(ingUse *ingModel.IngridientMaster) (bool, error) {

	validate := validator.New()

	err := validate.Struct(ingUse)
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
