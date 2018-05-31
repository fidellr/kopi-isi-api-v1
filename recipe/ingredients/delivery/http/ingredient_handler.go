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

type HttpIngredientHandler struct {
	IngredientUsecase ingUsecase.IngredientsUsecase
}

func (hIng *HttpIngredientHandler) Save(c echo.Context) error {
	ingredient := new(ingModel.IngredientMaster)
	ingredient.ID = ""

	if err := c.Bind(ingredient); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := isRequestValid(ingredient); !ok || err != nil {
		return c.JSON(http.StatusBadRequest, ResponseError{Message: err.Error()})
	}

	ing, existedIng, err := hIng.IngredientUsecase.Save(ingredient)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	if existedIng != nil {
		return c.JSON(http.StatusConflict, existedIng)
	}

	return c.JSON(http.StatusCreated, ing)
}

func (hIng *HttpIngredientHandler) FindByID(c echo.Context) error {
	qID := c.Param("id")
	ing, err := hIng.IngredientUsecase.FindByID(qID)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, ing)
}

func (hIng *HttpIngredientHandler) FindAll(c echo.Context) error {
	listOfIng, err := hIng.IngredientUsecase.FindAll()
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, listOfIng)
}

func (hIng *HttpIngredientHandler) Update(c echo.Context) error {
	ingredientPayload := new(ingModel.IngredientUpdate)
	qID := c.Param("id")

	if err := c.Bind(ingredientPayload); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	updatedIng, err := hIng.IngredientUsecase.Update(qID, ingredientPayload)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, updatedIng)
}

func (hIng *HttpIngredientHandler) Delete(c echo.Context) error {
	qID := c.Param("id")
	isDeleted, err := hIng.IngredientUsecase.Delete(qID)
	if err != nil || !isDeleted {
		return c.JSON(http.StatusUnprocessableEntity, "can't process your ingredient")
	}

	return c.JSON(http.StatusOK, "your ingredient has been deleted")

}

func NewIngredientHttpHandler(e *echo.Echo, ingUsecase ingUsecase.IngredientsUsecase) {
	handler := &HttpIngredientHandler{
		IngredientUsecase: ingUsecase,
	}
	e.POST("/recipe/ingredient", handler.Save)
	e.GET("/recipe/ingredient/:id", handler.FindByID)
	e.GET("/recipe/ingredient", handler.FindAll)
	e.PUT("/recipe/ingredient/:id", handler.Update)
	e.DELETE("/recipe/ingredient/:id", handler.Delete)
}

func isRequestValid(ingUse *ingModel.IngredientMaster) (bool, error) {

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
