package lab7

import (
	"golang.org/x/net/html"
)

type TokenNode struct {
	StartTag *html.Token
	Children []*TokenNode
	EndTag   *html.Token
}

type TokenForest struct {
	Trees []*TokenNode
}

func (forest *TokenForest) Find(tag string, attrs map[string]string) *TokenNode { // 1
	var result *TokenNode
	for _, tree := range forest.Trees { // 2
		result = tree.Find(tag, attrs) // 3
		if result != nil {             // 4
			return result // 5
		}
	}
	return nil // 6
}

func (root *TokenNode) Find(tag string, attrs map[string]string) *TokenNode { // 7
	if root.StartTag.Data == tag { // 8
		if attrs == nil { // 9
			return root // 10
		}
		for _, attr := range root.StartTag.Attr { // 11
			if attrs[attr.Key] == attr.Val { // 12
				return root // 13
			}
		}
	}
	for _, child := range root.Children { // 14
		result := child.Find(tag, attrs) // 15
		if result != nil {               // 16
			return result // 17
		}
	}
	return nil // 18
}
