package recipes

import (
	context "context"

	grpc "google.golang.org/grpc"
)

// Client comment
type Client struct {
	conn *grpc.ClientConn
}

// NewClient init
func NewClient() (Client, error) {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	con, err := grpc.Dial("recipes:5301", opts...)
	return Client{con}, err
}

// IngredientJSON struct ...
type IngredientJSON struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// RecipeInfoJSON struct ...
type RecipeInfoJSON struct {
	Title       string           `json:"title"`
	Description string           `json:"description"`
	Ingredients []IngredientJSON `json:"ingredients"`
}

// RecipeJSON struct
type RecipeJSON struct {
	ID         int64          `json:"id"`
	recipeInfo RecipeInfoJSON `json:"recipe_info"`
	Error      string         `json:"error"`
}

// GetRecipe function
func (client *Client) GetRecipe(id int64) RecipeJSON {
	cli := NewRecipesClient(client.conn)
	request := GetRecipeRequest{
		Id: id,
	}
	response, err := cli.GetRecipe(context.Background(), &request)

	if err != nil {
		return RecipeJSON{
			Error: err.Error(),
		}
	}

	ingredients := []IngredientJSON{}
	for _, element := range response.Recipe.RecipeInfo.Ingredients {
		ingredients = append(ingredients, IngredientJSON{
			Name:  element.Name,
			Value: element.Value,
		})
	}

	return RecipeJSON{
		ID: response.Recipe.Id,
		recipeInfo: RecipeInfoJSON{
			Title:       response.Recipe.RecipeInfo.Title,
			Description: response.Recipe.RecipeInfo.Description,
			Ingredients: ingredients,
		},
	}
}

// PostRecipe function
func (client *Client) PostRecipe(recipeInfo RecipeInfoJSON) RecipeJSON {
	cli := NewRecipesClient(client.conn)

	ingredients := []*Ingredient{}
	for _, element := range recipeInfo.Ingredients {
		ingredients = append(ingredients, &Ingredient{
			Name:  element.Name,
			Value: element.Value,
		})
	}

	request := PostRecipeRequest{
		RecipeInfo: &RecipeInfo{
			Title:       recipeInfo.Title,
			Description: recipeInfo.Description,
			Ingredients: ingredients,
		},
	}

	response, err := cli.PostRecipe(context.Background(), &request)

	if err != nil {
		return RecipeJSON{
			Error: err.Error(),
		}
	}

	return RecipeJSON{
		ID:         response.Id,
		recipeInfo: recipeInfo,
	}
}
