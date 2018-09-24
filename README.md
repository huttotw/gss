# Go Structured Streaming

Go Structured Streaming is a very basic derivitive of Spark Structured streaming, mostly relying on its backing principles to make effiecent pipelines. At a high level, this package provides a wrapper around Go pipelines, with the goal of being able to define pipelines in some configuration file.

# Example
```go
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
```

This example will find all MLB players named Adam, then print their average age. If you run this program, the average will output for every Adam we encounter. The last average is our final calculation.

Please see a couple of example implementations: [csv](cmd/csv/main.go), [stdin](cmd/stdin/main.go).

# Todo
- [ ] Add more tranforming pipelines, JSON, CSV, etc.
- [ ] Add more aggregating pipelines, min, max, etc.
- [ ] Add more filtering pipelines, in, not in, etc.
- [ ] Provide a way to load and save a pipeline, maybe YAML, JSON?
- [ ] Provide prebuilt sources, NSQ, Kafka, TCP, etc.

# License
Copyright © 2018 Trevor Hutto

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this work except in compliance with the License. You may obtain a copy of the License in the LICENSE file, or at:

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.