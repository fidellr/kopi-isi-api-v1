package report

import "github.com/kopi-isi-api-v1/model/report"

type SalesReport interface {
	Save(srPayload *report.SalesReportEntity) (*report.SalesReportEntity, error)
	FindByID(id string) (*report.SalesReportEntity, error)
	FindAll() ([]*report.SalesReportEntity, error)
	Update(id string, srPayload *report.SalesReportEntity) (*report.SalesReportEntity, error)
	Delete(id string) (bool, error)
}
