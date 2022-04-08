package main

import "fmt"

func BitmaskToNumset() {
	bitmask := uint64(6)
	numset := make([]uint64, 0, 8)

	for i := 0; i <= 63; i++ {
		filter := uint64(1 << i)
		fmt.Println(filter)
		if bitmask&filter == filter {
			fmt.Println("append ", filter)
			numset = append(numset, filter)
		}
	}
}
