package baz

import (
	"fmt"

	"github.com/binderclip/go-qux/v2/qux"
)

func Baz() {
	fmt.Println("Baz 003 >")
	qux.Qux()
	fmt.Println("Baz 003 <")
}
