package main

import "fmt"

func main() {
	// // スライス ⇒ 可変長配列
	// var slice1 []int
	// slice2 := []int{}
	// fmt.Println("slice1:", slice1)
	// fmt.Println("slice2:", slice2)

	// // var := のどちらかによってslice
	// fmt.Println(slice1 == nil)
	// fmt.Println(slice2 == nil)

	// // appendで要素を追加
	// slice1 = append(slice1, 1, 2, 3)
	// fmt.Println("slice1:", slice1)

	// // 引数に展開する
	// slice3 := []int{4, 5, 6}
	// slice1 = append(slice1, slice3...)
	// fmt.Println(slice1)
	// fmt.Println("slice1:", slice1, len(slice1), cap(slice1))

	// // makeで配列を作成する
	slice4 := make([]int, 0, 2)
	fmt.Println("slice4:", slice4, len(slice4), cap(slice4))
	slice4 = append(slice4, 1, 2, 3, 4)
	fmt.Println("slice4:", slice4, len(slice4), cap(slice4))

	slice5 := make([]int, 4, 6)
	fmt.Println("slice5:", slice5, len(slice5), cap(slice5))

	// // スライスから分割してスライスを作成する
	// s6 := slice5[1:3]
	// s6[1] = 10
	// fmt.Println("slice5:", slice5, len(slice5), cap(slice5))
	// fmt.Println("s6:", s6, len(s6), cap(s6))

	// s6 = append(s6, 2)
	// fmt.Println("slice5 after append:", slice5, len(slice5), cap(slice5))
	// fmt.Println("s6 after append:", s6, len(s6), cap(s6))

	// fmt.Println("slice5 length:", len(slice5[1:3]))

	// // コピーの使い方
	// sc6 := make([]int, len(slice5[1:3]))
	// fmt.Println("slice5 copy:", slice5, len(slice5), cap(slice5))
	// fmt.Println("sc6 copy before:", sc6, len(sc6), cap(sc6))

	// copy(sc6, slice5[1:3])
	// fmt.Println("sc6 copy after:", sc6, len(sc6), cap(sc6))

	// sc6[1] = 12
	// fmt.Println("slice5:", slice5, len(slice5), cap(slice5))
	// fmt.Println("sc6:", sc6, len(sc6), cap(sc6))

	// slice5 = make([]int, 4, 6)
	// fs6 := slice5[1:3:3]
	// fmt.Println("slice5:", slice5, len(slice5), cap(slice5))
	// fmt.Println("fs6:", fs6, len(fs6), cap(fs6))
	// fs6[0] = 6
	// fs6[1] = 7
	// fs6 = append(fs6, 8)
	// fmt.Println("slice5:", slice5, len(slice5), cap(slice5))
	// fmt.Println("fs6:", fs6, len(fs6), cap(fs6))
	// slice5[3] = 9
	// fmt.Println("slice5:", slice5, len(slice5), cap(slice5))
	// fmt.Println("fs6:", fs6, len(fs6), cap(fs6))

	// // 	map ⇒ 連想配列
	// var m1 map[string]int
	// m2 := map[string]int{}
	// fmt.Println(m1, m1 == nil)
	// fmt.Println(m2, m2 == nil)
	// m2["A"] = 10
	// m2["B"] = 20
	// m2["C"] = 0
	// fmt.Println(m2, len(m2), m2["A"])

	// delete(m2, "A")
	// fmt.Println(m2, len(m2), m2["A"])

	// value, ok := m2["A"]
	// fmt.Println(value, ok)
	// value, ok = m2["C"]
	// fmt.Println(value, ok)

	// for key, value := range m2 {
	// 	fmt.Println(key, value)
	// }
}
