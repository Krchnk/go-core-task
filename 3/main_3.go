package main

type StringIntMap struct {
	data map[string]int
}

func NewStringIntMap() *StringIntMap {
	return &StringIntMap{
		data: make(map[string]int),
	}
}

func (m *StringIntMap) Add(key string, value int) {
	m.data[key] = value
}

func (m *StringIntMap) Remove(key string) {
	delete(m.data, key)
}

func (m *StringIntMap) Copy() map[string]int {
	newMap := make(map[string]int)
	for key, value := range m.data {
		newMap[key] = value
	}
	return newMap
}

func (m *StringIntMap) Exists(key string) bool {
	_, exists := m.data[key]
	return exists
}

func (m *StringIntMap) Get(key string) (int, bool) {
	value, exists := m.data[key]
	return value, exists
}
