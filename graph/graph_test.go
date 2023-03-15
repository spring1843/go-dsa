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
			dfsVisit(v)
		}
	}
}

func dfsVisit(u *timedVertex) {
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
