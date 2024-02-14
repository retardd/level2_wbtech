package main

import "fmt"

// Интерфейс продукта
type Product interface {
	Use()
}

// Конкретный продукт A
type ConcreteProductA struct{}

func (p *ConcreteProductA) Use() {
	fmt.Println("Using ConcreteProductA")
}

// Конкретный продукт B
type ConcreteProductB struct{}

func (p *ConcreteProductB) Use() {
	fmt.Println("Using ConcreteProductB")
}

// Интерфейс фабрики
type Factory interface {
	CreateProduct() Product
}

// Конкретная фабрика A
type ConcreteFactoryA struct{}

func (f *ConcreteFactoryA) CreateProduct() Product {
	return &ConcreteProductA{}
}

// Конкретная фабрика B
type ConcreteFactoryB struct{}

func (f *ConcreteFactoryB) CreateProduct() Product {
	return &ConcreteProductB{}
}

func main() {
	// Создаем объект фабрики A
	factoryA := &ConcreteFactoryA{}
	// Используем фабрику A для создания продукта A
	productA := factoryA.CreateProduct()
	productA.Use()

	// Создаем объект фабрики B
	factoryB := &ConcreteFactoryB{}
	// Используем фабрику B для создания продукта B
	productB := factoryB.CreateProduct()
	productB.Use()
}
