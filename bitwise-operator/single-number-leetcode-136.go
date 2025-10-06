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


func main() {
    nums := []int{4,1,2,1,2}
    result := singleNumber(nums)
    fmt.Println("The single number is:", result) // Output: The single number is: 4
}