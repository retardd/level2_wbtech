package main

import "fmt"

// Element интерфейс, представляющий элемент, который может быть посещен
type Element interface {
	Accept(visitor Visitor)
}

// ConcreteElementA и ConcreteElementB конкретные элементы, которые реализуют интерфейс Element
type ConcreteElementA struct {
	Name string
}

func (e *ConcreteElementA) Accept(visitor Visitor) {
	visitor.VisitConcreteElementA(e)
}

type ConcreteElementB struct {
	Value int
}

func (e *ConcreteElementB) Accept(visitor Visitor) {
	visitor.VisitConcreteElementB(e)
}

// Visitor интерфейс, представляющий посетителя
type Visitor interface {
	VisitConcreteElementA(element *ConcreteElementA)
	VisitConcreteElementB(element *ConcreteElementB)
}

// ConcreteVisitor реализует конкретного посетителя
type ConcreteVisitor struct{}

func (v *ConcreteVisitor) VisitConcreteElementA(element *ConcreteElementA) {
	fmt.Printf("Посетитель выполняет операцию в Элементе А: %s\n", element.Name)
}

func (v *ConcreteVisitor) VisitConcreteElementB(element *ConcreteElementB) {
	fmt.Printf("Посетитель выполняет операцию в Элементе B: %d\n", element.Value)
}

// ObjectStructure структура объектов, которые могут быть посещены (можно просто слайсом в main func)
type ObjectStructure struct {
	elements []Element
}

func (os *ObjectStructure) AddElement(element Element) {
	os.elements = append(os.elements, element)
}

func (os *ObjectStructure) Accept(visitor Visitor) {
	for _, element := range os.elements {
		element.Accept(visitor)
	}
}

func main() {
	// Создаем объекты, которые могут быть посещены
	elementA := &ConcreteElementA{Name: "AAAA"}
	elementB := &ConcreteElementB{Value: 1111}

	// Добавляем объекты в структуру
	objectStructure := &ObjectStructure{}
	objectStructure.AddElement(elementA)
	objectStructure.AddElement(elementB)

	// Создаем посетителя
	visitor := &ConcreteVisitor{}

	// Посетитель посещает объекты в структуре
	objectStructure.Accept(visitor)
}
