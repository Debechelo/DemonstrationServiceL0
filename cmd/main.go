package main

import (
	"fmt"

	"DemonstrationServiceL0/internal/handler"
)

func main() {
	fmt.Println("Connected to ...!")

	handler.StartServer(":8080")
}
