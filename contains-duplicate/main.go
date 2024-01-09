package main

func containsDuplicate(nums []int) bool {
    numbers:= make(map[int]int,0)
    for _,val := range nums{
        numbers[val]++
        if numbers[val]>1{
            return true
        }
    }
    return false
}