package main

import (
	"sort"
	"strconv"
	"strings"
)

// Nodes struct
type Nodes struct {
	nodes  map[int]Node
	n      int // number of nodes
	orders []int
	cycles *[][]int
}

// NewNodes initialise a new node
func NewNodes(input []string) *Nodes {
	n, _ := strconv.Atoi(input[0])
	t := Nodes{map[int]Node{}, n, []int{}, &[][]int{}}
	// convert input into nodes digraph map
	for _, degree := range input[1:] {
		d := strings.Fields(degree)
		first, _ := strconv.Atoi(d[0])
		second, _ := strconv.Atoi(d[1])
		t.NodeInsert(first, second)
	}

	return &t
}

// NodeInsert initialise nodes map[int]Node
func (t Nodes) NodeInsert(out int, in int) {
	// initalise
	if _, okIn := t.nodes[in]; okIn {
		t.nodes[in].in[out] = true
	} else {
		t.nodes[in] = Node{in: map[int]bool{out: true}, out: map[int]bool{}}
	}

	if _, okOut := t.nodes[out]; okOut {
		t.nodes[out].out[in] = true
	} else {
		t.nodes[out] = Node{out: map[int]bool{in: true}, in: map[int]bool{}}
	}
}

// SameNode find all nodes with same in-degree and out-degree
func (t Nodes) SameNode() []int {
	same := []int{}
	for k, node := range t.nodes {
		if len(node.in) == len(node.out) {
			same = append(same, k)
		}
	}
	sort.Ints(same)
	return same
}

// AverageDegree find the average degree of all nodes
func (t Nodes) AverageDegree() (float32, float32) {
	in, out := 0, 0
	for _, node := range t.nodes {
		in = in + len(node.in)
		out = out + len(node.out)
	}
	return divideInt(in, t.n), divideInt(out, t.n)
}

//FindOrder give topological order sorted
func (t Nodes) FindOrder() ([]int, bool) {
	for len(t.nodes) != 0 {
		zero := t.zeroInDegree()
		if len(zero) == 0 {
			// return false cause not in topological ordered
			return []int{}, false
		}
		//sort layer in order
		sort.Ints(zero)
		for _, v := range zero {
			t.orders = append(t.orders, v)
			t.deleteZeroInNode(v)
			delete(t.nodes, v)
		}
	}
	return t.orders, true
}

// FindAllCycle give all cycles
func (t Nodes) FindAllCycle() ([][]int, string) {

	// Remove all nodes with an out degree of zero
	for len(t.nodes) != 0 {
		zero := t.zeroOutDegree()
		if len(zero) == 0 {
			break
		}
		for _, v := range zero {
			t.deleteZeroOutNode(v)
			delete(t.nodes, v)
		}
	}
	// recursive find cycle
	for node := range t.nodes {
		t.findCycleUnit(node, node, []int{})
		t.removeNode(node)
	}
	yes := "No"
	if len(*t.cycles) <= 3 {
		yes = "Yes"
	}
	return *t.cycles, yes
}

// findCycleUnit find loops
func (t Nodes) findCycleUnit(u int, v int, path []int) {
	for _, s := range path {
		if s == v {
			if s == u {
				*t.cycles = append(*t.cycles, path)
			}
			return
		}
	}
	path = append(path, v)
	for node := range t.nodes[v].out {
		// recursive call
		t.findCycleUnit(u, node, path)
	}
}

// zeroInDegree find all nodes that has zero In-Degree
func (t Nodes) zeroInDegree() []int {
	zeroInDegree := []int{}
	for k, v := range t.nodes {
		if len(v.in) == 0 {
			zeroInDegree = append(zeroInDegree, k)
		}
	}
	return zeroInDegree
}

// zeroOutDegree find all nodes that has zero In-Degree
func (t Nodes) zeroOutDegree() []int {
	zeroOutDegree := []int{}
	for k, v := range t.nodes {
		if len(v.out) == 0 {
			zeroOutDegree = append(zeroOutDegree, k)
		}
	}
	return zeroOutDegree
}

// removeNode remove node from nodes
func (t Nodes) removeNode(node int) {
	t.deleteZeroInNode(node)
	t.deleteZeroOutNode(node)
	delete(t.nodes, node)
}

// deleteZeroInNode delete a node with zero In-Degree
func (t Nodes) deleteZeroInNode(node int) {
	for in := range t.nodes[node].out {
		delete(t.nodes[in].in, node)
	}
}

// deleteZeroOutNode delete a node with zero Out-Degree
func (t Nodes) deleteZeroOutNode(node int) {
	for out := range t.nodes[node].in {
		delete(t.nodes[out].out, node)
	}
}
