package gss

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"strconv"
)

func CSVFieldAverage(in <-chan []byte, field int) <-chan float64 {
	out := make(chan float64)
	var total float64
	var count float64
	go func() {
		for item := range in {
			r := csv.NewReader(bytes.NewBuffer(item))
			record, err := r.Read()
			if err != nil {
				fmt.Printf("csv field avg: row %f: %s\n", count, err.Error())
				continue
			}
			if len(record) <= field {
				continue
			}
			val, err := strconv.ParseFloat(record[field], 64)
			if err != nil {
				fmt.Printf("csv field avg: row %f: %s\n", count, err.Error())
				continue
			}
			total += val
			count++
			out <- total / count
		}
		close(out)
	}()
	return out
}
