package main

import (
	"fmt"
	"math/rand"
	"time"
)

type RandomizedSet struct {
	hMap map[int]int
	list []int
}

func Constructor() RandomizedSet {
	rand.Seed(time.Now().UnixNano())
	return RandomizedSet{
		list: []int{},
		hMap: make(map[int]int),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, existed := this.hMap[val]; existed {
		return false
	}

	this.list = append(this.list, val)
	this.hMap[val] = len(this.list) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	idx, existed := this.hMap[val]
	if !existed {
		return false
	}
	last := len(this.list) - 1
	if idx != last {
		this.list[idx] = this.list[last]
		this.hMap[this.list[idx]] = idx
	}

	this.list = this.list[:last]
	delete(this.hMap, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.list[rand.Intn(len(this.list))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */

func main() {
	fmt.Println("Solution goes here!")
}
