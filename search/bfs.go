package search

import "github.com/bquenin/algorithms/ds"

type BFS struct {
	visited   []bool
	backtrack []int
	source    int
}

func NewBFS(g *ds.Graph, source int) *BFS {
	bfs := &BFS{
		visited:   make([]bool, g.V),
		backtrack: make([]int, g.V),
		source:    source,
	}
	bfs.bfs(g)
	return bfs
}

func (bfs *BFS) bfs(g *ds.Graph) {
	// Nodes to visit
	q := ds.NewListQueue()

	// Mark the source as visited and enqueue it
	bfs.visited[bfs.source] = true
	q.Enqueue(bfs.source)

	for !q.Empty() {
		from := q.Dequeue().(int)        // Get vertex 'from'
		for _, to := range g.Adj(from) { // Visit adjacencies 'to'
			if !bfs.visited[to] {
				bfs.backtrack[to] = from // Store the path between 'from' and 'to' in order to backtrack
				bfs.visited[to] = true   // Mark vertex as visited
				q.Enqueue(to)            // Enqueue it
			}
		}
	}
}

func (bfs *BFS) hasPathTo(to int) bool {
	return bfs.visited[to]
}

func (bfs *BFS) pathTo(to int) *ds.Stack {
	path := ds.NewStack()
	if !bfs.hasPathTo(to) {
		return path
	}
	for v := to; v != bfs.source; v = bfs.backtrack[v] {
		path.Push(v)
	}
	path.Push(bfs.source)
	return path
}
