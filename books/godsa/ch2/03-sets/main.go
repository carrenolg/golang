package main

import "fmt"

type Set struct {
	intMap map[int]bool
}

// Create
func (set *Set) New() {
	set.intMap = make(map[int]bool)
}

// Add
func (set *Set) Add(element int) {
	_, ok := set.intMap[element]
	if !ok {
		set.intMap[element] = true
	}
}

// Delete
func (set *Set) Delete(element int) {
	delete(set.intMap, element)
}

func (set *Set) ContainsElement(element int) bool {
	var exists bool
	_, exists = set.intMap[element]
	return exists
}

// Intersection
func (set *Set) Intersect(anotherSet *Set) *Set {
	var intersectSet = &Set{}
	intersectSet.New()
	// match
	for value := range set.intMap {
		if anotherSet.ContainsElement(value) {
			intersectSet.Add(value)
		}
	}
	return intersectSet
}

// Union
func (set *Set) Union(anotherSet *Set) *Set {
	var unionSet = &Set{}
	unionSet.New()
	for value := range set.intMap {
		unionSet.Add(value)
	}
	for value := range anotherSet.intMap {
		unionSet.Add(value)
	}
	return unionSet
}

func main() {
	set := &Set{}
	set.New()
	set.Add(1)
	set.Add(2)
	fmt.Println(set)
	anotherSet := &Set{}
	anotherSet.New()
	anotherSet.New()
	anotherSet.Add(1)
	anotherSet.Add(15)
	//anotherSet.Add(2)
	fmt.Println(set.Intersect(anotherSet))
	union := set.Union(anotherSet)
	fmt.Println(union)
}
