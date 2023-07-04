package main

type leafNode struct {
	key   []byte
	value interface{}
}
type meta struct {
	prefix    []byte
	prefixLen int
	size      int
}
type innerNode struct {
	// meta attributes
	meta
	nodeType int
	keys     []byte
	children []*Node
}
type Node struct {
	// inner nodes map partial keys to child pointers
	innerNode *innerNode

	// leaf is used to store possible leaf
	leaf *leafNode
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func replace(old, new *Node) {
	*old = *new
}
func replaceRef(oldNode **Node, newNode *Node) {
	*oldNode = newNode
}

func copyBytes(dest []byte, src []byte, numBytes int) {
	for i := 0; i < numBytes && i < len(src) && i < len(dest); i++ {
		dest[i] = src[i]
	}
}
func main() {

}
