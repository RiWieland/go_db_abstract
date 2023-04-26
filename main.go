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

type customer_address struct {
	street string
	city   string
	state  string
}

type customer struct {
	id      int
	name    string
	age     int
	address customer_address
}

func main() {
	filePath := "data_csv/customer_20230415.csv"
	filePathJson := "data_json/customer_20230412.json"
	fileExtension := path.Ext(filePathJson)
	if fileExtension == ".csv" {
		//extract := readCsvFile(filePath)
		extract_json := readJsonFile(filePath)
		fmt.Println(extract_json)
	} else if fileExtension == ".json" {
		extract_json := readJsonFile(filePath)
		fmt.Println(extract_json)
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
		records_customer = append(records_customer, customer{age_int, data[1], 7, customer_address{"", "", ""}})
		fmt.Println(records_customer)
	}
	return records_customer
}

func readJsonFile(filePath string) customer {
	var records customer
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	// Now let's unmarshall the data into `payload`
	err = json.Unmarshal(content, &records)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	return records
}
