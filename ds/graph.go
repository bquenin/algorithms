package ds

type Graph struct {
	V int            // Number of vertices
	a []map[int]bool // Adjacencies
}

func NewGraph(V int) *Graph {
	g := &Graph{
		V: V,
		a: make([]map[int]bool, V),
	}
	for i := 0; i < V; i++ {
		g.a[i] = map[int]bool{}
	}
	return g
}

func (g *Graph) AddEdge(from, to int) {
	g.a[from][to] = true
	g.a[to][from] = true
}

func (g *Graph) Adj(vertex int) []int {
	adjacencies := make([]int, 0, len(g.a[vertex]))
	for k := range g.a[vertex] {
		adjacencies = append(adjacencies, k)
	}
	return adjacencies
}
