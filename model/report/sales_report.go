package report

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type SalesReportEntity struct {
	ID             bson.ObjectId  `json:"id,omitempty" bson:"_id,omitempty"`
	CreatedAt      time.Time      `json:"created_at" bson:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at" bson:"updated_at"`
	Date           time.Time      `json:"date" bson:"date"`
	PersonName     string         `json:"person_name" bson:"person_name"`
	Item           string         `json:"item" bson:"item"`
	Status         string         `json:"status" bson:"status"`
	Quantity       float64        `json:"quantity" bson:"quantity"`
	Price          float64        `json:"price" bson:"price"`
	PaymentChannel paymentInfo    `json:"payment_channel" bson:"payment_channel"`
	GoFood         gofoodInfo     `json:"gofood" bson:"gofood"`
	HPP            float64        `json:"hpp" bson:"hpp"`
	RealMargin     realMarginInfo `json:"real_margin" bson:"real_margin"`
	Marketing      marketingInfo  `json:"marketing" bson:"markeitng"`
	NettProfit     float64        `json:"nett_profit" bson:"nett_profit"`
	Notes          string         `json:"notes" bson:"notes"`
	Final          finalDetail    `json:"final" bson:"final"`
}

type paymentInfo struct {
	Cash         float64 `json:"cash" bson:"cash"`
	BankTransfer float64 `json:"bank_transfer" bson:"bank_transfer"`
	GoPay        float64 `json:"gopay" bson:"gopay"`
}

type gofoodInfo struct {
	Percentage float64 `json:"percentage" bson:"percentage"`
	Value      float64 `json:"value" bson:"value"`
}

type realMarginInfo struct {
	FinalHPP    float64 `json:"final_hpp" bson:"final_hpp"`
	GrossProfit float64 `json:"gross_profit" bson:"gross_profit"`
}

type marketingInfo struct {
	Percentage float64 `json:"percentage" bson:"percentage"`
	Value      float64 `json:"value" bson:"value"`
}

type finalDetail struct {
	GrossIncome    float64 `json:"gross_income" bson:"gross_income"`
	IncomeCash     float64 `json:"income_cash" bson:"income_cash"`
	IncomeTransfer float64 `json:"income_transfer" bson:"income_transfer"`
	IncomeGoPay    float64 `json:"income_gopay" bson:"income_gopay"`
	HPP            float64 `json:"hpp" bson:"hpp"`
	Marketing      float64 `json:"marketing" bson:"marketing"`
	NettProfit     float64 `json:"nett_profit" bson:"nett_profit"`
}
