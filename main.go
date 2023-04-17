package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strconv"
)

type data_reader interface {
	reader()
}

type customer struct {
	id          int
	name        string
	age         int
	email       string
	adress      string
	phonenumber string
}

func main() {
	filePath := "data_csv/customer_20230415.csv"
	fileExtension := path.Ext(filePath)
	if fileExtension == ".csv" {
		extract := readCsvFile(filePath)
		fmt.Println(extract)
	} else if fileExtension == ".json" {
		fmt.Printf("not implemented yet")

	}

}

func readCsvFile(filePath string) []customer {
	records_customer := []customer{}
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	data_extract, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}
	fmt.Println(data_extract)
	_, data_content := data_extract[0], data_extract[1:]
	for _, data := range data_content {
		fmt.Println(data[0])
		age_int, err := strconv.Atoi(data[0])
		if err != nil {
			log.Fatal("can not convert to int: ", age_int, err)
		}
		records_customer = append(records_customer, customer{age_int, data[1], 7, data[2], "", ""})
		fmt.Println(records_customer)
	}
	return records_customer
}

func readJsonFile(filePath string) customer {
	var records customer
	content, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	// Now let's unmarshall the data into `payload`
	err = json.Unmarshal(content, records)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	return records
}
