package recipedb

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

const putRecipe = `
INSERT INTO recipes (id, issue_id, url, title, image_url)
VALUES ($1, $2, $3, $4, $5);
`

const insertIngredient = `
INSERT INTO ingredients (recipe_id, name, amount, unit)
VALUES ($1, $2, $3, $4);
`

const insertStep = `
INSERT INTO steps (recipe_id, step_number, step)
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
