package main

//Дана структура Human (с произвольным набором полей и методов). Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).
import "fmt"

type Human struct {
	name   string
	age    int
	height int
	weight int
	dist   uint64
	water  uint64
	calor  uint64
}

func (h *Human) move() uint64 {
	h.dist = h.dist + 1
	return h.dist
}

func (h *Human) drink() uint64 {
	h.water = h.water + 1
	return h.water
}

func (h *Human) eat() uint64 {
	h.calor = h.calor + 1
	return h.calor
}

func (h Human) all_information() {
	fmt.Println("Name:", h.name)
	fmt.Println("Age:", h.age)
	fmt.Println("Height:", h.height)
	fmt.Println("Weight:", h.weight)
	fmt.Println("Distance:", h.dist)
	fmt.Println("Water quantity", h.water)
	fmt.Println("Calories:", h.calor)
}

type action struct {
	Human
}

func main() {
	man := action{Human{name: "Eugene", age: 20, height: 188, weight: 70, dist: 0, water: 0, calor: 0}} //Создание элемента man типа action с встроенным типом Human
	man.move()
	man.drink()
	man.eat()
	man.all_information()

}
