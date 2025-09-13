package main

import "fmt"

func main() {
	// 配列をやる
	var a1 [3]int
	fmt.Println(a1, len(a1), cap(a1))

	// 宣言時に代入
	var a2 = [3]int{10, 20, 30}
	fmt.Println(a2, len(a2), cap((a2)))

	// 代入する値の数を要素数とする
	var a3 = [...]int{100, 200}
	fmt.Println(a3, len(a3), cap(a3))

	// := での宣言では代入が必須
	a4 := [3]int{300, 400}
	fmt.Println(a4, len(a4), cap(a4))

	// スライス ⇒ 可変長配列
	var s1 []int
	s2 := []int{}
	fmt.Println("s1:", s1, len(s1), cap(s1))
	fmt.Println("s2:", s2, len(s2), cap(s2))

	// var := のどちらかによってslice
	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)

	// appendで要素を追加
	s1 = append(s1, 1, 2, 3)
	fmt.Println("s1:", s1, len(s1), cap(s1))

	// 引数に展開する
	s3 := []int{4, 5, 6}
	s1 = append(s1, s3...)
	fmt.Println("s1:", s1, len(s1), cap(s1))

	// makeで配列を作成する
	s4 := make([]int, 0, 2)
	fmt.Println("s4:", s4, len(s4), cap(s4))
	s4 = append(s4, 1, 2, 3, 4)
	fmt.Println("s4:", s4, len(s4), cap(s4))

	s5 := make([]int, 4, 6)
	fmt.Println("s5:", s5, len(s5), cap(s5))

	// スライスから分割してスライスを作成する
	s6 := s5[1:3]
	s6[1] = 10
	fmt.Println("s5:", s5, len(s5), cap(s5))
	fmt.Println("s6:", s6, len(s6), cap(s6))

	s6 = append(s6, 2)
	fmt.Println("s5 after append:", s5, len(s5), cap(s5))
	fmt.Println("s6 after append:", s6, len(s6), cap(s6))

	fmt.Println("s5 length:", len(s5[1:3]))

	// コピーの使い方
	sc6 := make([]int, len(s5[1:3]))
	fmt.Println("s5 copy:", s5, len(s5), cap(s5))
	fmt.Println("sc6 copy before:", sc6, len(sc6), cap(sc6))

	copy(sc6, s5[1:3])
	fmt.Println("sc6 copy after:", sc6, len(sc6), cap(sc6))

	sc6[1] = 12
	fmt.Println("s5:", s5, len(s5), cap(s5))
	fmt.Println("sc6:", sc6, len(sc6), cap(sc6))

	s5 = make([]int, 4, 6)
	fs6 := s5[1:3:3]
	fmt.Println("s5:", s5, len(s5), cap(s5))
	fmt.Println("fs6:", fs6, len(fs6), cap(fs6))
	fs6[0] = 6
	fs6[1] = 7
	fs6 = append(fs6, 8)
	fmt.Println("s5:", s5, len(s5), cap(s5))
	fmt.Println("fs6:", fs6, len(fs6), cap(fs6))
	s5[3] = 9
	fmt.Println("s5:", s5, len(s5), cap(s5))
	fmt.Println("fs6:", fs6, len(fs6), cap(fs6))

	// 	map ⇒ 連想配列
	var m1 map[string]int
	m2 := map[string]int{}
	fmt.Println(m1, m1 == nil)
	fmt.Println(m2, m2 == nil)
	m2["A"] = 10
	m2["B"] = 20
	m2["C"] = 0
	fmt.Println(m2, len(m2), m2["A"])

	delete(m2, "A")
	fmt.Println(m2, len(m2), m2["A"])

	value, ok := m2["A"]
	fmt.Println(value, ok)
	value, ok = m2["C"]
	fmt.Println(value, ok)

	for key, value := range m2 {
		fmt.Println(key, value)
	}
}
