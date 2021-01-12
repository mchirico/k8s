package nodes

import (
	"fmt"
	"sync"
)

type Node struct {
	value string
}

func (n *Node) String() string {
	return fmt.Sprintf("%v", n.value)
}

// ItemGraph the Items graph
type ItemGraph struct {
	nodes []*Node
	edges map[Node][]*Node
	sync.RWMutex
}

// AddNode adds a node to the graph
func (g *ItemGraph) AddNode(n *Node) {
	g.Lock()
	defer g.Unlock()
	g.nodes = append(g.nodes, n)

}

// AddEdge adds an edge to the graph
func (g *ItemGraph) AddEdge(n1, n2 *Node) {
	g.Lock()
	defer g.Unlock()

	if g.edges == nil {
		g.edges = make(map[Node][]*Node)
	}
	g.edges[*n1] = append(g.edges[*n1], n2)
	//g.edges[*n2] = append(g.edges[*n2], n1)

}

func (g *ItemGraph) String() {
	g.RLock()
	g.RUnlock()
	s := ""
	for _, v := range g.nodes {
		s += v.String() + " -> "
		for _, v := range g.edges[*v] {
			s += v.String() + " "
		}
		s += "\n"
	}
	fmt.Println(s)

}

func (g *ItemGraph) Search(n *Node) *Node {
	g.RLock()
	g.RUnlock()

    if n == nil {
        return nil
	}

	
      



	s := ""

	fmt.Printf("%s\n", n.String())
	s += n.String() + " -> "

	for _, v := range g.edges[*n] {
		next := g.Search(v)
		if next != nil {
			s += next.String() + " "
		}
		
	}
	s += "\n"
	fmt.Printf("\n")
	fmt.Println(s)
	return n

}


func (g *ItemGraph) JustEdges(n *Node) *Node {
	g.RLock()
	g.RUnlock()

    if n == nil {
        return nil
	}
	s := ""

	fmt.Printf("%s\n", n.String())
	s += n.String() + " -> "

	var next *Node
	for _, v := range g.edges[*n] {
			s += v.String() + " "
			next = v
	}
	s += "\n"
	fmt.Printf("\n")
	fmt.Println(s)
	return next

}

