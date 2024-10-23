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

// Фронтендеры не закрывают теги на сайти, поэтому эти теги считаются блоками.
var SelfClosingTokens = map[string]struct{}{
	"area":   {},
	"base":   {},
	"br":     {},
	"col":    {},
	"embed":  {},
	"hr":     {},
	"img":    {},
	"input":  {},
	"link":   {},
	"meta":   {},
	"param":  {},
	"source": {},
	"track":  {},
	"wbr":    {},
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
	if _, ok := SelfClosingTokens[startToken.Data]; ok {
		now.StartTag.Type = html.SelfClosingTagToken
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
	result := make([]*TokenNode, 0, 10)
	if root.CheckNode(tag, attrs) {
		result = append(result, root)
	}
	for _, child := range root.Children {
		result = append(result, child.FindAll(tag, attrs)...)
	}
	return result
}

func (root *TokenNode) FindOneChildText() (string, error) {
	if root.Children == nil || len(root.Children) != 1 {
		return "", fmt.Errorf("no child nodes found")
	}
	child := root.Children[0]
	if child.StartTag.Type != html.TextToken {
		return "", fmt.Errorf("no text node found")
	}
	return child.StartTag.String(), nil
}

func (root *TokenNode) FindOneChildNode() *TokenNode {
	if root.Children == nil || len(root.Children) != 1 {
		return nil
	}
	return root.Children[0]
}

func (root *TokenNode) GetAttribute(attrName string) (string, bool) {
	for _, attr := range root.StartTag.Attr {
		if attr.Key == attrName {
			return attr.Val, true
		}
	}
	return "", false
}

var BannedTextTokens = map[string]struct{}{"noscript": {}, "script": {}, "style": {}}

func (root *TokenNode) GetText() (string, error) {
	var text strings.Builder
	for _, child := range root.Children {
		if child.StartTag.Type == html.TextToken {
			text.WriteString(child.StartTag.String())
		} else if child.StartTag.Type == html.StartTagToken {
			if _, ok := BannedTextTokens[child.StartTag.Data]; ok {
				continue
			}
			s, err := child.GetText()
			if err != nil {
				return "", err
			}
			text.WriteString(s)
		}
	}
	return text.String(), nil
}
