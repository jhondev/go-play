package parser

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"os"
	"rain-csv-parser/mapper"
	"strings"
)

func Parse(path string, mp *mapper.Mapper) error {
	log.Printf("Reading file '%s'.\n", path)
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	data, err := csvReader.ReadAll()
	if err != nil {
		return err
	}
	log.Print("Looking for a matching configuration map\n")
	name, cells := mp.FindMap(data[0])
	if cells == nil {
		return errors.New("configuration map for current file structure not found")
	}
	log.Printf("Configuration map found: %s.\n", name)
	log.Println("Creating result file.")
	rfile, err := os.Create(strings.ReplaceAll(path, ".csv", ".result.csv"))
	if err != nil {
		return err
	}
	defer rfile.Close()

	header := cells[0].Name
	for i := 1; i < len(cells); i++ {
		header = fmt.Sprintf("%s,%s", header, cells[i].Name)
	}
	rfile.WriteString(header + "\n")

	for ida := 1; ida < len(data); ida++ {
		row := cellValue(data[ida], &cells[0])
		for ice := 1; ice < len(cells); ice++ {
			row = fmt.Sprintf("%s,%s", row, cellValue(data[ida], &cells[ice]))
		}
		rfile.WriteString(row + "\n")
	}
	log.Println("Result file created successfully.")
	fmt.Println("")

	return nil
}

func cellValue(drow []string, cell *mapper.Cell) string {
	val := drow[cell.Indexes[0]]
	for i := 1; i < len(cell.Indexes); i++ {
		val = fmt.Sprintf("%s %s", val, drow[cell.Indexes[i]])
	}
	return val
}
