package usecase

import (
	"fmt"

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
