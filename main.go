// goのファイルは何かしらのpackageに属する必要がある
// main packageは初期状態から存在する
package main

import "fmt"

// main関数はファイル実行時に勝手に実行される関数
func main() {
	fmt.Println("hello world")
	sl := []int{1, 2, 3}
	if len(sl) < 0 {
		fmt.Println("unreachable code")
	}
}
