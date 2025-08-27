package calculator

import "fmt"

// 小文字の変数は他パッケージにエクスポートされない
var offset float64 = 1

// 大文字の変数・関数は他パッケージにエクスポートされる
var Offset float64 = 2

func Sum(a float64, b float64) float64 {
	fmt.Println("multiply: ", multiply(a, b))
	fmt.Println("Multiply: ", Multiply(a, b))

	return a + b + offset
}
