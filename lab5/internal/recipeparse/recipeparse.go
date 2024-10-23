package recipeparse

import (
	"io"
	"lab5/internal/parsing"
	"lab5/internal/recipedb"
)

// Retrieves a recipe from a parsed menunedeli.com token forest.
func ParseRecipe(reader io.ReadCloser, recipe *recipedb.Recipe) error {
	// Implement recipe parsing logic here
	forest, err := parsing.NewTokenForest(reader)
	if err != nil {
		return err
	}
	// forest.Print()
	// Extract recipe details from the parsed token forest
	title, err := ExtractTitle(forest)
	if err != nil {
		return err
	}
	imageUrl, err := ExtractImageURL(forest)
	if err != nil {
		return err
	}
	ingredients, err := ExtractIngredients(forest)
	if err != nil {
		return err
	}
	steps, err := ExtractSteps(forest)
	if err != nil {
		return err
	}
	recipe.Title = title
	recipe.ImageURL = imageUrl
	recipe.Ingredients = ingredients
	recipe.Steps = steps
	return nil
}

func ExtractTitle(forest *parsing.TokenForest) (string, error) {
	node := forest.Find("h1", map[string]string{"class": "hdr recipe-name", "itemprop": "name"})
	if node == nil {
		return "", ErrNoTitleFound
	}
	res, err := node.FindOneChildText()
	if err != nil {
		return "", err
	}
	return res, nil

}

func ExtractImageURL(forest *parsing.TokenForest) (string, error) {
	node := forest.Find("img", map[string]string{"class": "main-img", "itemprop": "image"})
	if node == nil {
		return "", ErrNoImageURLFound
	}
	url, ok := node.GetAttribute("src")
	if !ok {
		return "", ErrNoImageURLFound
	}
	return url, nil
}

func ExtractIngredients(forest *parsing.TokenForest) ([]recipedb.Ingredient, error) {
	ingredients := make([]recipedb.Ingredient, 0)
	node := forest.Find("ul", map[string]string{"class": "ingredients-lst"})
	if node == nil {
		return nil, ErrNoIngredientsFound
	}
	ingredientNodes := node.FindAll("li", nil)
	if len(ingredientNodes) == 0 {
		return nil, ErrNoIngredientsFound
	}
	for _, ingredientNode := range ingredientNodes {
		nameNode := ingredientNode.Find("span", map[string]string{"class": "name"})
		if nameNode == nil {
			continue
		}
		name, err := nameNode.FindOneChildText()
		if err != nil {
			return nil, err
		}
		quantityNode := ingredientNode.Find("span", map[string]string{"class": "value"})
		if quantityNode == nil {
			continue
		}
		quantity, err := quantityNode.FindOneChildText()
		if err != nil {
			quantity = ""
		}
		unitNode := ingredientNode.Find("span", map[string]string{"class": "type"})
		if unitNode == nil {
			continue
		}
		unit, err := unitNode.FindOneChildText()
		if err != nil {
			unit = ""
		}
		ingredients = append(ingredients, recipedb.Ingredient{
			Name:   name,
			Amount: quantity,
			Unit:   unit,
		})
	}
	return ingredients, nil
}

func ExtractSteps(forest *parsing.TokenForest) ([]string, error) {
	steps := make([]string, 0)
	node := forest.Find("ul", map[string]string{"class": "instructions-lst", "itemprop": "recipeInstructions"})
	if node == nil {
		return nil, ErrNoStepsFound
	}
	stepNodes := node.FindAll("div", map[string]string{"class": "desc"})
	if len(stepNodes) == 0 {
		return nil, ErrNoStepsFound
	}

	for _, stepNode := range stepNodes {
		s, err := stepNode.GetText()
		if err != nil {
			return nil, err
		}
		steps = append(steps, s)
	}
	return steps, nil
}
