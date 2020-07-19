package main

import (
	//"bufio"
	"encoding/csv"
	"flag"
	"io"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator
	filePath := flag.String("filePath", "", "path to csv")
	flag.Parse()
	// Open the file
	csvfile, err := os.Open(*filePath)
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	// Parse the file
	r := csv.NewReader(csvfile)
	r.LazyQuotes = true

	// Iterate through the records
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		// TODO: load record to DB
		LoadRecordToDatabase(record)
		//fmt.Printf("id: %s val %s\n", record[0], record[1])
	}
}
