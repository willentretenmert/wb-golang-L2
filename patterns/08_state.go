package main

import "fmt"

// Паттерн "Состояние" используется в программировании для изменения поведения объекта
// при изменении его внутреннего состояния.
// По сути, он позволяет объекту изменять свое поведение во время выполнения, когда его внутреннее состояние меняется.

type State interface {
	RespondLock(*BigSteelLock)
}

type BigSteelLock struct {
	locked       State
	unlocked     State
	currentState State
}

type LockedState struct{}

type UnlockedState struct{}

func (s *LockedState) RespondLock() {
	fmt.Println("Attempting to unlock...")
}

func (s *UnlockedState) RespondLock() {
	fmt.Println("Attempting to lock...")
}

func (b *BigSteelLock) setState(newState State) {
	b.currentState = newState
}

func (b *BigSteelLock) PressLock() {
	b.currentState.RespondLock(b)
}
