package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"

	//"io/ioutil"
	"log"
	"os"
	"path"
	"strconv"
)

type data_reader interface {
	reader()
}

type customerAddress struct {
	street string
	city   string
	state  string
}

type customer struct {
	id      int
	name    string
	age     int
	address customerAddress
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
		records_customer = append(records_customer, customer{age_int, data[1], 7, customerAddress{"", "", ""}})
		fmt.Println(records_customer)
	}
	return records_customer
}

func readJsonFile(filePath string) customer {

	/*content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}
	customerJson := `{"firstName": "Alex", "lastName": "Smith", "age": 36, "address": {"streetAddress": "Canal Street 3", "city": "New Orleans", "state": "Louisiana"}}`
	var result map[string]any

	// Now let's unmarshall the data into `payload`
	err := json.Unmarshal([]byte(customerJson), &records)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}*/
	customerJson := `{"firstName": "Alex", "lastName": "Smith", "age": 36, "address": {"streetAddress": "Canal Street 3", "city": "New Orleans", "state": "Louisiana"}}`
	//var customerReturn customer
	//birdJson := `{"birds":{"pigeon":"likes to perch on rocks","eagle":"bird of prey"},"animals":"none"}`
	var records map[string]interface{} // -> if we dont know the structure of the json
	err := json.Unmarshal([]byte(customerJson), &records)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}
	customerReturn := customer{1, records["firstName"].(string) + records["lastName"].(string), records["age"].(int), customerAddress{records["streetAddress"].(string), records["city"].(string), records["state"].(string)}}

	fmt.Println(records["firstName"])

	return customerReturn
}
