package Heaps

type MaxHeap struct {
	A		[]int
	size 	int
}

func New() (MaxHeap) {
	return MaxHeap{
		A: make([]int, 0),
		size: 0,
	}
}

func (m MaxHeap) parent(i int) int { return (i - 1) / 2 }
func (m MaxHeap) leftChild(i int) int { return (i * 2) + 1 }
func (m MaxHeap) rightChild(i int) int { return (i * 2) + 2 }

func (m *MaxHeap) heapifyUp(i int) {
	parentIndex := m.parent(i)
	if i >= 0 && m.A[parentIndex] < m.A[i] {
		m.A[i], m.A[parentIndex] = m.A[parentIndex], m.A[i]
		m.heapifyUp(parentIndex)
	}
}

func (m *MaxHeap) heapifyDown(i int) {
	left := m.leftChild(i)
	right := m.rightChild(i)

	largest := i 
	if left < m.size && m.A[left] > m.A[i] {
		largest = left 
	}
	if right < m.size && m.A[right] > m.A[largest] {
		largest = right 
	}

	if largest != i {
		m.A[i], m.A[largest] = m.A[largest], m.A[i]
		m.heapifyDown(largest)
	}
}

func (m *MaxHeap) Push(key int) {
	m.A = append(m.A, key)
	index := m.size 
	m.size++
	m.heapifyUp(index)
}

func (m *MaxHeap) Pop() {
	if m.IsEmpty() {
		return
	}
	m.A[0] = m.A[m.size - 1]
	m.A = m.A[:m.size - 1]
	m.size--
	m.heapifyDown(0)
}

func (m MaxHeap) Top() int {
	if m.IsEmpty() {
		return -1
	}
	return m.A[0]
}

func (m MaxHeap) IsEmpty() bool {
	return m.size == 0
}

func (m MaxHeap) GetSize() int {
	return m.size
}