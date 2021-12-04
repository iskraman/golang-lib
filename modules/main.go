package main

import (
	"fmt"

	"github.com/iskraman/golang-lib/modules/slice"
)

func main() {
	items := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(slice.SliceExists(items, 5))   // returns true
	fmt.Println(slice.SliceExists(items, "5")) // returns false
	fmt.Println(slice.SliceExists(items, 10))  // returns false
}
