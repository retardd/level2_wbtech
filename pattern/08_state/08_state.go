package main

import "fmt"

// Интерфейс состояния
type State interface {
	On(Printer) error
	Off(Printer) error
	Print(Printer) error
}

// Конкретное состояние - Включена
type OnState struct{}

func (s *OnState) On(p Printer) error {
	return fmt.Errorf("printer is already on")
}

func (s *OnState) Off(p Printer) error {
	fmt.Println("Turning printer off")
	p.setState(&OffState{})
	return nil
}

func (s *OnState) Print(p Printer) error {
	fmt.Println("Printing...")
	p.setState(&PrintState{})
	return nil
}

// Конкретное состояние - Выключена
type OffState struct{}

func (s *OffState) On(p Printer) error {
	fmt.Println("Turning printer on")
	p.setState(&OnState{})
	return nil
}

func (s *OffState) Off(p Printer) error {
	return fmt.Errorf("printer is already off")
}

func (s *OffState) Print(p Printer) error {
	return fmt.Errorf("cannot print, printer is off")
}

// Конкретное состояние - Печать
type PrintState struct{}

func (s *PrintState) On(p Printer) error {
	return fmt.Errorf("cannot turn on while printing")
}

func (s *PrintState) Off(p Printer) error {
	fmt.Println("Cancelling print and turning printer off")
	p.setState(&OffState{})
	return nil
}

func (s *PrintState) Print(p Printer) error {
	return fmt.Errorf("already printing")
}

// Контекст - принтер
type Printer struct {
	state State
}

func NewPrinter() *Printer {
	return &Printer{state: &OffState{}}
}

func (p *Printer) setState(state State) {
	p.state = state
}

func (p *Printer) On() error {
	return p.state.On(*p)
}

func (p *Printer) Off() error {
	return p.state.Off(*p)
}

func (p *Printer) Print() error {
	return p.state.Print(*p)
}

func main() {
	printer := NewPrinter()

	// Попробуем включить принтер
	printer.On()
	printer.On() // Попытка повторно включить принтер

	// Включаем печать
	printer.Print()

	// Попробуем выключить принтер во время печати
	printer.Off()

	// Отменяем печать и выключаем принтер
	printer.Off()
}
