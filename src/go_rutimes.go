
import (
	"fmt"
	"sync"
)
func say(text string,wg *sync.WaitGroup){
	defer wg.Done()
	fmt.Println(text)
}

func goRutime(){
	var wg = sync.WaitGroup()
	fmt.Println("start...")
	wg.Add(1)

	// se agrega la go rutime y la direccion en memoria de hilo del proceso principal
	go say("ahhhh nu ma ci es cierto",&wg)

	wg.Wait()
}