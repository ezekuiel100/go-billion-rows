package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Stats struct {
	sum   float64
	min   float64
	max   float64
	count int
}

func main() {
	stats := make(map[string]Stats)

	start := time.Now()

	file, err := os.Open("data.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	r := bufio.NewReader(file)

	for {
		line, _, err := r.ReadLine()
		if len(line) > 0 {
			data := strings.Split(string(line), ";")
			location, ok := stats[data[0]]
			n, _ := strconv.ParseFloat(data[1], 64)

			if !ok {
				stats[data[0]] = Stats{sum: n, min: n, max: n, count: 1}
			} else {
				stats[data[0]] = Stats{sum: location.sum + n, min: min(n, location.min), max: max(n, location.max), count: location.count + 1}
			}
		}

		if err != nil {
			break
		}
	}

	for city, value := range stats {
		avg := value.sum / float64(value.count)
		fmt.Printf("%s=%.1f/%.1f/%.1f \n", city, value.min, avg, value.max)
	}

	fmt.Print(time.Since(start))
}
