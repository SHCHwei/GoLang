package main


import (
    "fmt"
    "example.com/greetings"
    "log"
)


func main(){
    // Set properties of the predefined Logger, including
    // the log entry prefix and a flag to disable printing
    // the time, source file, and line number.
    log.SetPrefix("greetings:")
    log.SetFlags(0)

    // 宣告陣列 3個元素都是String
    arr := [3]string {"Gladys", ""}

    for index:= 0; index < 2 ; index++ {
        message, err := greetings.Hello(arr[index])
            if err != nil {
                log.Fatal("stop!")
            }
        	sayHello(message)
    }
}


func sayHello(enterName string ){
	fmt.Println(enterName)
}

