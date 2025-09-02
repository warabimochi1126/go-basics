package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// uint16 ⇒ 2byteのデータ型
	var ui1 uint16
	// 変数の先頭メモリアドレスを取得する ⇒ &を付ける
	fmt.Println("ui1 address", &ui1)
	// メモリアドレスの先頭が2Byteズレている事からuint16が2Byteのデータ型であることが分かる
	var ui2 uint16
	fmt.Println("ui2 address", &ui2)

	// ポインタ変数は変数の先頭メモリアドレスを持っている
	// 型を指定するのは、先頭アドレスから何byte分かを指定するため
	var p1 *uint16
	fmt.Println("p1 address", p1)
	p1 = &ui1
	fmt.Println("p1 address", p1)
	fmt.Println("p1 size", unsafe.Sizeof(p1))
	fmt.Println("p1 oneself address:", &p1)
	// ポインタ変数が指すメモリアドレスの値を取得する(dereference)
	fmt.Println("p1 pointer next value", *p1)
	*p1 = 1
	fmt.Println("ui1 value", ui1)

	// ダブルポインタ ⇒ ポインタ変数が更にポインタを持つ状態
	var pp1 **uint16 = &p1
	fmt.Println("pp1 value", pp1)
	fmt.Println("pp1 address", &pp1)
	fmt.Println("pp1 size", unsafe.Sizeof(pp1))

	fmt.Println("pp1 one dereference", *pp1)
	fmt.Println("pp1 two dereference", **pp1)

	**pp1 = 10
	fmt.Println("ui1 value", ui1)

	// shadowing ⇒ 別のスコープで変数を再定義することで、外側の変数にアクセスできなくする手法
	ok, result := true, "A"
	fmt.Println("outer result pointer", &result)
	if ok {
		result = "B"
		fmt.Println("innner result pointer", &result)
		println(result)
	} else {
		result := "C"
		println(result)
	}
	println(result)
}
