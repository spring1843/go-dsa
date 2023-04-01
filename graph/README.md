# Graph

A graph is a collection of vertices that are connected through directed or undirected edges. In an edge-weighted tree, each edge is assigned a value representing its cost or benefit. Graphs can be used to model various real-world scenarios, For example graph A in the below figure can represent university courses and their prerequisites, cities and their highways, followers in a social network, links on a website, and many more.


```ASCII
[Figure 1] Graph Examples - Numbers , arrows, and numbers in brackets each respectively represent vertices, edges, and edge weights.

     5		3     5	  3              6   	       1 - Is it a DAG?
  ↗  ↑  ↖     ↗   ↘    ↖   ↘	      (4)  ↖ (1)	 2- Is it a connected graph?
 2 → 3 ← 4   2     4    4 ← 2      2 ─────→ 4        A 1 1
  ↖  ↑  ↗     ↖   ↙        ↗   (1) ↑ (2)   ↗ (1)     B 0 1
     1		1         1        1 ──→ 5	     C 1 0
						     D 1 1
    (A)	       (B)	 (C)  	      (D)
```

An entire branch of mathematics named graph theory is dedicated to the study of graphs. Here are some essential graph concepts that are important to remember:

* **Directed Acyclic Graph - DAG**:  A directed graph with no cycles
* **Connected Graph**: There is a path between any two vertices
* **Minimum Spanning Tree**: Subset of edges in an undirected, edge weighted, and connected graph, that connects all the vertices at the lowest cost

## Implementation

Graphs are commonly represented using either an adjacency matrix or an adjacency list.

* **Adjacency Matrix**: Faster look up times and more suitable for dense graphs
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

When working with graphs, it is often necessary to perform a search to solve different problems. The efficiency of the algorithm used depends on the order in which the graph is searched. Two commonly used search methods are:

* **Breadth First Search - BFS** used to find the shortest path
* **Depth First Search - DFS** often a subroutine, in another algorithm e.g. in maze traversal, cycle finding, and path finding.

Both BFS and DFS can be [implemented iteratively using a container (queue, or stack)](#iteratively-implement-bfs-and-dfs).

#### Breadth First Search - BFS

BFS is a non-recursive algorithm that employs a [queue](../queue) and is a generalization of [post-order traversal](../tree) in a tree. When provided with a graph G and a vertex V, BFS systematically explores all nodes in G that are reachable from V, as illustrated in the following example:

```Go
package graph_test

import "container/list"

type vertex struct {
		val      int
		edges    []*vertex
		distance int
}

func bfs(source *vertex) {
	distance := 0
	seen := make(map[*vertex]struct{})
	queue := list.New()
	queue.PushBack(source)

	for queue.Len() != 0 {
		tmp := queue.Front().Value.(*vertex)
		queue.Remove(queue.Front())
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
```

* Vertices at distance K are discovered before those at distance K + 1
* The shortest distance (i.e., the smallest number of edges) from the source vertex to each reachable vertex is computed
* A BFS tree with root V containing all reachable vertices is produced

For any vertex S that is reachable from V, the simple path in the BFS tree from V to S corresponds to a _shortest_ path (minimum number of edges) from V to S in G.

#### Depth First Search - DFS

Depth First Search (DFS) is a graph traversal algorithm that explores a graph by visiting as far as possible along each branch before backtracking. It uses a [stack](../stack) data structure when implemented iteratively, is [recursive](../recursion), and is a generalization of pre-order traversal in trees.

When given a graph G and a vertex S, DFS systematically discovers all nodes in G reachable from S. It is typically implemented using a driver that discovers the edges of the most recently discovered vertex V that has unexplored edges. Once all of V's edges have been explored, the search [backtracks](../backtracking/) to explore all edges leaving the vertex from which V was discovered. This process continues until the discovery of all edges.

```Go
package graph_test

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

func dfsDriver(graph []*timedVertex) {
	for _, v := range graph {
		if _, ok := seen[v]; !ok {
			dfs(v)
		}
	}
}

func dfs(u *timedVertex) {
	time++
	u.discoveryTimeStart = time
	for _, v := range u.edges {
		if _, ok := seen[v]; !ok {
			dfsVisit(v)
		}
	}
	seen[u] = struct{}{}
	time++
	u.discoveryTimeFinish = time
}
```

* Sets the status of edges at the time of discovery, and at the time exploring is finished
* Produces a depth-first forest comprising several depth-first trees
* Generates a well-formed parenthesis structure, where discovery of a vertex V is represented with (V and finishing discovering it is represented with V). For instance applying DFS on the graph (A) in _Figure 1_ would result in the output `(1(2(3(5))4()))`.

DFS is capable of categorizing edges (u,v) into four types:

* Tree Edge: v is discovered by exploring edge (u,v), all green edges
* Back edge: v is an ancestor of u, 6 to 2.
* Forward edge: v is a decedent of u, 1 to 8
* Cross edge: all other edges, 5 to 4

If a DFS algorithm identifies a back edge, it indicates that the graph is cyclic.

#### Dijkstra's Algorithm

A [greedy](../greedy) algorithm that uses BFS-like ideas and a minimum [heap](../heap) to solve single-source shortest path problems in edge weighted directed graphs like (D) in _Figure 1_.

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

            //distance type

            if v.distance + cvw < w.distance {
                //update w
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

### Iteratively Implement BFS and DFS

Implement BFS and DFS on a graph without using recursion. [Solution](iterative_traversal.go), [Tests](iterative_traversal_test.go)

### Is Graph a DAG

A directed acyclic graph (DAG) is a directed graph that has no cycles. Given a directed graph determine if it's a DAG or not. [Solution](is_dag.go), [Tests](is_dag_test.go)

### Topological Sort

In a DAG the topological order returns elements such that if there's a path from v(i) to v(j), then v(i) appears before v(j) in the ordering. Given a graph like (B) of _Figure 1_ output the graph in topological order like `{1,2,3,4}`. [Solution](topological_sort.go), [Tests](topological_sort_test.go)

### Employee Head Count

Head count of a person in an organization is 1 + the number of all their reports (direct and indirect). Given a list of employee IDs and their direct reports in each line like `1,4,5`, `5,7`, `4`, `7`. Where 1 has two direct reports (4 and 5); 5 has one report (7); 4 and 5 have no reports; Return the head count of a given employee ID. For example head count of 7 is 1, and headcount of 1 is 4. [Solution](employee_headcount.go), [Tests](employee_headcount_test.go)

### Remove Invalid Parentheses

Given a string containing parentheses and other alphabet letters like `(z)())()`, remove the minimum amount of parentheses to make the string valid like `(z())()` and `(z)()()`. [Solution](remove_invalid_parentheses.go), [Tests](remove_invalid_parentheses_test.go)

### Cheapest Flights

Given a list of flights each with a source and destination and a price, a maximum number of stops, source, and destination cities, return the cheapest costs not exceeding the maximum number of stops to reach from source city to destination. [Solution](cheapest_flights.go), [Tests](cheapest_flights_test.go)

### Word Ladder

Given a start word like `pop` and an end word like `car`, a dictionary of same length words like  `{"top","cop","cap","car"}` return the minimum number of transformations like 4 to get from start to end where each transformation between two words can happen when they are different by only one letter. [Solution](word_ladder.go), [Tests](word_ladder_test.go)

### Network Delay Time

Given `n`, the number of nodes in a network, travel times a list of `{source, destination, delay}` format, and `k` and a node number, return the time it will take for a message sent from k to be received by all nodes. [Solution](network_delay_time.go), [Tests](network_delay_time_test.go)

### Number of Islands

Given a grid in which 0 represents water and 1 represents a piece of land, return the number of islands. An Island is one or more piece of land that could be connected to other pieces of land on left, right, top or bottom. [Solution](number_of_islands.go), [Tests](number_of_islands_test.go)
