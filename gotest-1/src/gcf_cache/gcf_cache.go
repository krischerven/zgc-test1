package gcf_cache

import (
	"fmt"
)

// A non-pointer key
type Key struct {
	Value int
}

// a [G]C [F]riendly cache that stores memory with extremely little dynamic
// allocation. Even massive caches will have little to no effect on the Go GC,
// so long as you do not make Key a pointer type.
type GCFcache struct {
	imap       map[uint64]Key
	kmap       map[Key]struct{}
	imap_start uint64
	cap        uint32
}

// if you want a pointer you have to dereference it
func New(cap uint32) GCFcache {
	return GCFcache{
		make(map[uint64]Key, int(cap)),
		make(map[Key]struct{}, int(cap)),
		0,
		cap,
	}
}

// or just use this
func NewPtr(cap uint32) *GCFcache {
	ret := New(cap)
	return &ret
}

// GCFcache logic
func (g *GCFcache) Refer(key Key) {
	if g.no(key) {
		g.emplace(key)
	}
}

func (g *GCFcache) Hit(key Key) bool {
	_, ok := g.kmap[key]
	return ok
}

func (g *GCFcache) no(key Key) bool {
	_, ok := g.kmap[key]
	return !ok
}

func (g *GCFcache) emplace(key Key) {
	g.imap[uint64(len(g.imap))+g.imap_start] = key
	g.kmap[key] = struct{}{}
	if uint32(len(g.imap)) > g.cap {
		delete(g.kmap, g.imap[g.imap_start])
		delete(g.imap, g.imap_start)
		g.imap_start++
	}
}

func (g *GCFcache) Display() {
	for key := range g.kmap {
		fmt.Println(key)
	}
}

func (g *GCFcache) Cap() int {
	return int(g.cap)
}

func (g *GCFcache) Size() int {
	return len(g.kmap)
}

func (g *GCFcache) Name() string {
	return "GCF Cache"
}

func (g *GCFcache) elements() []Key {
	ret := make([]Key, g.Size())
	i := 0
	for key := range g.kmap {
		ret[i] = key
		i++
	}
	return ret
}

func (g *GCFcache) is(elems ...Key) bool {
	if g.Size() != len(elems) {
		return false
	} else {
		i := 0
		for _, key := range g.elements() {
			if key != elems[i] {
				return false
			} else {
				i++
			}
		}
		return true
	}
}
