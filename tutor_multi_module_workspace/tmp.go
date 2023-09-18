package main

import "fmt"

func sum(arr ...int) int{
    t := 0
    for _,x := range(arr){
        t += x
    }
    return t
}

func main(){
    fmt.Println(sum(1,2,3))
    arr := []int{1,2,3,4}
    fmt.Println(sum(arr...))
}
