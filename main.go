package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

// deferを付けると関数実行終了時に実行される.複数ある場合はLIFOで実行される
func funcDefer() {
	defer fmt.Println("main func final-finish")
	defer fmt.Println("main func semi-finish")
	fmt.Println("hello world")
}

// 引数の型に...をつけるとsliceとして受け取る
func trimExtention(files ...string) []string {
	out := make([]string, 0, len(files))
	for _, file := range files {
		out = append(out, strings.TrimSuffix(file, ".csv"))
	}
	return out
}

// 戻り値を複数返す時は、タプルみたいに書く
func fileChecker(name string) (string, error) {
	f, err := os.Open(name)
	if err != nil {
		return "", errors.New("file not Found")
	}
	defer f.Close()
	return name, nil
}

// 無名関数を引数として受け取る
func addExtention(f func(file string) string, name string) {
	fmt.Println(f(name))
}

// 戻り値で関数を返す
func multiply() func(int) int {
	return func(n int) int {
		return n * 1000
	}
}

func countUp() func(int) int {
	count := 0
	return func(n int) int {
		count += n
		return count
	}
}

func main() {
	funcDefer()

	files := []string{"file1.csv", "file2.csv", "flie3.csv"}
	// スプレッド演算子でsliceを展開し、複数の文字列として渡せる
	fmt.Println(trimExtention(files...))

	name, err := fileChecker("file.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(name)

	// 無名関数の作り方
	i := 1
	func(i int) {
		fmt.Println(i)
	}(i)
	// 無名関数を変数に代入出来る.アロー関数を変数に代入する感覚
	f1 := func(i int) int {
		return i + 1
	}
	fmt.Println(f1(i))

	f2 := func(file string) string {
		return file + ".csv"
	}
	addExtention(f2, "file2")

	// 関数を戻り値として受け取る
	f3 := multiply()
	fmt.Println(f3(2))

	f4 := countUp()
	for i := 1; i <= 5; i++ {
		v := f4(2)
		fmt.Println(v)
	}
}
