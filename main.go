package main

import (
	"fmt"
	"strconv"
	"unicode"
)

var tracker = Tracker{}

func main() {
	tracker.trk = make(map[string][]int)
	fmt.Println("testing it")
	val := nextServerNumber([]int{5, 4, 1, 2})
	fmt.Println("result is: ",val)
	// server are numbered sequentially starting from 1
	// we don't know if server will stay up or not
	// if server crashes, it's number is made availble for use
	// new server should have lowest number
	// return the number of the next server to allocate from a slice
	// write unit test


	//part 2
	//server name only letetrs i nname concatenated with number

	//name tracking module with 2 ops
	//map host server number

}
func nextServerNumber(arr []int) int {
	next := 1
	foundNext := false

	for !foundNext {
		if !contains(arr,next) {
			return next
		}
		next++
	}


	return next
}

func contains(arr []int, val int) bool {
	for _,v := range arr {
		if v == val {
			return true
		}
	}
	return false
}

//returns server+number
func allocate(param string) string {
	_, ok := tracker.trk[param]
	if !ok {
		tracker.trk[param] = []int{}
	}

	nb := nextServerNumber(tracker.trk[param])
	tracker.Allocate(param,nb)

	return param+strconv.Itoa(nb)
}

//release the host name containing the number
func deallocate(param string) error {
	runes := []rune(param)
	index := 0
	for i,r := range runes {
		if unicode.IsDigit(r) {
			index = i
			break
		}
	}
	nb, err := strconv.Atoi(param[index+1:])
	if err != nil {
		return err
	}
	host := param[:index]
	tracker.Deallocate(host,nb)

	return nil
}

type Tracker struct {
	trk map[string][]int
}

func (t *Tracker) Allocate(host string, nb int) {
	v, ok := t.trk[host]

	if ok {
		t.trk[host] = append(v,nb)
	}
}

func (t *Tracker) Deallocate(host string, nb int) {
	v, ok := t.trk[host]

	if ok {
		t.trk[host] = remove(v,nb)
	}
}


func remove(sl []int, nb int) []int {
	ind := -1

	for i, n := range sl {
		if n == nb {
			ind = i
			break
		}
	}
	if ind != -1 {
		sl = append(sl[0:ind],sl[ind+1:]...)
	}
	return sl
}