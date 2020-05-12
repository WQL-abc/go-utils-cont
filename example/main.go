package main

import (
	"fmt"
	"time"

	"github.com/go-utils/cont"
)

func main() {
	var (
		now = time.Now()
		ex1 = []string{"1", "2", "3"}
		ex2 = []int{1, 2, 3}
		ex3 = "123"
		ex4 = map[string]interface{}{"a": "1", "b": 2, "c": "3", "d": now}
	)

	fmt.Printf("%#v\n", ex1)
	fmt.Printf("%#v\n", ex2)
	fmt.Printf("%#v\n", ex3)
	fmt.Printf("%#v\n", ex4)

	if cont.Contains(ex1, "1") {
		fmt.Println("ex1: OK!")
	}

	if cont.Contains(ex2, 1) {
		fmt.Println("ex2: OK!")
	}

	if cont.Contains(ex3, "1") {
		fmt.Println("ex3: OK!")
	}

	if cont.Contains(ex4, now) {
		fmt.Println("ex4: OK! - time.Time")
	}

	if cont.Contains(ex4, 2) {
		fmt.Println("ex4: OK! - int")
	}

	if cont.Contains(ex4, "c") {
		fmt.Println("ex4: OK! - string")
	}
}
