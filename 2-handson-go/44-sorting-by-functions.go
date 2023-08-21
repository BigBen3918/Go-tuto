// Sometimes we’ll want to sort a collection by something other than its natural order. For example, suppose we wanted to sort strings by their length instead of alphabetically. Here’s an example of custom sorts in Go.

package main

import (
    "fmt"
    "sort"
)

type byLength []string

func (s byLength) Len() int {
    return len(s)
}
func (s byLength) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}
func (s byLength) Less(i, j int) bool {
    return len(s[i]) < len(s[j])
}

func main() {
    fruits := []string{"peach", "banana", "kiwi"}
    sort.Sort(byLength(fruits))
    fmt.Println(fruits)
}