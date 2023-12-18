package main

import (
	"DemonstrationServiceL0/internal/database"
	"DemonstrationServiceL0/internal/transport/rest"
	"fmt"
)

func main() {
	fmt.Println("Connected to ...!")

	rest.StartServer(":8080")

	//Подключение к базе данных
	db := database.ConnectDB()
	defer db.Close()

}
