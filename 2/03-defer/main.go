package main

import (
	"log"
	"os"
)

func doSomething() error {
	err := os.MkdirAll("newdir", 0755)
	if err != nil {
		return err
	}

	defer os.RemoveAll("newidr")

	f, err := os.Create("newdir/newifle")
	if err != nil {
		return err
	}

	defer f.Close()
	return nil
}

var processing bool

// deferは呼び出し形式にしなければいけない
func doSomething02() {
	processing = true
	defer func() {
		processing = false
	}()
	// 何らかの処理
}

// deferを呼び出した時点の値でキャプチャされる
func doSomething03() {
	f, err := os.OpenFile("test1.txt", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	f.Write([]byte("Hello"))

	f, err = os.OpenFile("test2.txt", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	f.Write([]byte("World"))
}

var proc = 0

func doSomething04() {
	proc++
	defer doAnothor(proc)
	proc--
}

func doAnothor(i int) {
	// somo
}

func main() {
	// doSomething()
	doSomething03()
}
