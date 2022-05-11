package src

import (
	"fmt"
	"sync"
)

/*
using chanels
func sayChanel(text string, c chan<- string) {
	c <- text
}

func StartRutime() {
	/*
		create chanel
		params:
			1: type chanel
			2: concurrent tasks numbers (optional)
	*\/
	c := make(chan string, 1)
	fmt.Println("start process...")
	go sayChanel("start go rutime with chanels", c)

	fmt.Println(<-c)
}
*/

func say(text string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(text)
}

func GoRutime() {
	var wg sync.WaitGroup
	fmt.Println("start...")
	wg.Add(1)

	// se agrega la go rutime y la direccion en memoria de hilo del proceso wg
	go say("ahhhh nu ma ci es cierto", &wg)
	fmt.Println("stop service")

	// se espera a que el proceso wg termine
	wg.Wait()
}
