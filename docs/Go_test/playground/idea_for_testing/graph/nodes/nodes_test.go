package nodes

import (
	"sync"
	"testing"
	"fmt"
)

func TestNode_String(t *testing.T) {
	type fields struct {
		value string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
		{name: "Name test", fields: fields{value: "Bozo"}, want: "Bozo"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &Node{
				value: tt.fields.value,
			}
			if got := n.String(); got != tt.want {
				t.Errorf("Node.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestItemGraph_AddNode(t *testing.T) {
	type fields struct {
		nodes   []*Node
		edges   map[Node][]*Node
		RWMutex sync.RWMutex
	}
	type args struct {
		n *Node
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &ItemGraph{
				nodes:   tt.fields.nodes,
				edges:   tt.fields.edges,
				RWMutex: tt.fields.RWMutex,
			}
			g.AddNode(tt.args.n)
		})
	}
}

func TestItemGraph_AddEdge(t *testing.T) {
	type fields struct {
		nodes   []*Node
		edges   map[Node][]*Node
		RWMutex sync.RWMutex
	}
	type args struct {
		n1 *Node
		n2 *Node
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &ItemGraph{
				nodes:   tt.fields.nodes,
				edges:   tt.fields.edges,
				RWMutex: tt.fields.RWMutex,
			}
			g.AddEdge(tt.args.n1, tt.args.n2)
		})
	}
}

func fillGraph(g *ItemGraph) {

	nA := Node{"A"}
	nB := Node{"B"}
	nC := Node{"C"}
	nD := Node{"D"}
	nE := Node{"E"}
	nF := Node{"F"}
	nOutside := Node{"Outside"}

	_ = nOutside
	g.AddNode(&nA)
	g.AddNode(&nB)
	g.AddNode(&nC)
	g.AddNode(&nD)
	g.AddNode(&nE)
	g.AddNode(&nF)
	g.AddNode(&nOutside)

	g.AddEdge(&nA, &nB)
	g.AddEdge(&nA, &nOutside)
	g.AddEdge(&nA, &nC)
	g.AddEdge(&nB, &nE)
	//g.AddEdge(&nC, &nE)
	//g.AddEdge(&nE, &nF)
	//g.AddEdge(&nD, &nA)

}




	func Test_Thing(t *testing.T) {
		var g ItemGraph
		fillGraph(&g)
		//g.Search(g.nodes[0])
		//g.String()
		r := g.JustEdges(g.nodes[0])
		fmt.Printf("done: %v\n",r.String())
	
	}
