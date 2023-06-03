package main

import (
	//"encoding/csv"
	"encoding/json"
	//"fmt"
	"io/ioutil"
	//"log"
	//"os"
	//"strconv"
)

type rawStorage struct {
	path       string
	fileFormat string
	totalSize  float64
}

/*
func readCsvFile(filePath string) []customer {
	recordsCustomer := []customer{}
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
		recordsCustomer = append(recordsCustomer, customer{age_int, data[1], 7, customerAddress{"", "", ""}})
		fmt.Println(recordsCustomer)
	}
	return recordsCustomer
}
*/

func (r rawStorage) writer(filename string, dataCustomer string) {

	file, _ := json.MarshalIndent(dataCustomer, "", " ")
	filePath := r.path + filename
	_ = ioutil.WriteFile(filePath, file, 0644)

}

/*
func (r rawStorage) reader(fileNames string) customer {
	filePath := r.path + fileNames
	fmt.Println(fileNames)
	customerJson, err := os.Open(filePath)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	defer customerJson.Close()

	byteJson, _ := ioutil.ReadAll(customerJson)

	var records map[string]interface{} // -> if we dont know the structure of the json
	err = json.Unmarshal(byteJson, &records)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}
	//build_str :=
	for key, element := range records {
		fmt.Println("Key:", key, "=>", "Element:", element)
	}

	nestedMap := records["address"].(map[string]interface{})
	customerReturn := customer{1, records["lastName"].(string), int(records["age"].(float64)), customerAddress{nestedMap["streetAddress"].(string), nestedMap["city"].(string), nestedMap["state"].(string)}}

	return customerReturn
}
*/
