package main

import (
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Create("log.txt")
	if err != nil {
		// エラーログを出力した後に強制終了
		log.Fatalln(err)
	}
	defer file.Close()
	// log.Lshortfile ⇒ エラーが発生した行数をlogに追加するオプション
	flags := log.Lshortfile
	// log.New ⇒ ログを作成
	warnLogger := log.New(io.MultiWriter(file, os.Stderr), "WARN:", flags)
	errorLogger := log.New(io.MultiWriter(file, os.Stderr), "ERROR:", flags)

	warnLogger.Println("warning A")
	errorLogger.Fatalln("critical Error")
}
