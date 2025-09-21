package main

func main() {
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
	// fslice6 := slice5[1:3:3]
	// fmt.Println("slice5:", slice5, len(slice5), cap(slice5))
	// fmt.Println("fslice6:", fslice6, len(fslice6), cap(fslice6))
	// fslice6[0] = 6
	// fslice6[1] = 7
	// fslice6 = append(fslice6, 8)
	// fmt.Println("slice5:", slice5, len(slice5), cap(slice5))
	// fmt.Println("fslice6:", fslice6, len(fslice6), cap(fslice6))
	// slice5[3] = 9
	// fmt.Println("slice5:", slice5, len(slice5), cap(slice5))
	// fmt.Println("fslice6:", fslice6, len(fslice6), cap(fslice6))

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
