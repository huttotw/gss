package gss

func ToString(in <-chan []byte) <-chan string {
	out := make(chan string)
	go func() {
		for item := range in {
			out <- string(item)
		}
		close(out)
	}()
	return out
}
