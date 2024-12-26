func (forest *TokenForest) Find(tag string, attrs map[string]string) *TokenNode { // 1
	var result *TokenNode
	for _, tree := range forest.Trees { // 2
		result = tree.Find(tag, attrs) // 3
		if result != nil {             // 4
			return result // 5
		}
	}
	return nil
}

func (root *TokenNode) Find(tag string, attrs map[string]string) *TokenNode { // 6
	if root.StartTag.Data == tag { // 7
		if attrs == nil { // 8
			return root // 9
		}
		for _, attr := range root.StartTag.Attr { // 10
			if attrs[attr.Key] == attr.Val { // 11
				return root // 12
			}
		}
	}
	for _, child := range root.Children { // 13
		result := child.Find(tag, attrs) // 14
		if result != nil {               // 15
			return result // 16
		}
	}
	return nil // 17
}
func (node *TokenNode) CheckNode(tag string, attrs map[string]string) bool {

	return false
}