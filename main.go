package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type CsvLine struct {
	history_id string
	machine_group_id string
	MacNOtmh string
	MacNOtm string
	customer_id string
	MacSTTtmh string
	MacSTTtm string
	action_amount string
	action_amount_unsign string
	price string
	machine_price string
	action_by string
	starttime string
}

var rec [][]string

func main() {
	lines, err := ReadCsv("amount5.csv")
	if err != nil {
		panic(err)
	}

	// Loop through lines & turn into object
	for _, line := range lines {
		data := CsvLine{
			history_id: line[0],
			machine_group_id:line[1],
			MacNOtmh:line[2],
			MacNOtm:line[3],
			customer_id :line[4],
			MacSTTtmh :line[5],
			MacSTTtm :line[6],
			action_amount :line[7],
			action_amount_unsign: line[8],
			price: fill_price(line[8], line[9]),
			machine_price: line[9],
			action_by: line[10],
			starttime: line[11],
		}
		/*
		fmt.Println(data.history_id + " " +
			data.machine_group_id + " " +
			data.MacNOtmh + " " +
			data.MacNOtm + " " +
			data.customer_id + " " +
			data.MacSTTtmh + " " +
			data.MacSTTtm + " " +
			data.action_amount + " " +
			data.action_amount_unsign + " " +
			data.price + " " +
			data.machine_price + " " +
			data.action_by + " " +
			data.starttime)*/
		//fmt.Println(data.action_amount_unsign + " " + data.price + " " + data.machine_price)
		//fmt.Println(data.action_amount_unsign[0],"Hello GO")
		fmt.Println(data)

		/*temp := [][]string{{data.history_id + "," +
			data.machine_group_id + "," +
			data.MacNOtmh + "," +
			data.MacNOtm + "," +
			data.customer_id + "," +
			data.MacSTTtmh + "," +
			data.MacSTTtm + "," +
			data.action_amount + "," +
			data.action_amount_unsign + "," +
			data.price + "," +
			data.machine_price + "," +
			data.action_by + "," +
			data.starttime},
		}*/
		//rec = append(rec, temp)
	}
	//WriteCsv()
}

// ReadCsv accepts a file and returns its content as a multi-dimentional type
// with lines and each column. Only parses to string type.
func ReadCsv(filename string) ([][]string, error) {

	// Open CSV file
	f, err := os.Open(filename)
	if err != nil {
		return [][]string{}, err
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)

	// Read File into a Variable
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return lines, nil
}

func WriteCsv() {
	records := [][]string{
		{"first_name", "last_name", "occupation"},
		{"John", "Doe", "gardener"},
		{"Lucy", "Smith", "teacher"},
		{"Brian", "Bethamy", "programmer"},
	}

	f, err := os.Create("users.csv")
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)

	if err != nil {

		log.Fatalln("failed to open file", err)
	}

	w := csv.NewWriter(f)
	err = w.WriteAll(records) // calls Flush internally

	if err != nil {
		log.Fatal(err)
	}
}

func fill_price(aau string, mp string) string {
	var price string
	if aau == "0.0" {
		price = mp
	} else {
		price = aau
	}
	return price
}