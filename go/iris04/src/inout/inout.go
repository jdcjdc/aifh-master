package inout

import (
	"os"
	"encoding/csv"
	"io"
	"fmt"
	"bufio"
	"log"
)

func ReadFileLines (file02 string) (int, []string) {
	file, err := os.Open(file02)
	lines := make([]string, 0, 1000)  // lines is a slice

	if err != nil {
		// todo later return nil, err
		fmt.Printf("err has type %T\n", err)
		if err2, ok := err.(*os.PathError); ok {
			fmt.Printf("err2 has type %T\n", err2.Error)
		}
		return 0, nil
	}
	defer file.Close()  // this needs to be after the err check

	scanner := bufio.NewScanner(file)
	var nbrOfLines int
	for scanner.Scan() {
		var line = scanner.Text()
		// fmt.Println(scanner.Text())
		fmt.Println(line)
		lines = append(lines, line)
		nbrOfLines += 1

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return nbrOfLines, lines
}


// based on http://stackoverflow.com/questions/24999079/reading-csv-file-in-go
func ReadCsv(file string) {
	f, err := os.Open(file)
	if err != nil {
		// todo later return nil, err
		fmt.Printf("err has type %T\n", err)
		if err2, ok := err.(*os.PathError); ok {
			fmt.Printf("err2 has type %T\n", err2.Error)
		}
		return
	}
	defer f.Close()  // this needs to be after the err check

	csvr := csv.NewReader(f)

	for {
		row, err := csvr.Read()
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			return
		}
		fmt.Printf("%d: %s\n", row)
	}
}
