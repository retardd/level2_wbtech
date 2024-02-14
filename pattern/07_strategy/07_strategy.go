package main

import "fmt"

// Интерфейс стратегии
type Strategy interface {
	ExecuteStrategy(int, int) int
}

// Конкретная стратегия для сложения
type ConcreteStrategyAdd struct{}

func (s *ConcreteStrategyAdd) ExecuteStrategy(a, b int) int {
	return a + b
}

// Конкретная стратегия для вычитания
type ConcreteStrategySubtract struct{}

func (s *ConcreteStrategySubtract) ExecuteStrategy(a, b int) int {
	return a - b
}

// Контекст, который использует стратегию
type Context struct {
	strategy Strategy
}

func (c *Context) SetStrategy(strategy Strategy) {
	c.strategy = strategy
}

func (c *Context) ExecuteStrategy(a, b int) int {
	return c.strategy.ExecuteStrategy(a, b)
}

func main() {
	// Создаем контекст
	context := &Context{}

	// Устанавливаем стратегию сложения
	context.SetStrategy(&ConcreteStrategyAdd{})
	// Выполняем операцию
	result := context.ExecuteStrategy(10, 5)
	fmt.Println("10 + 5 =", result)

	// Устанавливаем стратегию вычитания
	context.SetStrategy(&ConcreteStrategySubtract{})
	// Выполняем операцию
	result = context.ExecuteStrategy(10, 5)
	fmt.Println("10 - 5 =", result)
}
