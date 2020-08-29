package main

import (
	"fmt"
	"log"
	"os/exec"

	"golang.org/x/text/encoding/japanese"
)

func main() {
	cmd := exec.Command("my-app")
	b, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	// my-appからの出力はCP932であると限定し、UTF-8へ変換する
	b, err = japanese.ShiftJIS.NewDecoder().Bytes(b)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(b))
}
