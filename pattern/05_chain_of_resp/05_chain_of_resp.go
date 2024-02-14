package main

import "fmt"

// Интерфейс обработчика
type Handler interface {
	SetNext(handler Handler)
	HandleRequest(request int)
}

// Конкретный обработчик
type ConcreteHandler struct {
	next Handler
}

func (c *ConcreteHandler) SetNext(handler Handler) {
	c.next = handler
}

func (c *ConcreteHandler) HandleRequest(request int) {
	if request <= 10 {
		fmt.Println("Request handled by ConcreteHandler")
	} else if c.next != nil {
		fmt.Println("Request passed to next handler")
		c.next.HandleRequest(request)
	} else {
		fmt.Println("No handler available")
	}
}

func main() {
	// Создаем цепочку обработчиков
	handler1 := &ConcreteHandler{}
	handler2 := &ConcreteHandler{}
	handler3 := &ConcreteHandler{}

	// Устанавливаем следующий обработчик для каждого обработчика
	handler1.SetNext(handler2)
	handler2.SetNext(handler3)

	// Обработка запросов
	handler1.HandleRequest(5)
	handler1.HandleRequest(15)
}
