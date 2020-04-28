package main

import (
	"fmt"
	"github.com/tj/go-spin"
	lru "lru_cache/simple"
	"runtime"
	"time"
)

const (
	lruCacheMillionsOfItems = 50
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
  			fmt.Printf("\r\033 [Allocating the LRU cache\033[m %s (%d)", s.Next(), int(secs))
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
	for i := 0; i < gcIterations; i++ {
		print2(nil)
		print2(heap(0))
		go func() {
			t0 := time.Now()
			time.Sleep(time.Millisecond * 10)
			print2(fmt.Sprintf("Latency: %d µs", time.Now().Sub(t0).Microseconds()-10000))
		}()
		t1 := time.Now()
		runtime.GC()
		t2 := time.Now().Sub(t1)
		print2(heap(1))
		print2("Time to perform a full GC: " + fmt.Sprintf("%d", t2.Milliseconds()) + " ms")
	}
	// force memory to stay alive
	println(c.Size())
}
