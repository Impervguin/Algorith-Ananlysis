package parsing

import (
	"fmt"
	"strings"

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

func newTokenTree(tokenizer *html.Tokenizer, initToken *html.Token) (*TokenNode, error) {
	var startToken html.Token
	if initToken == nil {
		startToken = tokenizer.Token()
	} else {
		startToken = *initToken
	}

	now := &TokenNode{
		StartTag: &startToken,
	}

	nextType := tokenizer.Next()
	if startToken.Type != html.StartTagToken || nextType == html.ErrorToken {
		return now, nil
	}
	now.Children = make([]*TokenNode, 0, 1)

	var token *html.Token
	for nextType != html.EndTagToken && nextType != html.ErrorToken {
		child, err := newTokenTree(tokenizer, token)
		if err != nil {
			return nil, err
		}
		now.Children = append(now.Children, child)
		tk := tokenizer.Token()
		token = &tk
		nextType = tk.Type
	}
	if nextType == html.EndTagToken {
		now.EndTag = token
		tokenizer.Next()
		return now, nil
	}
	return now, nil
}

func (root *TokenNode) Print(indent int) {
	for i := 0; i < indent; i++ {
		fmt.Printf(" ")
	}
	fmt.Println(strings.TrimSpace((*root.StartTag).String()))
	if root.Children != nil {
		for _, child := range root.Children {
			child.Print(indent + 4)
		}
	}
	if root.EndTag != nil {
		for i := 0; i < indent; i++ {
			fmt.Printf(" ")
		}
		fmt.Println(strings.TrimSpace((*root.EndTag).String()))
	}
}

func (node *TokenNode) CheckNode(tag string, attrs map[string]string) bool {
	if node.StartTag.Data == tag {
		if attrs == nil {
			return true
		}
		for _, attr := range node.StartTag.Attr {
			if attrs[attr.Key] == attr.Val {
				return true
			}
		}
	}
	return false
}

func (root *TokenNode) Find(tag string, attrs map[string]string) *TokenNode {
	if root.CheckNode(tag, attrs) {
		return root
	}
	for _, child := range root.Children {
		result := child.Find(tag, attrs)
		if result != nil {
			return result
		}
	}
	return nil
}

func (root *TokenNode) FindAll(tag string, attrs map[string]string) []*TokenNode {
	var result []*TokenNode
	if root.CheckNode(tag, attrs) {
		result = append(result, root)
	}
	for _, child := range root.Children {
		result = append(result, child.FindAll(tag, attrs)...)
	}
	return result
}
