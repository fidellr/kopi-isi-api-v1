package usecase

import (
	"fmt"
	"math"

	model "github.com/kopi-isi-api-v1/model/recipe"
	"github.com/kopi-isi-api-v1/recipe"
	"github.com/kopi-isi-api-v1/validator"
)

type coffeeUsecase struct {
	cofRepos recipe.Coffee
}

func NewCoffeeUsecase(cof recipe.Coffee) recipe.Coffee {
	return &coffeeUsecase{
		cofRepos: cof,
	}
}

func (cofUse *coffeeUsecase) Save(cofPayload *model.CoffeeEntity) (*model.CoffeeEntity, map[string]string, error) {
	if errValid := validator.Validate(cofPayload); errValid != nil {
		return nil, nil, recipe.NewErrorInvalidRecipeData(errValid.Error())
	}

	// Ingredient calculations
	cing := cofPayload.Ingredient
	var totalRrp float64
	for k := range cing {
		if cing[k].PerGlass == 0 {
			cing[k].PerGlass = math.Round(cing[k].Gram / cing[k].GramPerGlass)
		}

		cing[k].PricePerGlass = math.Round(cing[k].Price / cing[k].PerGlass)
		cing[k].RRP = math.Round(cing[k].PricePerGlass * cing[k].TargetMargin / 100)
		cofPayload.TotalPricePerGlass += cing[k].PricePerGlass
		totalRrp += cing[k].RRP
	}
	cofPayload.TotalRRP = totalRrp

	// ThirdPartyRevenue calculations
	cThirdRev := cofPayload.ThirdPartyRevenue
	var tpVal float64
	for k := range cThirdRev {
		tpVal = cofPayload.SellingPrice * cThirdRev[k].Percentage / 100
		cThirdRev[k].Value = tpVal
		cofPayload.GrossPrice = math.Round(cofPayload.SellingPrice - (cofPayload.SellingPrice * cThirdRev[k].Percentage / 100))
	}

	// RealMargin nett price calculations
	cofPayload.RealMargin.NettPrice = cofPayload.GrossPrice - cofPayload.RealMargin.HPP

	// Marketing value calculations
	cofPayload.Marketing.Value = math.Round(cofPayload.RealMargin.NettPrice * cofPayload.Marketing.Percentage / 100)

	// NettProfit calculations
	cofPayload.NettProfit = math.Round(cofPayload.SellingPrice - tpVal - cofPayload.RealMargin.HPP - cofPayload.Marketing.Value)

	cof, existedCof, errIns := cofUse.cofRepos.Save(cofPayload)
	if errIns != nil {
		fmt.Println("cuse save errIns:", errIns)
		return nil, nil, errIns
	}

	if existedCof != nil {
		return nil, existedCof, nil
	}

	return cof, nil, nil
}

func (cofUse *coffeeUsecase) FindByID(id string) (*model.CoffeeEntity, error) {
	cof, err := cofUse.cofRepos.FindByID(id)
	if err != nil {
		return nil, err
	}

	return cof, nil
}

func (cofUse *coffeeUsecase) FindAll() ([]*model.CoffeeEntity, error) {
	cof, err := cofUse.cofRepos.FindAll()
	if err != nil {
		return nil, err
	}

	return cof, nil
}

func (cofUse *coffeeUsecase) Update(id string, cofPayload *model.CoffeeEntity) (*model.CoffeeEntity, error) {
	cof, err := cofUse.cofRepos.Update(id, cofPayload)
	if err != nil {
		return nil, err
	}

	return cof, nil
}

func (cofUse *coffeeUsecase) Delete(id string) (bool, error) {
	_, err := cofUse.cofRepos.Delete(id)
	if err != nil {
		return false, err
	}

	return true, nil
}
