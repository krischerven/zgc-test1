package simple

import (
	"container/list"
	"fmt"
)

func New(cap int) *LRUcache {
	return &LRUcache{
		cap:  cap,
		list: list.New(),
		map_: make(map[*int]*list.Element, cap),
	}
}

type LRUcache struct {
	cap  int
	list *list.List
	map_ map[*int]*list.Element
}

func (l *LRUcache) Refer(key *int) {
	if l.no(key) {
		l.emplace(key)
	}
}

func (l *LRUcache) no(key *int) bool {
	if e, ok := l.map_[key]; ok {
		l.list.Remove(e)
		l.list.PushFront(e)
		return false
	} else {
		return true
	}
}

func (l *LRUcache) emplace(key *int) {
	if e, ok := l.map_[key]; ok {
		l.list.Remove(e)
	} else if len(l.map_) == l.cap {
		delete(l.map_, l.list.Back().Value.(*int))
		l.list.Remove(l.list.Back())
	}
	l.list.PushFront(key)
	l.map_[key] = l.list.Front()
}

func (l *LRUcache) Display() {
	for e := l.list.Front(); e != nil; e = e.Next() {
		fmt.Println(*(e.Value.(*int)))
	}
}

func (l *LRUcache) Cap() int {
	return l.cap
}

func (l *LRUcache) Size() int {
	return len(l.map_)
}

func (l *LRUcache) elements() []int {
	ret := make([]int, l.Size())
	i := 0
	for e := l.list.Front(); e != nil; e = e.Next() {
		ret[i] = *(e.Value.(*int))
		i++
	}
	return ret
}

func (l *LRUcache) is(elems ...*int) bool {
	if l.Size() != len(elems) {
		return false
	} else {
		i := 0
		for e := l.list.Front(); e != nil; e = e.Next() {
			if *(e.Value.(*int)) != *elems[i] {
				return false
			} else {
				i++
			}
		}
		return true
	}
}
