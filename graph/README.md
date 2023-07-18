# Graph

A graph is a collection of vertices connected through directed or undirected edges. In an edge-weighted tree, each edge is assigned a value representing its cost or benefit. Graphs can be used to model various real-world scenarios. For example, graph A in the figure below can represent university courses and their prerequisites, cities and their highways, social network followers, site links, and many more.

```ASCII
[Figure 1] Graph Examples - Numbers, arrows, and numbers in brackets represent vertices, edges, and edge weights.

     5		3     5	  3              6   	       1 - Is it a DAG?
  ↗  ↑  ↖     ↗   ↘    ↖   ↘	      (4)  ↖ (1)	 2- Is it a connected graph?
 2 → 3 ← 4   2     4    4 ← 2      2 ─────→ 4        A 1 1
  ↖  ↑  ↗     ↖   ↙        ↗   (1) ↑ (2)   ↗ (1)     B 0 1
     1		1         1        1 ──→ 5	     C 1 0
						     D 1 1
    (A)	       (B)	 (C)  	      (D)
```

Graph theory a branch of mathematics is dedicated to studying graphs. Here are some essential graph concepts to remember:

* **Directed Acyclic Graph - DAG**: A directed graph with no cycles
* **Connected Graph**: There is a path between any two vertices
* **Minimum Spanning Tree**: Subset of edges in an undirected, edge-weighted, and connected graph that connects all the vertices at the lowest cost

## Implementation

Graphs are commonly represented using an adjacency matrix or an adjacency list.

* **Adjacency Matrix**: Faster lookup times and more suitable for dense graphs
* **Adjacency List**: More space efficient, suitable for graphs with fewer edges

```Go
package main

import "fmt"

func main() {
	adjacencyMatrix := [][]int{
		//    1  2  3  4  5
		[]int{0, 1, 1, 1, 0}, // 1
		[]int{0, 0, 1, 0, 1}, // 2
		[]int{0, 0, 0, 0, 1}, // 3
		[]int{0, 0, 1, 0, 1}, // 4
		[]int{0, 0, 0, 0, 0}, // 5
	}

	adjacencyList := [][]int{
		[]int{2, 3, 4},
		[]int{3, 5},
		[]int{5},
		[]int{3, 5},
		[]int{},
	}

	fmt.Println(adjacencyMatrix, adjacencyList)
}
```

One approach to make the graph more manageable is to store the adjacency list in each node and represent the graph as an array or slice of vertices.

```Go
package main

import "fmt"

// Vertex is a vertex in a Graph that has a value and can be connected to more vertices
type Vertex struct {
	// Val is the value of the vertex
	Val int

	// The edges that this vertex is connected to
	Edges []*Vertex
}

var graph []*Vertex

func main() {
	vertex := &Vertex{
		Val: 1,
	}
	graph = append(graph, vertex)
	fmt.Println(graph[0].Val) // Prints 1
}
```

### Searching Graphs

When working with graphs, searching to solve problems is often necessary. The algorithm's effectiveness depends on the order in which the graph is searched. Two commonly used search methods are:

* **Breadth First Search - BFS** used to find the shortest path
* **Depth First Search - DFS** often a subroutine, in another algorithm. Used in maze traversal, cycle finding, and pathfinding

Both BFS and DFS can be [implemented iteratively using a container (queue or stack)](./iterative_traversal_test.go).

#### Breadth First Search - BFS

BFS is a non-recursive algorithm that employs a [queue](../queue) and is a generalization of [post-order traversal](../tree) in a tree. When provided with a graph G and a vertex V, BFS systematically explores all nodes in G that are reachable from V, as illustrated in the following example:

```Go
package main

import (
	"container/list"
	"fmt"
	"log"
)

type vertexWithDistance struct {
	val   int
	edges []*vertexWithDistance

	distance int
}

func bfsRecursive(source *vertexWithDistance, distance int) {
	source.distance = distance
	for _, edge := range source.edges {
		bfsRecursive(edge, distance+1)
	}
}

func bfsIterative(source *vertexWithDistance) {
	distance := 0
	seen := make(map[*vertexWithDistance]struct{})
	queue := list.New()
	queue.PushBack(source)

	for queue.Len() != 0 {
		tmp := queue.Remove(queue.Front()).(*vertexWithDistance)
		distance++
		for _, v := range tmp.edges {
			if _, ok := seen[v]; ok {
				continue
			}
			seen[v] = struct{}{}
			v.distance = distance
			queue.PushBack(v)
		}
	}
}

func main() {
	graph1 := makeGraph()
	bfsIterative(graph1[0])

	graph2 := makeGraph()
	bfsRecursive(graph2[0], 0)

	for i, vertex := range graph1 {
		if vertex.distance == graph2[i].distance {
			fmt.Printf("vertex val: %d, distance: %d\n", vertex.val, vertex.distance)
		} else {
			log.Fatal("mismatch between iterative and recursive")
		}
	}
}

func makeGraph() []*vertexWithDistance {
	graph := []*vertexWithDistance{
		{val: 1},
		{val: 2},
		{val: 3},
		{val: 4},
	}
	graph[0].edges = []*vertexWithDistance{graph[1]}
	graph[1].edges = []*vertexWithDistance{graph[2]}
	graph[2].edges = []*vertexWithDistance{graph[3]}
	return graph
}
```

* Vertices at distance K are discovered before those at distance K + 1
* The shortest distance (i.e., the smallest number of edges) from the source vertex to each reachable vertex is computed
* A BFS tree with root V containing all reachable vertices is produced

For any vertex S that is reachable from V, the simple path in the BFS tree from V to S corresponds to the _shortest_ path (minimum number of edges) from V to S in G.

#### Depth First Search - DFS

Depth First Search (DFS) is a graph traversal algorithm that explores a graph as far as possible along each branch before backtracking. When implemented iteratively, it uses a [stack](../stack) data structure, is [recursive](../recursion), and is a generalization of pre-order traversal in trees.

When given a graph G and a vertex S, DFS systematically discovers all nodes in G reachable from S. It is typically implemented using a driver that discovers the edges of the most recently discovered vertex V with unexplored edges. Once all of V's edges have been explored, the search [backtracks](../backtracking/) to explore all edges leaving the vertex from which V was discovered. This process continues until all the edges are discovered.

```Go
package main

import "fmt"

type timedVertex struct {
	val   int
	edges []*timedVertex

	discoveryTimeStart  int
	discoveryTimeFinish int
}

var (
	seen = make(map[*timedVertex]struct{})
	time = 0
)

func dfs(graph []*timedVertex) {
	for _, v := range graph {
		if _, ok := seen[v]; !ok {
			dfsRecursive(v)
		}
	}
}

func dfsRecursive(u *timedVertex) {
	time++
	u.discoveryTimeStart = time
	for _, v := range u.edges {
		if _, ok := seen[v]; !ok {
			dfsRecursive(v)
		}
	}
	seen[u] = struct{}{}
	time++
	u.discoveryTimeFinish = time
}

func main() {
	graph := makeGraph()
	dfs(graph)
	for _, vertex := range graph {
		fmt.Printf("vertex val: %d, discoveryTimeStart: %d, discoveryTimeFinish: %d\n", vertex.val, vertex.discoveryTimeStart, vertex.discoveryTimeFinish)
	}
}

func makeGraph() []*timedVertex {
	graph := []*timedVertex{
		{val: 1},
		{val: 2},
		{val: 3},
		{val: 4},
	}
	graph[0].edges = []*timedVertex{graph[1]}
	graph[1].edges = []*timedVertex{graph[2]}
	graph[2].edges = []*timedVertex{graph[3]}
	return graph
}
```

* Sets the status of edges at the time of discovery and at the time exploring is finished
* Produces a depth-first forest comprising several depth-first trees
* Generates a well-formed parentheses structure, where the discovery of a vertex V is represented with (V and finishing discovering it is represented with V). For instance, applying DFS on graph (A) in _Figure 1_ would result in the output `(1(2(3(5))4()))`.

DFS is capable of categorizing edges (u,v) into four types:

* Tree Edge: v is discovered by exploring edge (u,v), all green edges
* Back edge: v is an ancestor of u, 6 to 2.
* Forward edge: v is a decedent of u, 1 to 8
* Cross edge: all other edges, 5 to 4

Discovery of a back edge in a DFS algorithm indicates a cyclic graph.

### Dijkstra's Algorithm

A [greedy](../greedy) algorithm that uses BFS-like ideas and a minimum [heap](../heap) to solve single-source shortest path problems in edge-weighted directed graphs like (D) in _Figure 1_.

 ```Go
package graph_test

import "container/heap"

type (
    dijkstraVertex struct {
        val int
        distance int
        discovered bool
        edges map[]*dijkstraVertex
    }
    verticesHeap []*dijkstraVertex
)

func dijkstra(graph []*timedVertex) {
    h := new(pointsHeap)
	for _, vertex := range graph {
		heap.Push(h, vertex)
    }

    s := make(map[*Vertex]struct{})
    for h.Len() != 0 {
        u := h.Pop().(*Vertex)
        u.discovered = true

        for _, v := range u.Edges {
            if v.discovered {
                continue
            }

            if v.distance + cvw < w.distance {
                w.distance
            }
        }
    }
}

func (p verticesHeap) Len() int            { return len(p) }
func (p verticesHeap) Less(i, j int) bool  { return p[i].distance < p[j].distance }
func (p verticesHeap) Swap(i, j int)       { p[i], p[j] = p[j], p[i] }
func (p *verticesHeap) Push(x interface{}) { *p = append(*p, x.(*Vertex)) }

func (p *pointsHeap) Pop() interface{} {
	old := *p
	tmp := old[len(old)-1]
	*p = old[0 : len(old)-1]
	return tmp
}
```

## Complexity

BFS and DFS are both O(E + V), where E is the number of edges and V is the number of vertices.

Dijkstra's algorithm time complexity depends on the minimum heap implementation. Adding all items to the heap is O(V). Furthermore, the time complexity for each pop operation in a binary [heap](../heap/) is O(log V), and there are E such operations at worst. As a result, the total time complexity is O((V + E) log V), which is equivalent to O(E log V).

## Applications

Graph algorithms find extensive usage in addressing real-world challenges, including but not limited to map navigation, flight and airport management, computer networking, image processing, and social media connectivity.

## Rehearsal

* [Iteratively Implement BFS and DFS](./iterative_traversal_test.go), [Solution](./iterative_traversal.go)
* [Is Graph a DAG](./is_dag_test.go), [Solution](./is_dag.go)
* [Topological Sort](./topological_sort_test.go), [Solution](./topological_sort.go)
* [Employee Head Count](./employee_headcount_test.go), [Solution](./employee_headcount.go)
* [Remove Invalid Parentheses](./remove_invalid_parentheses_test.go), [Solution](./remove_invalid_parentheses.go)
* [Cheapest Flights](./cheapest_flights_test.go), [Solution](./cheapest_flights.go)
* [Word Ladder](./word_ladder_test.go), [Solution](./word_ladder.go)
* [Network Delay Time](./network_delay_time_test.go), [Solution](./network_delay_time.go)
* [Number of Islands](./number_of_islands_test.go), [Solution](./number_of_islands.go)
* [Dependency Order](./dependency_order_test), [Solution](./dependency_order_test.go)
