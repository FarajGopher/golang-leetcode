package main

import (
	"fmt"
	"sort"
)

func successfulPairs(spells []int, potions []int, success int64) []int {
	sort.Ints(potions)
    for i := 0 ;i< len(spells);i++{
        s:=0
        l:= len(potions)-1
        for s<=l{
            mid := s + (l-s)/2

            if spells[i] * potions[mid] >= int(success) {
                l = mid - 1 
            }else {
                s = mid + 1
            }
        }
        spells[i] = len(potions)-s
    } 

    return spells
}

func main() {
	spells := []int{5, 1, 3}
	potions := []int{1, 2, 3, 4, 5}
	success := int64(7)

	fmt.Println(successfulPairs(spells, potions, success)) // Output: [4 0 3]
}
