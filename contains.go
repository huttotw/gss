package gss

import (
	"bytes"
)

func Contains(in <-chan []byte, needle []byte) <-chan []byte {
	out := make(chan []byte)
	go func() {
		for item := range in {
			if bytes.Contains(item, needle) {
				out <- item
			}
		}
		close(out)
	}()
	return out
}
