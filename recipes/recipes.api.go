package recipes

import (
	context "context"
	"fmt"
	"time"

	grpc "google.golang.org/grpc"
)

// RecipeClient comment
type RecipeClient struct {
	conn *grpc.ClientConn
}

// NewClient init
func NewClient(port string) (RecipeClient, error) {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	con, err := grpc.Dial(fmt.Sprintf("recipes:%s", port), opts...)
	return RecipeClient{
		conn: con,
	}, err
}

// GetRecipe function
func (recipeClient *RecipeClient) GetRecipe(recipeID *Recipe_ID) *RecipeResponse {
	client := NewRecipesClient(recipeClient.conn)

	// Создаем запрос
	request := &GetRecipeRequest{
		RecipeId: recipeID,
	}

	// Отправляем запрос
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	response, err := client.GetRecipe(ctx, request)

	if err != nil {
		return &RecipeResponse{
			Error: &Error{
				ErrorCode: 1,
				ErrorStr:  err.Error(),
			},
		}
	}

	return response
}

// PostRecipe function
func (recipeClient *RecipeClient) PostRecipe(recipeInfo *Recipe_Info) *RecipeResponse {
	client := NewRecipesClient(recipeClient.conn)

	request := &PostRecipeRequest{
		Info: recipeInfo,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	response, err := client.PostRecipe(ctx, request)

	if err != nil {
		return &RecipeResponse{
			Error: &Error{
				ErrorCode: 1,
				ErrorStr:  err.Error(),
			},
		}
	}

	return response
}

// DeleteRecipe method
func (recipeClient *RecipeClient) DeleteRecipe(recipeID *Recipe_ID) *Error {
	client := NewRecipesClient(recipeClient.conn)

	// Создаем запрос
	request := &DeleteRecipeRequest{
		RecipeId: recipeID,
	}

	// Отправляем запрос
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	response, err := client.DeleteRecipe(ctx, request)

	if err != nil {
		return &Error{
			ErrorCode: 1,
			ErrorStr:  err.Error(),
		}
	}

	return response
}
