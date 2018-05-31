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
	TotalPricePerGlass int64            `json:"total_price_per_glass"`
	TotalRRP           int64            `json:"total_rrp" bson:"total_rrp"`
	ThirdPartyRevenue  []thirdPartyInfo `json:"third_party_revenue" bson:"third_party_revenue"`
	RealMargin         marginInfo       `json:"real_margin" bson:"real_margin"`
	Marketing          mrktgInfo        `json:"marketing" bson:"marketing"`
}

type ingredientInfo struct {
	ID            bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Name          string        `json:"ingredient_name" bson:"ingredient_name"`
	Qty           int64         `json:"quantity" bson:"quantity"`
	GramPerGlass  float64       `json:"gram_per_glass" bson:"gram_per_glass"`
	PerGlass      int64         `json:"per_glass" bson:"per_glass"`
	PricePerGlass float64       `json:"price_per_glass" bson:"price_per_glass"`
	TargetMargin  float64       `json:"target_margin" bson:"target_margin"`
	RRP           int64         `json:"rrp" bson:"rrp"`
}

type thirdPartyInfo struct {
	CompanyName  string  `json:"company_name" bson:"company_name"`
	SellingPrice int64   `json:"selling_price" bson:"selling_price"`
	Percentage   float64 `json:"percentage" bson:"percentage"`
	Value        int64   `json:"value" bson:"value"`
	GrossPrice   int64   `json:"gross_price" bson:"gross_price"`
}

type marginInfo struct {
	HPP       int64 `json:"hpp" bson:"hpp"`
	NettPrice int64 `json:"nett_price" bson:"nett_price"`
}

type mrktgInfo struct {
	Percentage float64 `json:"percentage" bson:"percentage"`
	Value      int64   `json:"value" bson:"value"`
	NettProfit int64   `json:"nett_profit" bson:"nett_profit"`
}
