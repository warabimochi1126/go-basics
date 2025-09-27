package main

import "fmt"

type Task struct {
	Title    string
	Estimate int
}

func main() {
	task1 := Task{
		Title:    "Learn Golang",
		Estimate: 3,
	}
	task1.Title = "Learning Go"
	fmt.Println(task1, task1.Title)

	var task2 Task = task1
	fmt.Println(task2)
	task2.Title = "new"
	fmt.Println(task2, "task1.title:", task1.Title, "task2.title:", task2.Title)

	task1Pointer := &Task{
		Title:    "Learn concurrency",
		Estimate: 2,
	}
	fmt.Println(task1Pointer, *task1Pointer)
	task1Pointer.Title = "Changed"
	fmt.Println(*task1Pointer)

	var task2Pointer *Task = task1Pointer
	task2Pointer.Title = "Changed by Task2"
	fmt.Println(*task2Pointer)
	fmt.Println(*task1Pointer)

	task1.extendEstimate()
	fmt.Println("task1 estimate:", task1.Estimate)
	task1.extendEstimatePointer()
	fmt.Println("task1 estimate:", task1.Estimate)
}

// 値レシーバーを持つメソッド
// 特定の型の実体に対してメソッドを追加出来る
// 値レシーバーの引数は値渡しなので、元の構造体に影響を与えない
func (task Task) extendEstimate() {
	task.Estimate += 10
}

// ポインターレシーバーを持つメソッド
// ポインターレシーバーの引数は参照渡しなので、元の構造体に影響を与える
func (taskp *Task) extendEstimatePointer() {
	taskp.Estimate += 10
}
