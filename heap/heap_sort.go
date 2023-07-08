package heap

type (
	// Vertex is a node in a heap
	Vertex struct {
		// Val is the value of the vertex
		Val int
		// Left is the left child of the vertex
		Left *Vertex
		// Right is the right child of the vertex
		Right *Vertex
	}
	// MinHeap is a heap where the root is always the minimum value
	MinHeap struct {
		Data []*Vertex
	}
)

// HeapSort solves the problem in O(n*Log n) time and O(n) space.
func HeapSort(list []int) []int {
	sorted := []int{}
	heap := NewMinHeap()
	for _, val := range list {
		heap.Push(val)
	}
	for heap.Len() > 0 {
		sorted = append(sorted, heap.Pull())
	}
	return sorted
}

// Returns a new Min Heap
func NewMinHeap() *MinHeap {
	return &MinHeap{
		Data: []*Vertex{},
	}
}

// Push inserts a new value into the heap
func (m *MinHeap) Push(value int) {
	vertex := &Vertex{
		Val: value,
	}
	m.Data = append(m.Data, vertex)
	m.heapifyUp(len(m.Data) - 1)
}

// Pull removes the root value from the heap
func (m *MinHeap) Pull() int {
	if len(m.Data) == 0 {
		return 0 // or any other appropriate default value
	}

	rootValue := m.Data[0].Val
	lastIndex := len(m.Data) - 1
	m.Data[0] = m.Data[lastIndex]
	m.Data = m.Data[:lastIndex]
	m.heapifyDown(0)
	return rootValue
}

// Len returns the number of elements in the heap
func (m *MinHeap) Len() int {
	return len(m.Data)
}

// heapifyUp moves the vertex up the heap to maintain the heap property
func (m *MinHeap) heapifyUp(index int) {
	for index > 0 {
		parentIndex := (index - 1) / 2
		if m.Data[parentIndex].Val <= m.Data[index].Val {
			break
		}
		m.swap(parentIndex, index)
		index = parentIndex
	}
}

// heapifyDown moves the vertex down the heap to maintain the heap property
func (m *MinHeap) heapifyDown(index int) {
	for {
		leftChildIndex := 2*index + 1
		rightChildIndex := 2*index + 2
		smallestIndex := index

		if leftChildIndex < len(m.Data) && m.Data[leftChildIndex].Val < m.Data[smallestIndex].Val {
			smallestIndex = leftChildIndex
		}

		if rightChildIndex < len(m.Data) && m.Data[rightChildIndex].Val < m.Data[smallestIndex].Val {
			smallestIndex = rightChildIndex
		}

		if smallestIndex == index {
			break
		}

		m.swap(index, smallestIndex)
		index = smallestIndex
	}
}

// swap swaps the positions of two vertices in the heap
func (m *MinHeap) swap(i, j int) {
	m.Data[i], m.Data[j] = m.Data[j], m.Data[i]
}
