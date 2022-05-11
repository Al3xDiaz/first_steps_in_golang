package main

import (
	"fmt"

	"github.com/al3xdiaz/first_steps_in_golang/src"
)

func message(text string, c chan string) {
	c <- text
}

func main() {

	sala1 := make(chan string)
	sala2 := make(chan string)

	go message("sala 1: bienvenidos ...", sala1)
	go message("sala 2: bienvenidas ...", sala2)

	for i := 0; i < 2; i++ {
		select {
		case m1 := <-sala1:
			fmt.Println(m1)
		case m2 := <-sala2:
			fmt.Println(m2)
		}
	}
	src.GoRutime()

}
