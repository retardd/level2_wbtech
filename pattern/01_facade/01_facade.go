package main

import "fmt"

// SubsystemA представляет часть подсистемы A
type SubsystemA struct {
}

func (s *SubsystemA) OperationA() {
	fmt.Println("Subsystem A: Operation A")
}

// SubsystemB представляет часть подсистемы B
type SubsystemB struct {
}

func (s *SubsystemB) OperationB() {
	fmt.Println("Subsystem B: Operation B")
}

// SubsystemC представляет часть подсистемы C
type SubsystemC struct {
}

func (s *SubsystemC) OperationC() {
	fmt.Println("Subsystem C: Operation C")
}

// Facade предоставляет унифицированный интерфейс к подсистеме
type Facade struct {
	A *SubsystemA
	B *SubsystemB
	C *SubsystemC
}

func NewFacade() *Facade {
	return &Facade{
		A: &SubsystemA{},
		B: &SubsystemB{},
		C: &SubsystemC{},
	}
}

// Operation предоставляет унифицированный интерфейс к операциям подсистемы
func (f *Facade) Operation() {
	fmt.Println("Facade: Operation")
	f.A.OperationA()
	f.B.OperationB()
	f.C.OperationC()
}

func main() {
	// Используем фасад для упрощения работы с подсистемой
	facade := NewFacade()
	facade.Operation()
}
