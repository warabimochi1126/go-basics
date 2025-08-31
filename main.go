package main

// constは定数宣言
const secret int = 100

type Os int

// 連番の生成
const (
	Mac Os = iota + 1
	Windows
	Linux
)

// 複数の変数定義
var (
	i int
	s string
	b bool
) // i ⇒ 0 s ⇒ "" b ⇒ false

func main() {
	// // # 初期化しなかった場合、型のfalsyな値が挿入される
	// var i int
	// // # varでも初期化できる
	// var i int = 2
	// // # varで型省略すると代入値から型推論される
	// var i = 3
	// # := は初期化によって型推論される
	// i := 1 // i(int) ⇒ 1
	// # キャストすることで型を付けられる
	// ui := uint16(2)

	// Printfは書式を指定して出力する
	// %vは値の表示・%Tは型の表示
	// fmt.Printf("i: %v %T\n", i, ui)
	// fmt.Printf("i: %[1]v %[1]T ui: %[2]v %[2]T", i, ui)

	// # 複数個の変数を一気に初期化することが出来る
	// pi, title := 3.14, "go"
	// fmt.Printf("pi: %v title: %v", pi, title)
}
