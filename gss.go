package gss

import (
	"bufio"
	"io"
)

type Stream struct {
	r *bufio.Reader
}

func NewStream(r *bufio.Reader) *Stream {
	s := Stream{
		r: r,
	}
	return &s
}

func (s *Stream) Start() <-chan []byte {
	out := make(chan []byte)
	go func() {
		for {
			b, _, err := s.r.ReadLine()
			if err != nil && err != io.EOF {
				panic(err)
			} else if err == io.EOF {
				break
			}

			out <- b
		}
		close(out)
	}()
	return out
}
