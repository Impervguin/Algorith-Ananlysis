package parsing

import (
	"io"

	"golang.org/x/net/html"
)

func NewTokenForest(htmlpage io.Reader) (*TokenForest, error) {
	tokenizer := html.NewTokenizer(htmlpage)
	forest := &TokenForest{
		Trees: make([]*TokenNode, 0, 1),
	}
	nextType := tokenizer.Next()
	for nextType != html.ErrorToken {
		tree, err := newTokenTree(tokenizer, nil)
		if err != nil {
			return nil, err
		}
		forest.Trees = append(forest.Trees, tree)
		nextType = tokenizer.Token().Type
	}
	return forest, nil
}

func (forest *TokenForest) Print() {
	if forest.Trees != nil {
		for _, tree := range forest.Trees {
			tree.Print(0)
		}
	}
}

func (forest *TokenForest) Find(tag string, attrs map[string]string) *TokenNode {
	var result *TokenNode
	for _, tree := range forest.Trees {
		result = tree.Find(tag, attrs)
		if result != nil {
			return result
		}
	}
	return nil
}

func (forest *TokenForest) FindAll(tag string, attrs map[string]string) []*TokenNode {
	var result []*TokenNode
	for _, tree := range forest.Trees {
		result = append(result, tree.FindAll(tag, attrs)...)
	}
	return result
}
