package simpletree

// Tree Simple tree type
type Tree struct {
	Text string
	N0 *Tree
	N1 *Tree
}

// NewT returns root
func NewT() *Tree {
	return &Tree{Text: "Root"}
}

// AddN0 adds N0 node
func (tree *Tree)AddN0(text string) {
	n0 := NewT()
	n0.Text = text
	tree.N0 = n0
}