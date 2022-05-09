import (
	"fmt"
)
func say(text string,c chan<- string){
	c <- text
}

func main(){
	c:= make(chan string, 1)
	fmt.Println("start process...")
	go say("start go rutime with chanels",c)

	fmt.Println(<-c)
}