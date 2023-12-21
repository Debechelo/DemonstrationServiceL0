package service

import (
	"encoding/json"
	"log"
	"os"
)

const filePath = "model.json"

// Чтение содержимого файла model.json
func ReadJSON() []byte {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	defer file.Close()

	// Получение информации о файле
	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatal("Error getting file info:", err)
	}

	fileContent := make([]byte, fileInfo.Size())

	// Считывание содержимого файла в буфер
	_, err = file.Read(fileContent)
	if err != nil {
		log.Fatal("Error reading file content:", err)
	}
	return fileContent
}

// Декодирование JSON
func DecodeJSON(content []byte) *Order {
	var order *Order
	if err := json.Unmarshal(content, &order); err != nil {
		log.Fatal("Error decoding JSON:", err)
	}
	return order
}

// Раскодирование JSON
func CodeJSON(order Order) []byte {
	jsonBytes, err := json.Marshal(order)
	if err != nil {
		log.Fatal("Error marshalling JSON:", err)
	}
	return jsonBytes
}
