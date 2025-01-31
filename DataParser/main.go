// package main

// import (
// 	"encoding/csv"
// 	"encoding/json"
// 	"fmt"
// 	"os"
// )

// func parseCSV(filePath string) {
// 	file, _ := os.Open(filePath)
// 	defer file.Close()
// 	reader := csv.NewReader(file)
// 	data, _ := reader.ReadAll()
// 	fmt.Println(data)
// }

// func parseJSON(filePath string) {
// 	file, _ := os.Open(filePath)
// 	defer file.Close()
// 	var data interface{}
// 	json.NewDecoder(file).Decode(&data)
// 	fmt.Println(data)
// }

// func main() {
// 	parseCSV("data.csv")
// 	parseJSON("data.json")
// }

// The code is mostly correct, but it lacks proper error handling. It's important to handle errors to avoid unexpected crashes and to provide meaningful error messages.

package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
)

func parseCSV(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening CSV file: %v\n", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	data, err := reader.ReadAll()
	if err != nil {
		fmt.Printf("Error reading CSV file: %v\n", err)
		return
	}
	fmt.Println(data)
}

func parseJSON(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening JSON file: %v\n", err)
		return
	}
	defer file.Close()

	var data interface{}
	if err := json.NewDecoder(file).Decode(&data); err != nil {
		fmt.Printf("Error decoding JSON file: %v\n", err)
		return
	}
	fmt.Println(data)
}

func main() {
	parseCSV("data.csv")
	parseJSON("data.json")
}
