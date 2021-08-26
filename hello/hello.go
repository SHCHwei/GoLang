package main


import (
    "fmt"
    "example.com/greetings"
    "log"
)


func main(){

	message := greetings.Hello("Gladys")
	fmt.Println(message)
	sayHello()

}

func sayHello(){
	fmt.Println("Hello master")
}

