package baz

import (
	"fmt"

	"github.com/binderclip/go-qux/qux"
)

func Baz() {
	fmt.Println("Baz 001 >")
	qux.Qux()
	fmt.Println("Baz 001 <")
}
