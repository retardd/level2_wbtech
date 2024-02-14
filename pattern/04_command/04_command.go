package main

import "fmt"

// Интерфейс Команды
type Command interface {
	Execute()
}

// Конкретная Команда
type ConcreteCommand struct {
	receiver *Receiver
}

func (c *ConcreteCommand) Execute() {
	c.receiver.Action()
}

// Получатель команды
type Receiver struct {
	Name string
}

func (r *Receiver) Action() {
	fmt.Printf("Receiver %s is performing the action\n", r.Name)
}

// Инициатор (Invoker)
type Invoker struct {
	command Command
}

func (i *Invoker) SetCommand(command Command) {
	i.command = command
}

func (i *Invoker) ExecuteCommand() {
	i.command.Execute()
}

func main() {
	// Создаем объекты
	receiver := &Receiver{Name: "Receiver-1"}
	concreteCommand := &ConcreteCommand{receiver: receiver}
	invoker := &Invoker{}

	// Устанавливаем команду для инициатора
	invoker.SetCommand(concreteCommand)

	// Инициатор выполняет команду
	invoker.ExecuteCommand()
}
