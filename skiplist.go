package godsa

import (
	"errors"
	"math/rand"
)

var insertMap map[int]*node

type node struct {
	key   int
	value int
	next  []*node // next[0] represents the lowest level
	prevs []*node
}

type SkipList struct {
	sentinel    *node
	maxLevel    int
	probability float64
}

func (sl *SkipList) flipCoin() bool {
	return rand.Float64() < sl.probability
}

func (sl *SkipList) newNode(key, value int) *node {
	level := 1 // all nodes always have next[0]
	for sl.flipCoin() && level < sl.maxLevel {
		level += 1
	}
	return &node{
		key:   key,
		value: value,
		next:  make([]*node, level),
		prevs: make([]*node, level),
	}
}

func NewSkipList(maxLevel int, probability float64) *SkipList {
	sl := &SkipList{
		maxLevel:    maxLevel,
		probability: probability,
	}
	sl.sentinel = &node{
		key:   -1,
		value: -1,
		next:  make([]*node, maxLevel),
	}
	insertMap = make(map[int]*node)
	return sl
}

func (sl *SkipList) Search(key int) (*node, error) {
	if key < 0 {
		return nil, errors.New("key must >= 0")
	}
	insertMap = make(map[int]*node)

	cur := sl.sentinel
	for i := sl.maxLevel - 1; i >= 0; i-- {
		flag := false
		if key == cur.key {
			insertMap[i] = cur
			return cur, nil
		}
		if cur.next[i] != nil {
			for key > cur.next[i].key {
				cur = cur.next[i]
				if cur.next[i] == nil {
					flag = true
					break
				}
			}
			if flag {
				insertMap[i] = cur
				continue
			}
			if key == cur.next[i].key {
				insertMap[i] = cur.next[i]
				return cur.next[i], nil
			}
		}
		insertMap[i] = cur
	}
	return nil, errors.New("key does not exist")
}

func completeInsertMap() {
	minLevel := 100000
	for k := range insertMap {
		if k <= minLevel {
			minLevel = k
		}
	}
	for i := 0; i < minLevel; i++ {
		insertMap[i] = insertMap[minLevel]
	}
}

func (sl *SkipList) Insert(key, value int) error {
	if key < 0 {
		return errors.New("key must >= 0")
	}
	n := sl.newNode(key, value)
	res, err := sl.Search(key)
	if err != nil {
		completeInsertMap()
		for i := 0; i < len(n.next); i++ {
			nextNode := insertMap[i].next[i]
			n.next[i] = nextNode
			insertMap[i].next[i] = n

			n.prevs[i] = insertMap[i]
			if nextNode != nil {
				nextNode.prevs[i] = n
			}
		}
	} else {
		res.value = value
	}
	return nil
}

func (sl *SkipList) Delete(key int) error {
	if key < 0 {
		return errors.New("key must >= 0")
	}
	res, err := sl.Search(key)
	if err != nil {
		return errors.New("key does not exist")
	}
	for i := 0; i < len(res.prevs); i++ {
		res.prevs[i].next[i] = res.next[i]
		if res.next[i] != nil {
			res.next[i].prevs[i] = res.prevs[i]
		}
	}
	return nil
}
