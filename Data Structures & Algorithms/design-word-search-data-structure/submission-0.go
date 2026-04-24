type WordDictionary struct {
    root *Node
}

type Node struct {
	Val byte
	Neighbors [26]*Node
	IsEnd bool
}

func Constructor() WordDictionary {
    return WordDictionary{root: &Node{Val: '_', Neighbors: [26]*Node{}}}
}

func (this *WordDictionary) AddWord(word string)  {
    root := this.root

	for i:=0;i<len(word);i++{
		pos := int(word[i] - 'a')
		if pos > 25 {
			fmt.Println("unexpected")
			return
		}
		if root.Neighbors[pos] == nil {
			root.Neighbors[pos] = &Node{Val: word[i], Neighbors: [26]*Node{}}
		}
		root = root.Neighbors[pos]
	}
	root.IsEnd = true
}

func (this *WordDictionary) Search(word string) bool {
    root := this.root
	return this.SearchFromChar(word, root, 0)
}

func (this *WordDictionary) SearchFromChar(word string, node *Node, i int) bool {
	if node == nil {return false}
	if i == len(word) {return node.IsEnd}
	if word[i] == '.' {
		for j:=0;j<26;j++ {
			if node.Neighbors[j] != nil && this.SearchFromChar(word, node.Neighbors[j], i+1) {return true}
		}
		return false
	} else {
		pos := int(word[i] - 'a')
		return this.SearchFromChar(word, node.Neighbors[pos], i+1)
	}
}
