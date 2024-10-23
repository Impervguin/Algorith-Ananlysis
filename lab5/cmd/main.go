package main

import (
	"fmt"
	"lab5/internal/recipedb"
	"lab5/internal/recipeparse"
	"net/http"
)

func main() {
	resp, _ := http.Get("https://menunedeli.ru/recipe/samyj-vkusnyj-grushevyj-pirog/")

	// b, _ := os.ReadFile("test.html")
	// tk := html.NewTokenizer(bytes.NewReader(b))
	// fmt.Println(tk.Token())
	// parsing.NewTokenForest(tk)
	// forest, _ := parsing.NewTokenForest(resp.Body)
	// forest.Print()
	recipe := &recipedb.Recipe{}
	_ = recipeparse.ParseRecipe(resp.Body, recipe)
	// fmt.Println(recipe)
	resp.Body.Close()
	fmt.Println(recipe)
	// fmt.Println(recipe.Title)
	// fmt.Println(recipe.ImageURL)
	// fmt.Println(recipe.Ingredients)
	// fmt.Println(recipe.Steps)
}
