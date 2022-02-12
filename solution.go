package main

import (
	"fmt"
	"github.com/kyokomi/emoji/v2"
)

func GetMessage(emj string) string {
	msg := fmt.Sprintf("Hello :%s:", emj)
	return emoji.Sprint(msg)
}

func main() {
	fmt.Println(GetMessage("beer"))
}
