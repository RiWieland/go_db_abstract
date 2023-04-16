package main

import (
	"encoding/csv"
	//"fmt"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path"
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

	fileExtension := path.Ext(filePath)
	if fileExtension == ".csv" {

	} else if fileExtension == ".json" {

	}

}

func readCsvFile(filePath string) customer {
	var records customer
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	fileExtension := path.Ext(filePath)
	if fileExtension == ".csv" {
		csvReader := csv.NewReader(f)
		records, err := csvReader.ReadAll()
		if err != nil {
			log.Fatal("Unable to parse file as CSV for "+filePath, err)
		}

		return records
	}
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
