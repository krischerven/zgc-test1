package main

import (
	"fmt"
	gcf "github.com/krischerven/zgc-test1/gotest-1/src/gcf_cache"
	lru "github.com/krischerven/zgc-test1/gotest-1/src/lru_cache/fast"
	"github.com/tj/go-spin"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"math"
	"runtime"
	"time"
)

const (
	lruCacheMillionsOfItems = 20
	gcIterations            = 10
	gcf_cache               = false
)

func main() {
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
	done := func() {
		panic("Done")
	}
	_ = done
	print2("Original " + heap(-1))
	allocated := false
	var c interface {
		Size() int
		Name() string
	}
	if gcf_cache {
		c = gcf.NewPtr(1000 * 1000 * lruCacheMillionsOfItems)
	} else {
		c = lru.New(1000 * 1000 * lruCacheMillionsOfItems)
	}
	go func() {
		s := spin.New()
		secs := float64(0)
		for !allocated {
			fmt.Printf("\r\033 [Allocating the %s \033[m %s (%.1f)", c.Name(), s.Next(), secs)
			time.Sleep(time.Millisecond * 100)
			secs += 0.1
		}
	}()
	for i := 0; i < 1000*1000*lruCacheMillionsOfItems; i++ {
		if gcf_cache {
			c.(*gcf.GCFcache).Refer(gcf.Key{Value: i})
		} else {
			c.(*lru.LRUcache).Refer(func(i int) *int { return &i }(i))
		}
	}
	allocated = !allocated
	print2(nil)
	print2(
		fmt.Sprintf("Finished allocating the %s cache. (size=%s)", c.Name(),
			message.NewPrinter(language.English).Sprintf("%d", c.Size())),
	)
	latency := struct{ min, max, mean, mean2, c, c2 int64 }{0, 0, 0, 0, 0, 0}
	for i := 0; i < gcIterations; i++ {
		print2(nil)
		print2(heap(0))
		go func() {
			t0 := time.Now()
			time.Sleep(time.Millisecond * 10)
			latency_ := time.Now().Sub(t0).Microseconds() - 10000
			print2(fmt.Sprintf("Latency: %d µs", latency_))
			if latency.min == 0 || latency_ < latency.min {
				latency.min = latency_
			}
			latency.max = int64(math.Max(float64(latency.max), float64(latency_)))
			latency.mean += latency_
			// only count small latencies for mean(2)
			if latency_ < 100_000 {
				latency.mean2 += latency_
				latency.c2++
			}
			latency.c++
		}()
		t1 := time.Now()
		runtime.GC()
		t2 := time.Now().Sub(t1)
		print2(heap(1))
		print2("Time to perform a full GC: " + fmt.Sprintf("%d", t2.Milliseconds()) + " ms")
		print2("Sleeping for " + fmt.Sprintf("%d", t2.Milliseconds()/2) + " ms")
		time.Sleep(time.Millisecond * time.Duration(t2.Milliseconds()/2))
	}
	// latency stats
	latency.mean /= latency.c
	latency.mean2 /= latency.c2
	// best results so far on the fast cache: Latency (min, max, mean): 68 µs, 721 µs, 302 µs
	print2(fmt.Sprintf("Latency (min, max, mean, mean(2)): %d µs, %d µs, %d µs, %d µs",
		latency.min, latency.max, latency.mean, latency.mean2))
	// force memory to stay alive (this branch will never execute)
	if c.Size() == -1 {
		println(c.Size())
	}
}
