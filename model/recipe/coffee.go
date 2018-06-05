package recipe

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type CoffeeEntity struct {
	ID                 bson.ObjectId    `json:"id,omitempty" bson:"_id,omitempty"`
	CreatedAt          time.Time        `json:"created_at" bson:"created_at"`
	UpdatedAt          time.Time        `json:"updated_at" bson:"updated_at"`
	CoffeeName         string           `json:"coffee_name" bson:"coffee_name" validate:"required"`
	Ingredient         []ingredientInfo `json:"ingredient" bson:"ingredient"`
	TotalPricePerGlass float64          `json:"total_price_per_glass" bson:"total_price_per_glass"`
	TotalRRP           float64          `json:"total_rrp" bson:"total_rrp"`
	SellingPrice       float64          `json:"selling_price" bson:"selling_price"`
	ThirdPartyRevenue  []thirdPartyInfo `json:"third_party_revenue" bson:"third_party_revenue"`
	GrossPrice         float64          `json:"gross_price" bson:"gross_price"`
	RealMargin         marginInfo       `json:"real_margin" bson:"real_margin"`
	Marketing          mrktgInfo        `json:"marketing" bson:"marketing"`
	NettProfit         float64          `json:"nett_profit" bson:"nett_profit"`
}

type ingredientInfo struct {
	ID            bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Name          string        `json:"name" bson:"name"`
	Gram          float64       `json:"gram" bson:"gram"`
	Price         float64       `json:"price" bson:"price"`
	Qty           float64       `json:"quantity" bson:"quantity"`
	GramPerGlass  float64       `json:"gram_per_glass" bson:"gram_per_glass"`
	PerGlass      float64       `json:"per_glass" bson:"per_glass"`
	PricePerGlass float64       `json:"price_per_glass" bson:"price_per_glass"`
	TargetMargin  float64       `json:"target_margin" bson:"target_margin"`
	RRP           float64       `json:"rrp" bson:"rrp"`
}

type thirdPartyInfo struct {
	CompanyName string  `json:"company_name" bson:"company_name"`
	Percentage  float64 `json:"percentage" bson:"percentage"`
	Value       float64 `json:"value" bson:"value"`
}

type marginInfo struct {
	HPP       float64 `json:"hpp" bson:"hpp"`
	NettPrice float64 `json:"nett_price" bson:"nett_price"`
}

type mrktgInfo struct {
	Percentage float64 `json:"percentage" bson:"percentage"`
	Value      float64 `json:"value" bson:"value"`
}
