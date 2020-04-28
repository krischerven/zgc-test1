package fast

import (
	"container/list"
	"fmt"
	"unsafe"
)

func ptr(p *int) uintptr {
	return uintptr(unsafe.Pointer(p))
}

func ptr2(p *list.Element) uintptr {
	return uintptr(unsafe.Pointer(p))
}

func New(cap int) *LRUcache {
	return &LRUcache{
		cap:  cap,
		list: list.New(),
		map_: make(map[uintptr]uintptr, cap),
	}
}

type LRUcache struct {
	cap  int
	list *list.List
	// using uintptr as a key is probably equivalent to using
	// a *int as the key, but for the sake of consistency this way is preferred
	map_ map[uintptr]uintptr
}

func (l *LRUcache) Refer(key *int) {
	if l.no(key) {
		l.emplace(key)
	}
}

func (l *LRUcache) no(key *int) bool {
	if e, ok := l.map_[ptr(key)]; ok {
		l.list.Remove((*list.Element)(unsafe.Pointer(e)))
		l.list.PushFront(e)
		return false
	} else {
		return true
	}
}

func (l *LRUcache) Display() {
	for e := l.list.Front(); e != nil; e = e.Next() {
		fmt.Println(*(e.Value.(*int)))
	}
}

func (l *LRUcache) emplace(key *int) {
	if e, ok := l.map_[ptr(key)]; ok {
		l.list.Remove((*list.Element)(unsafe.Pointer(e)))
	} else if len(l.map_) == l.cap {
		delete(l.map_, ptr(l.list.Back().Value.(*int)))
		l.list.Remove(l.list.Back())
	}
	l.list.PushFront(key)
	l.map_[ptr(key)] = ptr2(l.list.Front())
}

func (l *LRUcache) Cap() int {
	return l.cap
}

func (l *LRUcache) Size() int {
	return len(l.map_)
}

func (l *LRUcache) Is(keys ...*int) bool {
	if l.Size() != len(keys) {
		return false
	} else {
		i := 0
		for e := l.list.Front(); e != nil; e = e.Next() {
			if *(e.Value.(*int)) != *keys[i] {
				return false
			} else {
				i++
			}
		}
		return true
	}
}

func (l *LRUcache) Free() {
	l.list.Init()
}
