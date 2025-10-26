package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Stats struct {
	sum   float64
	min   float64
	max   float64
	count int
}

func main() {
	file, _ := os.Open("data.txt")
	defer file.Close()

	var mu sync.Mutex
	stats := make(map[string]Stats)

	start := time.Now()

	scanner := bufio.NewScanner(file)
	lines := make(chan string, 1000)

	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go processData(&wg, stats, lines, &mu)
	}

	for scanner.Scan() {
		lines <- scanner.Text()
	}
	close(lines)
	wg.Wait()

	for city, value := range stats {
		avg := value.sum / float64(value.count)
		fmt.Printf("%s=%.1f/%.1f/%.1f \n", city, value.min, avg, value.max)
	}

	fmt.Print(time.Since(start))
}

func processData(wg *sync.WaitGroup, stats map[string]Stats, lines chan string, mu *sync.Mutex) {
	defer wg.Done()

	for line := range lines {
		if len(line) > 0 {
			data := strings.Split(line, ";")
			if len(data) < 2 {
				continue
			}

			n, err := strconv.ParseFloat(data[1], 64)
			if err != nil {
				continue
			}

			mu.Lock()
			location, ok := stats[data[0]]
			if !ok {
				stats[data[0]] = Stats{
					sum: n,
					min: n, max: n,
					count: 1,
				}

			} else {
				stats[data[0]] = Stats{
					sum:   location.sum + n,
					min:   min(n, location.min),
					max:   max(n, location.max),
					count: location.count + 1,
				}
			}
			mu.Unlock()
		}

	}
}
