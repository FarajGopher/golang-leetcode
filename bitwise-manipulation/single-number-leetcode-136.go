package main
import "fmt"


func singleNumber(nums []int) int {
    m1 := make(map[int]int)
    for _,v := range nums{
        if _,ok := m1[v] ; ok{
            m1[v] = m1[v]+1
        }else{ 
            m1[v] = 1 
        }
    }

    for key,v := range m1 {
        if v == 1{
            return key
        }
    } 
    return -1
}

// optimized approach using XOR
//XOR help in finding the unique number in an array where every other number appears twice.
//XOR of a number with itself is 0 (a ^ a = 0).
//XOR of a number with 0 is the number itself (a ^ 0 = a).
//XOR is commutative and associative, meaning the order of operations does not matter.
func singleNumberXOR(nums []int) int {
    result := 0

    for _, num := range nums {
        result ^= num
    }

    return result
}


func main() {
    nums := []int{4,1,2,1,2}
    result := singleNumber(nums)
    fmt.Println("The single number is:", result)
    result = singleNumberXOR(nums)
    fmt.Println("The single number using XOR is:", result)

}