// в этом коде у нас структура Human с некоторыми полями и методом Sleep
// затем определяем структуру Action, которая эмбедит в себя структуру Human
// это позволяет инстансам Action получать доступ к полям и методам Human,
// как будто они собственные

package main

import "fmt"

// определяем структуру Human с полями
type Human struct {
	Name   string
	Age    int
	Gender string
	Weight float32
}

// метод структуры Human
func (h *Human) Sleep() {
	fmt.Printf("%s is sleeping.\n", h.Name)
}

// определяем структуру Action, встраивая в нее Human
type Action struct {
	Human // "наследуем" поля и методы Human
	Type  string
}

// метод структуры Action
func (a *Action) Do() {
	fmt.Printf("%s is doing a %s action.\n", a.Name, a.Type)
}

func main() {
	// создаем инстанс структуры Action
	action := Action{
		Human: Human{
			Name:   "Alice",
			Age:    30,
			Gender: "Female",
			Weight: 54.7,
		},
		Type: "Run",
	}

	// вызываем методы
	action.Sleep()
	action.Do()
}
