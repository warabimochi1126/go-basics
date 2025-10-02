package main

import (
	"errors"
	"fmt"
	"os"
)

var ErrCustom = errors.New("not found")

func main() {
	// errors.NewでErrorオブジェクトを作る
	// error01 := errors.New("something wrong")
	// fmt.Println("error01 value:", error01)
	// fmt.Println("error01 type:", reflect.TypeOf(error01))

	// error02 := errors.New("something wrong")
	// fmt.Println("error pointer match:", error01 == error02)
	// fmt.Println("error01 pointer address:", &error01)
	// fmt.Println("error02 pointer address:", &error02)

	// fmt.ErrorfでErrorオブジェクトを作る
	// errorOf := fmt.Errorf("aiueo")
	// fmt.Println("errorOf value:", errorOf)
	// fmt.Println("errorOf type:", reflect.TypeOf(errorOf))
	// fmt.Errorfは%wでエラーをラップする
	// errorWrap := fmt.Errorf("twoHikisuu=%w", errorOf)
	// fmt.Println("errorWrap value:", errorWrap)
	// fmt.Println("errorWrap type:", reflect.TypeOf(errorWrap))
	// %vだとエラーをラップしない
	// errorNoWrap := fmt.Errorf("No Wrap=%v", errorOf)
	// fmt.Println("errorNoWrap value:", errorNoWrap)
	// fmt.Println("errorNoWrap type:", reflect.TypeOf(errorNoWrap))

	// ラップしたエラーの元エラーを辿る ⇒ あれば返ってくる.無ければnil
	// fmt.Println("unWrap errorOf:", errors.Unwrap(errorOf))
	// fmt.Println("unWrap errorWrap:", errors.Unwrap(errorWrap))
	// fmt.Println("unWrap errorNoWrap:", errors.Unwrap(errorNoWrap))

	// センチネルエラー ⇒ 予め定義されているerrorString構造体
	// sentinelErrorWrap := fmt.Errorf("in repository layer: %w", ErrCustom)
	// fmt.Println("sentinelErrorWrap:", sentinelErrorWrap)
	// fmt.Println("sentinelErrorWrap Type:", reflect.TypeOf(sentinelErrorWrap))

	// sentinelErrorWrap = fmt.Errorf("in service layer: %w", sentinelErrorWrap)
	// fmt.Println("sentinelErrorWrap:", sentinelErrorWrap)
	// fmt.Println("sentinelErrorWrap Type:", reflect.TypeOf(sentinelErrorWrap))

	// fmt.Println("error matched?:", errors.Is(sentinelErrorWrap, ErrCustom))

	file := "dummy.txt"
	fileCheckerError := fileChecker(file)
	if fileCheckerError != nil {
		if errors.Is(fileCheckerError, os.ErrNotExist) {
			fmt.Println("%v file not found\n", file)
		} else {
			fmt.Println("unknown error")
		}
	}
}

func fileChecker(name string) error {
	f, err := os.Open(name)
	if err != nil {
		return fmt.Errorf("in checker: %w", err)
	}
	defer f.Close()
	return nil
}
