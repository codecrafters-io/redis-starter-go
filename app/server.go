package main

import (
	"fmt"
)

func main() {
	fmt.Println("boot your server here!")

	// Comment out the following to pass the first stage
	//
	// l, err := net.Listen("tcp", "0.0.0.0:6379")
	// if err != nil {
	// 	fmt.Println("Failed to bind to port 6379")
	// 	os.Exit(1)
	// }
	// _, err = l.Accept()
	// if err != nil {
	// 	fmt.Println("Error accepting connection: ", err.Error())
	// 	os.Exit(1)
	// }
}
