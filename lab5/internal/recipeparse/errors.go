package recipeparse

import "fmt"

var ErrNoTitleFound = fmt.Errorf("no title found in the recipe")
var ErrNoImageURLFound = fmt.Errorf("no image URL found in the recipe")
var ErrNoIngredientsFound = fmt.Errorf("no ingredients found in the recipe")
var ErrNoStepsFound = fmt.Errorf("no steps found in the recipe")
