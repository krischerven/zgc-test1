package main

import (
	"fmt"
	"math"
	"github.com/tj/go-spin"
	lru "github.com/krischerven/zgc-test1/gotest-1/src/lru_cache/fast"
	"runtime"
	"time"
)

const (
	lruCacheMillionsOfItems = 20
	gcIterations            = 20
)

func main() {
	newInt := func(i int) *int {
		return &i
	}
	heap := func(run int) string {
		truncate := func(s string) string {
			if len(s) > 5 {
				return s[0:4]
			} else {
				return s
			}
		}
		var super string
		if run == 0 {
			super = "⁰"
		} else if run == 1 {
			super = "¹"
		}
		var ms runtime.MemStats
		runtime.ReadMemStats(&ms)
		return "Heap" + super + ": " + truncate(
			fmt.Sprintf("%f", float64(ms.Alloc)/(1000*1000*1000))) + " GB"
	}
	print2 := func(x interface{}) {
		if x == nil {
			fmt.Println()
		} else {
			fmt.Println(x)
		}
	}
	print2("Original " + heap(-1))
	allocated := false
	go func() {
		s := spin.New()
		secs := float64(0)
		for !allocated {
  			fmt.Printf("\r\033 [Allocating the LRU cache\033[m %s (%.1f)", s.Next(), secs)
			time.Sleep(time.Millisecond*100)
			secs += 0.1
		}
	} ()
	c := lru.New(1000 * 1000 * lruCacheMillionsOfItems)
	for i := 0; i < 1000*1000*lruCacheMillionsOfItems; i++ {
		c.Refer(newInt(i))
	}
	allocated = !allocated
	print2(nil)
	print2("Finished allocating the LRU cache.")
	latency := struct{min, max, mean, c int64} {0, 0, 0, 0}
	for i := 0; i < gcIterations; i++ {
		print2(nil)
		print2(heap(0))
		go func() {
			t0 := time.Now()
			time.Sleep(time.Millisecond * 10)
			latency_ := time.Now().Sub(t0).Microseconds()-10000
			print2(fmt.Sprintf("Latency: %d µs", latency_))
			if latency.min == 0 || latency_ < latency.min {
				latency.min = latency_
			}
			latency.max = int64(math.Max(float64(latency.max), float64(latency_)))
			latency.mean += latency_
			latency.c++
		}()
		t1 := time.Now()
		runtime.GC()
		t2 := time.Now().Sub(t1)
		print2(heap(1))
		print2("Time to perform a full GC: " + fmt.Sprintf("%d", t2.Milliseconds()) + " ms")
		print2("Sleeping for " + fmt.Sprintf("%d", t2.Milliseconds()/2) + " ms");
		time.Sleep(time.Millisecond*time.Duration(t2.Milliseconds()/2))
	}
	// force memory to stay alive
	println(c.Size())
	// latency stats
	latency.mean /= latency.c
	print2(fmt.Sprintf("Latency (min, max, mean): %d µs, %d µs, %d µs", latency.min, latency.max, latency.mean))
}
