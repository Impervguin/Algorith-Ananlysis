package main

import (
	"context"
	"fmt"
	"lab5/internal/recipedb"
	"os"
)

const JsonDir = "json"

func main() {
	pgs, err := recipedb.NewPgsStorage(context.TODO())
	if err != nil {
		fmt.Printf("Can't initialize PGS storage: %v\n", err)
		return
	}

	recipes, err := pgs.GetAllRecipes()
	if err != nil {
		fmt.Printf("Can't get all recipes from PGS: %v\n", err)
		return
	}
	for _, recipe := range recipes {
		ingredients, err := pgs.GetAllIngredients(recipe.ID)
		if err != nil {
			fmt.Printf("Can't get ingredients for recipe %d: %v\n", recipe.ID, err)
			continue
		}
		steps, err := pgs.GetAllSteps(recipe.ID)
		if err != nil {
			fmt.Printf("Can't get steps for recipe %d: %v\n", recipe.ID, err)
			continue
		}
		recipe.Ingredients = ingredients
		recipe.Steps = steps
		jsonRecipe := recipe.ToJson()
		os.WriteFile(fmt.Sprintf("%s/%s.json", JsonDir, recipe.Title), jsonRecipe, 0644)
	}
	fmt.Println("All recipes saved to JSON files")
}
