package graph

type (
	// Vertex is a vertex in a Graph that has a value and can be connected to more vertices.
	Vertex struct {
		// Val is the value of the vertex
		Val int

		// The edges that this Vertex is connected to
		Edges []*Vertex
	}
	queue     struct{ collection []*Vertex }
	stack     struct{ collection []*Vertex }
	container interface {
		Push(*Vertex)
		Pop() *Vertex
		Len() int
	}
)

// IterativeTraversal performs a BFS and DFS on the given graph and returns
// the value of each vertex in the order it was visited.
func IterativeTraversal(graph []*Vertex) ([]int, []int) {
	dfs := traverseAllNodes(new(stack), graph)
	bfs := traverseAllNodes(new(queue), graph)
	return bfs, dfs
}

// traverseAllNodes performs BFS or DFS if a graph and a container that is
// either respectively a stack or queue passed to it.
func traverseAllNodes(c container, graph []*Vertex) []int {
	output := []int{}
	visited := make(map[*Vertex]struct{})
	c.Push(graph[0])

	for c.Len() != 0 {
		tmp := c.Pop()
		for _, neighbour := range tmp.Edges {
			c.Push(neighbour)
		}
		if _, ok := visited[tmp]; !ok {
			output = append(output, tmp.Val)
			visited[tmp] = struct{}{}
		}
	}
	return output
}

func newVertex(v int) *Vertex { return &Vertex{Val: v, Edges: []*Vertex{}} }

func (q *queue) Len() int       { return len(q.collection) }
func (q *queue) Push(n *Vertex) { q.collection = append(q.collection, n) }
func (q *queue) Pop() *Vertex {
	tmp := q.collection[0]
	q.collection = q.collection[1:len(q.collection)]
	return tmp
}

func (s *stack) Len() int       { return len(s.collection) }
func (s *stack) Push(n *Vertex) { s.collection = append(s.collection, n) }
func (s *stack) Pop() *Vertex {
	tmp := s.collection[len(s.collection)-1]
	s.collection = s.collection[:len(s.collection)-1]
	return tmp
}
