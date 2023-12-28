package main

import "fmt"

// Этот паттерн используется для создания объектов, передавая логику создания из конструктора в отдельный метод.
// В данном случае фабричный метод реализован через функцию New.

const (
	ComputerType string = "Computer"
	PhoneType    string = "phone"
)

type StoreObject interface {
	GetType() string
	PrintDetails()
}

type Phone struct {
	Display      string
	Manufacturer string
	Frameless    bool
}

type Computer struct {
	CPU string
	RAM int
}

func NewComputer() Computer {
	return Computer{
		CPU: "Intel",
		RAM: 32,
	}
}

func NewPhone() Phone {
	return Phone{
		Display:      "6.3\"",
		Manufacturer: "Apple",
		Frameless:    true,
	}
}

func (n Phone) GetType() string {
	return "Phone"
}

func (n Phone) PrintDetails() {
	fmt.Printf("Display %s, Manufacturer %s, Frameless %v\n", n.Display, n.Manufacturer, n.Frameless)
}

func (s Computer) GetType() string {
	return "Computer"
}

func (s Computer) PrintDetails() {
	fmt.Printf("CPU %s, RAM %d\n", s.CPU, s.RAM)
}

func New(typeName string) StoreObject {
	switch typeName {
	default:
		fmt.Printf("Несуществующий тип %s\n", typeName)
		return nil
	case ComputerType:
		return NewComputer()
	case PhoneType:
		return NewPhone()
	}
}
