package main

import "auth-service/internal/server"
import "fmt"

func main() {
	fmt.Println("Starting the server...")
	server.Run()
	fmt.Println("Server has started.") 
}
