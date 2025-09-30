package main

import "fmt"

// メソッド群を一覧に出来る ← interface
type ageIncrementer interface {
	ageUp() int
}

type human1 struct {
	age int
}
type human2 struct {
	age int
}

// 構造体にinterfaceと同名のメソッドを実装する
func (human *human1) ageUp() int {
	human.age += 1
	return human.age
}
func (human *human2) ageUp() int {
	human.age += 2
	return human.age
}

func (human *human1) String() string {
	return fmt.Sprint("String method")
}

func ageUp(human ageIncrementer) {
	fmt.Println("human Age:", human.ageUp())
}

func main() {
	human1 := &human1{1}
	ageUp(human1) // human Age: 2
	human2 := &human2{1}
	ageUp(human2) // human Age: 3

	fmt.Println(human1) // String method
	fmt.Println(human2) // &{3}
}
