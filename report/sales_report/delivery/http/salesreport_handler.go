package http

import (
	"net/http"

	"github.com/kopi-isi-api-v1/model"
	srModel "github.com/kopi-isi-api-v1/model/report"
	srUsecase "github.com/kopi-isi-api-v1/report"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	validator "gopkg.in/go-playground/validator.v9"
)

type ResponseError struct {
	Message string `json:"message" bson:"message"`
}

type HttpSalesReportHandler struct {
	SalesReportUsecase srUsecase.SalesReportUsecase
}

func (hSr *HttpSalesReportHandler) Save(c echo.Context) error {
	salesReport := new(srModel.SalesReportEntity)
	salesReport.ID = ""

	if err := c.Bind(salesReport); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error)
	}

	if ok, err := isRequestValid(salesReport); !ok || err != nil {
		return c.JSON(http.StatusBadRequest, ResponseError{Message: err.Error()})
	}

	sr, err := hSr.SalesReportUsecase.Save(salesReport)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, sr)
}

func (hSr *HttpSalesReportHandler) FindByID(c echo.Context) error {
	qID := c.Param("id")
	sr, err := hSr.SalesReportUsecase.FindByID(qID)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, sr)
}

func (hSr *HttpSalesReportHandler) FindAll(c echo.Context) error {
	listOfSr, err := hSr.SalesReportUsecase.FindAll()
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, listOfSr)
}

func (hSr *HttpSalesReportHandler) Update(c echo.Context) error {
	srPayload := new(srModel.SalesReportEntity)
	qID := c.Param("id")

	if err := c.Bind(srPayload); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	updatedSr, err := hSr.SalesReportUsecase.Update(qID, srPayload)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, updatedSr)
}

func (hSr *HttpSalesReportHandler) Delete(c echo.Context) error {
	qID := c.Param("id")
	isDeleted, err := hSr.SalesReportUsecase.Delete(qID)
	if err != nil || !isDeleted {
		return c.JSON(http.StatusUnprocessableEntity, "can't process your sales report")
	}

	return c.JSON(http.StatusOK, "your sales report has been deleted")
}

func NewSalesReportHandler(e *echo.Echo, srUsecase srUsecase.SalesReportUsecase) {
	handler := &HttpSalesReportHandler{
		SalesReportUsecase: srUsecase,
	}

	e.POST("/report/sales", handler.Save)
	e.GET("/report/sales/:id", handler.FindByID)
	e.GET("/report/sales", handler.FindAll)
	e.PUT("/report/sales/:id", handler.Update)
	e.DELETE("/report/sales/:id", handler.Delete)
}

func isRequestValid(srUse *srModel.SalesReportEntity) (bool, error) {

	validate := validator.New()

	err := validate.Struct(srUse)
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
