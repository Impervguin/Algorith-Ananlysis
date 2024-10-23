package main

import (
	"lab5/internal/parsing"
	"net/http"
)

func main() {
	resp, _ := http.Get("https://menunedeli.ru/recipe/pirog-s-brusnikoj-iz-drozhzhevogo-testa-v-duxovke/")

	// b, _ := os.ReadFile("test.html")
	// tk := html.NewTokenizer(bytes.NewReader(b))
	// fmt.Println(tk.Token())
	// parsing.NewTokenForest(tk)
	forest, _ := parsing.NewTokenForest(resp.Body)
	forest.Print()
}
