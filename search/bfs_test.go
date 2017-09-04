package search

import (
	"testing"

	"github.com/bquenin/algorithms/ds"
	"github.com/stretchr/testify/assert"
)

func TestBFS(t *testing.T) {
	V := 4
	E := [][2]int{
		{0, 1},
		{0, 2},
	}
	source := 0
	length := []int{1, 2, 2, 0}

	g := ds.NewGraph(V)
	for _, e := range E {
		from, to := e[0], e[1]
		g.AddEdge(from, to)
	}
	bfs := NewBFS(g, source)

	for i := 0; i < V; i++ {
		assert.Equal(t, length[i], bfs.pathTo(i).Size())
	}
}
