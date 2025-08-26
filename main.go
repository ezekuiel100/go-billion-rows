package main

import (
	"bufio"
	"fmt"
	"math"
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
			value, _ := stats[data[0]]

			n, _ := strconv.ParseFloat(data[1], 64)
			stats[data[0]] = Stats{sum: value.sum + n, min: min(n, value.min), max: max(n, value.max), count: value.count + 1}
		}

		if err != nil {
			break
		}
	}

	for city, value := range stats {
		avg := math.Round(value.sum/float64(value.count)*100) / 100
		fmt.Println(city, "Avg=", avg, "min=", value.min, "max=", value.max)
	}

	finish := time.Now()
	duration := finish.Sub(start)
	fmt.Print(duration.Seconds())
}
