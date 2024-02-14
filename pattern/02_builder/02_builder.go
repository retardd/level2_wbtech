package main

import "fmt"

// Product представляет конечный продукт, который мы строим
type Product struct {
	PartA string
	PartB string
	PartC string
}

// Builder предоставляет интерфейс для создания различных частей продукта
type Builder interface {
	BuildPartA()
	BuildPartB()
	BuildPartC()
	GetProduct() *Product
}

// ConcreteBuilder реализует интерфейс Builder и строит конкретный продукт
type ConcreteBuilder struct {
	product *Product
}

func NewConcreteBuilder() *ConcreteBuilder {
	return &ConcreteBuilder{
		product: &Product{},
	}
}

func (b *ConcreteBuilder) BuildPartA() {
	b.product.PartA = "PartA"
}

func (b *ConcreteBuilder) BuildPartB() {
	b.product.PartB = "PartB"
}

func (b *ConcreteBuilder) BuildPartC() {
	b.product.PartC = "PartC"
}

func (b *ConcreteBuilder) GetProduct() *Product {
	return b.product
}

// Director управляет процессом строительства и скрывает сложность клиентского кода
type Director struct {
	builder Builder
}

func NewDirector(builder Builder) *Director {
	return &Director{
		builder: builder,
	}
}

func (d *Director) Construct() {
	d.builder.BuildPartA()
	d.builder.BuildPartB()
	d.builder.BuildPartC()
}

func main() {
	// Создаем строителя
	builder := NewConcreteBuilder()

	// Создаем директора и передаем ему строителя
	director := NewDirector(builder)

	// Директор управляет процессом строительства
	director.Construct()

	// Получаем готовый продукт от строителя
	product := builder.GetProduct()

	// Выводим результат
	fmt.Printf("Product Parts: %s, %s, %s\n", product.PartA, product.PartB, product.PartC)
}
