package recipe

import "github.com/kopi-isi-api-v1/model/recipe"

type Ingredients interface {
	Save(ingPayload *recipe.IngredientMaster) (*recipe.IngredientMaster, map[string]string, error)
	FindByID(id string) (*recipe.IngredientMaster, error)
	FindAll() ([]*recipe.IngredientMaster, error)
	Update(id string, ingPayload *recipe.IngredientUpdate) (*recipe.IngredientUpdate, error)
	Delete(id string) (bool, error)
}

type Coffee interface {
	Save(cofPayload *recipe.CoffeeEntity) (*recipe.CoffeeEntity, map[string]string, error)
	FindByID(id string) (*recipe.CoffeeEntity, error)
	FindAll() ([]*recipe.CoffeeEntity, error)
	Update(id string, cofPayload *recipe.CoffeeEntity) (*recipe.CoffeeEntity, error)
	Delete(id string) (bool, error)
}
