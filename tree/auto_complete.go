package tree

const (
	alphabetSize = 26
)

type (
	trieNode struct {
		children  [alphabetSize]*trieNode
		isWordEnd bool
	}
	trie struct {
		root *trieNode
	}
)

// AutoComplete suggests auto-completion ideas based on the input and given dictionary.
func (t *trie) AutoComplete(word string) []string {
	output := []string{}
	current := t.root
	for i := 0; i < len(word); i++ {
		index := word[i] - 'a'
		if current.children[index] == nil {
			return output
		}
		current = current.children[index]
	}
	return current.readWords()
}

func newTrie(dict []string) *trie {
	trie := &trie{
		root: new(trieNode),
	}
	for _, word := range dict {
		trie.insert(word)
	}
	return trie
}

func (t *trie) insert(word string) {
	current := t.root
	for i := 0; i < len(word); i++ {
		index := word[i] - 'a'
		if current.children[index] == nil {
			current.children[index] = new(trieNode)
		}
		current = current.children[index]
	}
	current.isWordEnd = true
}

func (t *trieNode) readWords() []string {
	var (
		words        = []string{}
		bfsRecursive func(node *trieNode, parent string)
	)
	bfsRecursive = func(node *trieNode, parent string) {
		if node.isWordEnd && parent != "" {
			words = append(words, parent)
		}

		for i := 0; i < alphabetSize; i++ {
			if node.children[i] != nil {
				bfsRecursive(node.children[i], parent+string(rune(i+'a')))
			}
		}
	}

	bfsRecursive(t, "")
	return words
}
