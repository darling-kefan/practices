package main

import (
	"fmt"
)

type T struct {
	a int
}

func (tv T) Mv(a int) int          { return 0 }
func (tp *T) Mp(f float32) float32 { return 1 }

var t T

func main() {
	fmt.Println(t.Mv(7), T.Mv(t, 7), (T).Mv(t, 7))

	f1 := T.Mv; a := f1(t, 7)
	fmt.Println(a)
}
