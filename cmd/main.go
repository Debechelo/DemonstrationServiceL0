package main

import (
	"DemonstrationServiceL0/internal/transport/rest"
	"fmt"
)

func main() {
	fmt.Println("Connected to ...!")

	rest.StartServer(":8080")
}
