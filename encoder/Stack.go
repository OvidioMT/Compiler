package encoder


import (
"errors"
"github.com/llir/llvm/ir"
)

// Stack define una estructura de pila
type Stack struct {
	items []*ir.Block
}

// Push agrega un elemento al tope de la pila
func (s *Stack) Push(item *ir.Block) {
	s.items = append(s.items, item)
}

// Pop elimina y devuelve el elemento al tope de la pila
func (s *Stack) Pop() (*ir.Block, error) {
	if len(s.items) == 0 {
		return nil, errors.New("la pila está vacía")
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, nil
}

// Peek devuelve el elemento al tope de la pila sin eliminarlo
func (s *Stack) Peek() (*ir.Block, error) {
	if len(s.items) == 0 {
		return nil, errors.New("la pila está vacía")
	}
	return s.items[len(s.items)-1], nil
}

// GetFromTop devuelve el elemento en el índice especificado desde el tope de la pila
func (s *Stack) GetFromTop(index int) (*ir.Block, error) {
	if index < 0 || index >= len(s.items) {
		return nil, errors.New("índice fuera de rango")
	}
	return s.items[len(s.items)-1-index], nil
}

// IsEmpty verifica si la pila está vacía
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

// Size devuelve el número de elementos en la pila
func (s *Stack) Size() int {
	return len(s.items)
}

// Peek devuelve el elemento al tope de la pila sin eliminarlo
func (s *Stack) Peek2(valor int) (*ir.Block, error) {
	if len(s.items) == 0 {
		return nil, errors.New("la pila está vacía")
	}
	return s.items[len(s.items)-valor], nil
}
