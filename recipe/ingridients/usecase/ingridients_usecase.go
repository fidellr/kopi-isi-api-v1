package usecase

import (
	"fmt"

	errorModel "github.com/kopi-isi-api-v1/model"
	model "github.com/kopi-isi-api-v1/model/recipe"
	"github.com/kopi-isi-api-v1/recipe"
	"github.com/kopi-isi-api-v1/validator"
)

type ingridientUsecase struct {
	ingRepos recipe.Ingridients
}

func NewIngridientUsecase(ing recipe.Ingridients) recipe.Ingridients {
	return &ingridientUsecase{
		ingRepos: ing,
	}
}

func (ingUse *ingridientUsecase) Save(ingPayload *model.IngridientMaster) (*model.IngridientMaster, map[string]string, error) {
	if err := validator.Validate(ingPayload); err != nil {
		return nil, nil, recipe.NewErrorInvalidRecipeData(err.Error())
	}
	ing, existedIng, err := ingUse.ingRepos.Save(ingPayload)
	if err != nil {
		return nil, nil, err
	}
	if existedIng != nil {
		return nil, existedIng, nil
	}

	return ing, nil, nil
}

func (ingUse *ingridientUsecase) FindByID(id string) (*model.IngridientMaster, error) {
	ing, err := ingUse.ingRepos.FindByID(id)
	if err != nil {
		return nil, err
	}
	return ing, nil
}

func (ingUse *ingridientUsecase) FindAll() ([]*model.IngridientMaster, error) {
	listOfIng, err := ingUse.ingRepos.FindAll()
	if err != nil {
		fmt.Println("err nih")
		return nil, err
	}
	return listOfIng, nil
}

func (ingUse *ingridientUsecase) Update(id string, ingPayload *model.IngridientUpdate) (*model.IngridientUpdate, error) {
	ing, err := ingUse.ingRepos.Update(id, ingPayload)
	if err != nil {
		return nil, err
	}
	return ing, err
}

func (ingUse *ingridientUsecase) Delete(id string) (bool, error) {
	_, errFind := ingUse.FindByID(id)
	if errFind != nil {
		return false, errorModel.NOT_FOUND_ERROR
	}

	_, errDel := ingUse.ingRepos.Delete(id)
	if errDel != nil {
		return false, errDel
	}

	return true, nil
}
