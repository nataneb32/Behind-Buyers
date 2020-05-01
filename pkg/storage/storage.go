package storage

import "fmt"

type Storage interface {
	Store(group string, v interface{})
	Read(group string) interface{}
	GetRaw() map[string]interface{}
}

type InMemory struct {
	data map[string]interface{}
}

func (s *InMemory) Store(group string, v interface{}) {
	fmt.Println(v)
	fmt.Println(group)
	s.data[group] = v
}

func (s *InMemory) Read(group string) interface{} {
	return s.data[group]
}

func (s *InMemory) GetRaw() map[string]interface{} {
	return s.data
}

func NewInMemory() *InMemory {
	return &InMemory{
		data: make(map[string]interface{}),
	}
}
