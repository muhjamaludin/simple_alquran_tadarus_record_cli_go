package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
)

type Record struct {
	Id      int
	Juz     string
	Surah   string
	Ayat    string
	Tanggal string
	Jam     string
}

func main() {
	fmt.Println("Program Pencatatan Tadarus Alquran")
	fmt.Printf("==================================\n\n")
	var juz, surah, ayat, tanggal, jam string

	for {
		fmt.Printf("Daftar Menu: \n\n")

		fmt.Println("1. Records")
		fmt.Println("2. Add Record")
		fmt.Printf("0. Keluar\n\n")

		// read from input cli
		fmt.Printf("Your choice: ")
		inp := scanInputCli()
		fmt.Printf("\n")

		// show list
		records := readFileJson()
		if inp == "1" {
			// setup table
			t := table.NewWriter()
			t.SetOutputMirror(os.Stdout)
			t.AppendHeader(table.Row{"#", "Juz", "Surah", "Ayat", "Tanggal", "Jam"})
			for i := range records {
				t.AppendRows([]table.Row{{records[i].Id, records[i].Juz, records[i].Surah, records[i].Ayat, records[i].Tanggal, records[i].Jam}})
			}
			t.Render()
			fmt.Printf("\n")
		}

		// add record
		if inp == "2" {
			var tmpRecord Record
			fmt.Println("Tambah Catatan")

			id := records[0].Id + 1
			tmpRecord.Id = id

			fmt.Printf("Juz: ")
			juz = scanInputCli()
			tmpRecord.Juz = juz

			fmt.Printf("Surah: ")
			surah = scanInputCli()
			tmpRecord.Surah = surah

			fmt.Printf("Ayat: ")
			ayat = scanInputCli()
			tmpRecord.Ayat = ayat

			fmt.Printf("Tanggal: ")
			tanggal = scanInputCli()
			tmpRecord.Tanggal = tanggal

			fmt.Printf("Jam: ")
			jam = scanInputCli()
			tmpRecord.Jam = jam

			records = append([]Record{tmpRecord}, records...)
			file, err := json.MarshalIndent(records, "", "  ")
			if err != nil {
				log.Fatal(err)
			}
			err = os.WriteFile("data.json", file, 0644)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("\n\n")
		}

		// exit
		if inp == "0" {
			break
		}
	}
}

func readFileJson() []Record {
	var record []Record
	// no 1.
	data_json, err := os.ReadFile("data.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(data_json, &record)
	if err != nil {
		log.Fatal(err)
	}
	return record
}

func scanInputCli() string {
	reader := bufio.NewScanner(os.Stdin)
	reader.Scan()
	inp := reader.Text()
	err := reader.Err()
	if err != nil {
		log.Fatal(err)
	}
	return inp
}
