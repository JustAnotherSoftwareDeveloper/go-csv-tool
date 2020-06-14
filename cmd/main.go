package main
import (
	"time"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)
func main() {
	operations := readCSV()
	printTotalUsers(operations)
	printLargerThan50Kb(operations)
	printJeffTwentyTwoOnDay(operations)
}


type ServerOperation struct {
	username string 
	operation string
	size int64 
	timestamp time.Time 
}
func printJeffTwentyTwoOnDay(operations []ServerOperation) {
	uploads := 0

	for _, operation := range operations {
		correctDate := operation.timestamp.Year() == 2020 && operation.timestamp.Month() == 4 && operation.timestamp.Day() == 15
		if operation.operation == "upload" && operation.username == "jeff22" && correctDate {
			uploads += 1
		}
	}

	fmt.Println("Times jeff22 Uploaded on April 15th, 2020:",uploads)
}
func printTotalUsers(operations []ServerOperation) {
	uniqueUsers := make(map[string]bool)

	for _, operation := range operations {
		uniqueUsers[operation.username] = true
	}

	fmt.Println("Unique Users:",len(uniqueUsers))
}

func printLargerThan50Kb(operations []ServerOperation) {
	uploads := 0

	for _, operation := range operations {
		if operation.operation == "upload" && operation.size > 50 {
			uploads += 1
		}
	}

	fmt.Println("Uploads Larger Than 50kb:",uploads)
}
func readCSV() []ServerOperation {
	csvfile, err := os.Open("server_log.csv")
	if err != nil {
		log.Fatalln("Could Not Open CSV File",err)
	}

	reader := csv.NewReader(csvfile)

	var operations []ServerOperation

	timeLayout := "Mon Jan 2 15:04:05 MST 2006"
	for {
		record , err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		timestamp,err := time.Parse(timeLayout,record[0])
		user := record[1]
		operation := record[2]
		size,err := strconv.ParseInt(record[3],10,64)
		operations = append(operations, ServerOperation{username: user, operation: operation, size: size, timestamp: timestamp})

	}

	return operations

}
