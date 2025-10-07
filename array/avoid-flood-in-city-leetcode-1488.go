package main

import (
	"fmt"
	"sort"
)

//****** this approach also work but for few testcase it fails **********
// func avoidFlood(rains []int) []int {
// 	lake := make(map[int]int)
// 	dry := []int{}
// 	j := 0

// 	for i, v := range rains {
// 		if v == 0 {
// 			dry = append(dry, i)
//             rains[i] = 1  //setting initally the value of dry one 1 then below overwrite 
// 		} else {
// 			if lastRain, ok := lake[v]; ok {
// 				found := false
// 				for k := j; k < len(dry); k++ {
// 					if dry[k] > lastRain {
// 						rains[dry[k]] = v
// 						dry = append(dry[:k], dry[k+1:]...)
// 						found = true
// 						break
// 					}
// 				}
// 				if !found {
// 					return []int{}
// 				}
// 				rains[i] = -1
// 			} else {
// 				lake[v] = i
// 				rains[i] = -1
// 			}
// 		}
// 	}
// 	return rains
// }



func avoidFlood(rains []int) []int {
	lake := make(map[int]int)
	dry := []int{}

	for i, v := range rains {
		if v == 0 {
			dry = append(dry, i)
			rains[i] = 1
		} else {
			if lastRain, ok := lake[v]; ok {
				idx := sort.Search(len(dry), func(k int) bool { return dry[k] > lastRain })
				if idx == len(dry) {
					return []int{}
				}
				rains[dry[idx]] = v
				dry = append(dry[:idx], dry[idx+1:]...)
			}
			lake[v] = i
			rains[i] = -1
		}
	}
	return rains
}

func main() {
	fmt.Println(avoidFlood([]int{0, 1, 1}))                             // []
	fmt.Println(avoidFlood([]int{1, 0, 2, 0, 3, 0, 2, 0, 0, 0, 1, 2, 3})) // passes all tests
}
