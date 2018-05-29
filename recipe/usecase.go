package recipe

import "github.com/kopi-isi-api-v1/model/recipe"

type IngridientsUsecase interface {
	Save(ingPayload *recipe.IngridientMaster) (*recipe.IngridientMaster, map[string]string, error)
	FindByID(id string) (*recipe.IngridientMaster, error)
	FindAll() ([]*recipe.IngridientMaster, error)
	Update(id string, ingPayload *recipe.IngridientUpdate) (*recipe.IngridientUpdate, error)
	Delete(id string) (bool, error)
}
