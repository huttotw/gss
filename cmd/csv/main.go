package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/huttotw/gss"
)

func main() {
	f, err := os.Open("mlb-players.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	r := bufio.NewReader(f)
	stream := gss.NewStream(r)

	out := stream.Start()
	adams := gss.Contains(out, []byte("Adam"))
	avgAge := gss.CSVFieldAverage(adams, 5)

	for avg := range avgAge {
		fmt.Println(avg)
	}
}
