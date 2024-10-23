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

func (p *PgsStorage) PutRecipe(recipe *Recipe) error {
	_, err := p.conn.Exec(putRecipe, recipe.ID, recipe.IssueID, recipe.Url, recipe.Title, recipe.ImageURL)
	return err
}

const insertIngredient = `
INSERT INTO ingredients (recipe_id, name, amount, unit)
VALUES ($1, $2, $3, $4);
`

func (p *PgsStorage) InsertIngredient(recipeID int64, ingredient *Ingredient) error {
	_, err := p.conn.Exec(insertIngredient, recipeID, ingredient.Name, ingredient.Amount, ingredient.Unit)
	return err
}

const insertStep = `
INSERT INTO steps (recipe_id, step_number, step)
VALUES ($1, $2, $3);
`

func (p *PgsStorage) InsertStep(recipeID int64, stepNumber int, step string) error {
	_, err := p.conn.Exec(insertStep, recipeID, stepNumber, step)
	return err
}
