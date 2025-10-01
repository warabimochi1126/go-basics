package main

import "fmt"

// interface ⇒ メソッドの一覧を定義

func main() {
	items := []item{
		{price: 10},
		{price: 20},
		{price: 30},
	}

	for index, value := range items {
		fmt.Println(index) // 0, 1, 2
		fmt.Println(value) // [{10}, {20}, {30}]
	}

	for _, value := range items {
		value.price = 100
	}
	fmt.Println(items) // [{10}, {20}, {30}]

	for index := range items {
		items[index].price = 100
	}
	fmt.Println(items) // [{100}, {100}, {100}]
}

type item struct {
	price float32
}
