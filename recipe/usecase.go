package recipe

import "github.com/kopi-isi-api-v1/model/recipe"

type IngredientsUsecase interface {
	Save(ingPayload *recipe.IngredientMaster) (*recipe.IngredientMaster, map[string]string, error)
	FindByID(id string) (*recipe.IngredientMaster, error)
	FindAll() ([]*recipe.IngredientMaster, error)
	Update(id string, ingPayload *recipe.IngredientUpdate) (*recipe.IngredientUpdate, error)
	Delete(id string) (bool, error)
}

type CoffeeUsecase interface {
	Save(cofPayload *recipe.CoffeeEntity) (*recipe.CoffeeEntity, map[string]string, error)
	Update(id string, ingPayload *recipe.CoffeeEntity) (*recipe.CoffeeEntity, error)
	FindByID(id string) (*recipe.CoffeeEntity, error)
	FindAll() ([]*recipe.CoffeeEntity, error)
	Delete(id string) (bool, error)
}
