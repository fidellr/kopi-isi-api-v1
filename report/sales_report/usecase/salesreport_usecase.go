package usecase

import (
	"math"

	model "github.com/kopi-isi-api-v1/model/report"
	"github.com/kopi-isi-api-v1/report"
	"github.com/kopi-isi-api-v1/validator"
)

type salesReportUsecase struct {
	srRepos report.SalesReport
}

func NewSalesReportUsecase(sr report.SalesReport) report.SalesReport {
	return &salesReportUsecase{
		srRepos: sr,
	}
}

func (srUse *salesReportUsecase) Save(srPayload *model.SalesReportEntity) (*model.SalesReportEntity, error) {
	if err := validator.Validate(srPayload); err != nil {
		return nil, report.NewErrorInvalidReportData(err.Error())
	}

	// Sales cash calculations
	srPayC := srPayload.PaymentChannel
	srPayC.Cash = math.Round(srPayload.Price * srPayload.Quantity)

	// Sales Go food calculations
	srGF := srPayload.GoFood
	srGF.Value = math.Round(srPayC.Cash * srGF.Percentage / 100)

	// Fix Gross Income that should be the result of SUM Cash and Gofood value
	srPayload.RealMargin.FinalHPP = srPayload.Final.GrossIncome

	sr, err := srUse.srRepos.Save(srPayload)
	if err != nil {
		return nil, err
	}

	return sr, nil
}

func (srUse *salesReportUsecase) FindByID(id string) (*model.SalesReportEntity, error) {
	sr, err := srUse.srRepos.FindByID(id)
	if err != nil {
		return nil, err
	}

	return sr, nil
}

func (srUse *salesReportUsecase) FindAll() ([]*model.SalesReportEntity, error) {
	listOfSr, err := srUse.srRepos.FindAll()
	if err != nil {
		return nil, err
	}

	return listOfSr, nil
}

func (srUse *salesReportUsecase) Update(id string, srPayload *model.SalesReportEntity) (*model.SalesReportEntity, error) {
	sr, err := srUse.srRepos.Update(id, srPayload)
	if err != nil {
		return nil, err
	}

	return sr, nil
}

func (srUse *salesReportUsecase) Delete(id string) (bool, error) {
	_, err := srUse.srRepos.Delete(id)
	if err != nil {
		return false, err
	}

	return true, nil
}
