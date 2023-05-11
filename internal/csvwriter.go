package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	records := [][]string{
		{"first_name", "last_name", "username"},
		{"Rob", "Pike", "rob"},
		{"Ken", "Thompson", "ken"},
		{"Robert", "Griesemer", "gri"},
	}
	file, err := os.Create("hello.csv")

	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	w := csv.NewWriter(file)
	w.WriteAll(records) // вызывает Flush внутри

	if err := w.Error(); err != nil {
		log.Fatalln("error writing csv:", err)
	}

}
