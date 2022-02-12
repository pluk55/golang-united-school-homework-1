package main

import (
	"fmt"
	"github.com/kyokomi/emoji/v2"
)

func GetMessage() string {
	fmt.Println("Hello World Emoji!")

	emoji.Sprint(":beer: Beer!!!")
	return emoji.Sprint("Hello :beer:")
}

func main() {
	fmt.Println(GetMessage())
}
