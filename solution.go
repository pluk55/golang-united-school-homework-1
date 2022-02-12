package solutions

import (
	"fmt"
	"github.com/kyokomi/emoji/v2"
)

func GetMessage() string {
	msg := fmt.Sprintf("Hello :world-map:")
	return emoji.Sprint(msg)
}
