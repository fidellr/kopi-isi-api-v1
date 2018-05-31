package usecase

import (
	model "github.com/kopi-isi-api-v1/model/recipe"
	"github.com/kopi-isi-api-v1/recipe"
	"github.com/kopi-isi-api-v1/validator"
)

type ingredientUsecase struct {
	ingRepos recipe.Ingredients
}

func NewIngredientUsecase(ing recipe.Ingredients) recipe.Ingredients {
	return &ingredientUsecase{
		ingRepos: ing,
	}
}

func (ingUse *ingredientUsecase) Save(ingPayload *model.IngredientMaster) (*model.IngredientMaster, map[string]string, error) {
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

func (ingUse *ingredientUsecase) FindByID(id string) (*model.IngredientMaster, error) {
	ing, err := ingUse.ingRepos.FindByID(id)
	if err != nil {
		return nil, err
	}
	return ing, nil
}

func (ingUse *ingredientUsecase) FindAll() ([]*model.IngredientMaster, error) {
	listOfIng, err := ingUse.ingRepos.FindAll()
	if err != nil {
		return nil, err
	}
	return listOfIng, nil
}

func (ingUse *ingredientUsecase) Update(id string, ingPayload *model.IngredientUpdate) (*model.IngredientUpdate, error) {

	ing, err := ingUse.ingRepos.Update(id, ingPayload)
	if err != nil {
		return nil, err
	}
	return ing, nil
}

func (ingUse *ingredientUsecase) Delete(id string) (bool, error) {
	_, errDel := ingUse.ingRepos.Delete(id)
	if errDel != nil {
		return false, errDel
	}

	return true, nil
}
