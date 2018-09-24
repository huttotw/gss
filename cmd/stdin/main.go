package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/huttotw/gss"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	stream := gss.NewStream(r)

	out := stream.Start()
	strings := gss.ToString(out)
	groups := gss.GroupByString(strings)

	for group := range groups {
		fmt.Println("-- result --")
		fmt.Println(group)
	}
}
