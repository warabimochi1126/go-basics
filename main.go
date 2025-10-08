package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// 型に~を付けると、その型を基底型とした独自の型も許容する
// ex) この場合int型を基底型とした独自の型も許容する
type customConstraints interface {
	~int | int16 | float32 | float64 | string
}

// int型を基底型として独自の型を定義
type NewInt int

func add[T customConstraints](x, y T) T {
	return x + y
}

func min[T constraints.Ordered](x, y T) T {
	if x < y {
		return x
	}
	return y
}
func sumValues[K int | string, V constraints.Float | constraints.Integer](m map[K]V) V {
	var sum V
	for _, v := range m {
		sum += v
	}
	return sum
}

func main() {
	fmt.Println(add(1, 2))
	fmt.Println(add(1.1, 2.1))
	fmt.Println(add("file", ".txt"))

	var int1, int2 NewInt = 3, 4
	fmt.Println(add(int1, int2))
	fmt.Println(min(int1, int2))

	m1 := map[string]uint{
		"A": 1,
		"B": 2,
		"C": 3,
	}
	m2 := map[int]float32{
		1: 1.23,
		2: 4.56,
		3: 7.89,
	}
	fmt.Println(sumValues(m1))
	fmt.Println(sumValues(m2))
}
