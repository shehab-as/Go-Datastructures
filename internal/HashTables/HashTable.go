package HashTables

import "fmt"


type Map struct {
	array		[][]int
	size 		int 
}

func New(size int) Map {
	return Map{
		array: make([][]int, size),
		size: size,
	}
}

func (m Map) hash(key int) int {
	return key % m.size
}

func (m Map) keyExists(key int) bool {
	hashedKey := m.hash(key)
	for i := 0; i < len(m.array[hashedKey]); i++ {
		if key == m.array[hashedKey][i] {
			return true
		}
	}
	return false
}

func (m *Map) Insert(key int) {
	fmt.Printf("Inserting key %d.\n", key)
	if m.keyExists(key) {
		fmt.Printf("Key %d already exists.\n", key)
		return
	}
	hashedKey := m.hash(key)
	m.array[hashedKey] = append(m.array[hashedKey], key)
}

func (m *Map) Remove(key int) {
	fmt.Printf("Deleting key %d.\n", key)
	if !m.keyExists(key) {
		fmt.Printf("Key %d does not exist.\n", key)
		return
	}
	hashedKey := m.hash(key)
	length := len(m.array[hashedKey])
	for i := 0; i < length; i++ {
		if key == m.array[hashedKey][i] {
			// Swap with last element, then slice till before last element.
			m.array[hashedKey][length - 1], m.array[hashedKey][i] = m.array[hashedKey][i], m.array[hashedKey][length- 1]
			m.array[hashedKey] = m.array[hashedKey][:length - 1]
		}
	}
}

func (m Map) Display() {
	for i := 0; i < m.size; i++ {
		fmt.Printf("Key (%d) -> ", i)
		fmt.Printf("%+v\n", m.array[i])
	}
}

