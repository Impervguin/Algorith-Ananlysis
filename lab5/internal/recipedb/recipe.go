package recipedb

import (
	"encoding/json"

	"github.com/jackc/pgx"
)

type Ingredient struct {
	Name   string
	Amount string
	Unit   string
}

type Recipe struct {
	ID          int64
	IssueID     int64
	Url         string
	Title       string
	ImageURL    string
	Ingredients []Ingredient
	Steps       []string
}

func (r *Recipe) ToJson() []byte {
	json, err := json.Marshal(r)
	if err != nil {
		return nil
	}
	return json
}

const putRecipe = `
INSERT INTO recipe (id, issue_id, url, title, image_url)
VALUES ($1, $2, $3, $4, $5);
`

const insertIngredient = `
INSERT INTO ingredient (recipe_id, name, amount, unit)
VALUES ($1, $2, $3, $4);
`

const insertStep = `
INSERT INTO step (recipe_id, step_number, step)
VALUES ($1, $2, $3);
`

func (p *PgsStorage) PutRecipe(recipe *Recipe) error {
	tx, err := p.conn.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec(putRecipe, recipe.ID, recipe.IssueID, recipe.Url, recipe.Title, recipe.ImageURL)
	if err != nil {
		tx.Rollback()
		return err
	}

	for _, ingredient := range recipe.Ingredients {
		_, err := tx.Exec(insertIngredient, recipe.ID, ingredient.Name, ingredient.Amount, ingredient.Unit)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	for i, step := range recipe.Steps {
		_, err := tx.Exec(insertStep, recipe.ID, i+1, step)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	err = tx.Commit()
	return err
}

const getRecipe = `
SELECT id, issue_id, url, title, image_url
FROM recipe
WHERE title = $1;
`

func (p *PgsStorage) GetRecipeByTitle(title string) (*Recipe, error) {
	row := p.conn.QueryRow(getRecipe, title)
	recipe := &Recipe{}
	err := row.Scan(&recipe.ID, &recipe.IssueID, &recipe.Url, &recipe.Title, &recipe.ImageURL)
	if err == pgx.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return recipe, nil
}

const getAllRecipes = `
SELECT id, issue_id, url, title, image_url
FROM recipe
ORDER BY title;
`

func (p *PgsStorage) GetAllRecipes() ([]Recipe, error) {
	rows, err := p.conn.Query(getAllRecipes)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	recipes := make([]Recipe, 0, 1)
	for rows.Next() {
		recipe := Recipe{}
		err := rows.Scan(&recipe.ID, &recipe.IssueID, &recipe.Url, &recipe.Title, &recipe.ImageURL)
		if err != nil {
			return nil, err
		}
		recipes = append(recipes, recipe)
	}
	return recipes, nil
}

const getAllIngredients = `
SELECT name, amount, unit
FROM ingredient
WHERE recipe_id = $1;
`

func (p *PgsStorage) GetAllIngredients(recipeID int64) ([]Ingredient, error) {
	rows, err := p.conn.Query(getAllIngredients, recipeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ingredients := make([]Ingredient, 0, 1)
	for rows.Next() {
		ingredient := Ingredient{}
		err := rows.Scan(&ingredient.Name, &ingredient.Amount, &ingredient.Unit)
		if err != nil {
			return nil, err
		}
		ingredients = append(ingredients, ingredient)
	}
	return ingredients, nil
}

const getAllSteps = `
SELECT step
FROM step
WHERE recipe_id = $1
ORDER BY step_number;
`

func (p *PgsStorage) GetAllSteps(recipeID int64) ([]string, error) {
	rows, err := p.conn.Query(getAllSteps, recipeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	steps := make([]string, 0, 1)
	for rows.Next() {
		var step string
		err := rows.Scan(&step)
		if err != nil {
			return nil, err
		}
		steps = append(steps, step)
	}
	return steps, nil
}
