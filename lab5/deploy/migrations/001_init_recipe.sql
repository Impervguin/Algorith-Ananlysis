-- +goose Up

CREATE TABLE IF NOT EXISTS recipe (
    id int PRIMARY KEY,
    issue_id int NOT NULL,
    url text NOT NULL,
    title text NOT NULL,
    image_url text NOT NULL
);

CREATE TABLE IF NOT EXISTS ingredient (
    recipe_id INT REFERENCES recipe(id),
    name TEXT NOT NULL,
    amount TEXT NOT NULL,
    unit TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS step (
    recipe_id INT REFERENCES recipe(id),
    step_number INT NOT NULL,
    step text NOT NULL
);

ALTER TABLE step ADD CONSTRAINT unique_step_per_recipe UNIQUE (recipe_id, step_number);
-- +goose Down

DROP TABLE IF EXISTS step;
DROP TABLE IF EXISTS ingredient;
DROP TABLE IF EXISTS recipe;