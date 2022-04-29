package main

import (
    "encoding/csv"
    "fmt"
    "log"
    "os"
	"io"
	"strconv"
)

type Person struct {
    Id int
	Name string
    Age int
}

func main() {
	
	//open csv file
	objfile, err := os.Open("csvdata.csv")
	
	//handle error, if file not found
    if err != nil {
        log.Fatal(err)
    }
	
	//array to hold person data  
	var Persons []Person
	
	// read csv data using csv.ReadAll
    csvReader := csv.NewReader( objfile)
    
	for {
        fields, err := csvReader.Read()
		
		//if reached end of file, break the loop
        if err == io.EOF {
            break
        }
		
		//if encounted with error exit
        if err != nil {
            log.Fatal(err)
        }
		
		//person object to hold the translated csv line
		var objPerson Person
		
		//loop through all csv fields
		for fieldindex, field := range fields {
			switch fieldindex {
				case 0:
					objPerson.Id, _ = strconv.Atoi( field)
					break
				case 1:
					objPerson.Name = field
					break
				case 2:
					objPerson.Age, _ = strconv.Atoi( field)
					break
			}
		}
		
		//append to array
		Persons = append(Persons, objPerson)
		
    }
	
	// close the file, we are done with reading
    objfile.Close()
	
	fmt.Printf("%+v\n", Persons)
}

