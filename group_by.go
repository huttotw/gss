package gss

func GroupByString(in <-chan string) <-chan map[string]int {
	out := make(chan map[string]int)
	var m = make(map[string]int)
	go func() {
		for item := range in {
			if _, ok := m[item]; !ok {
				m[item] = 1
			} else {
				m[item]++
			}
			out <- m
		}
		close(out)
	}()
	return out
}
