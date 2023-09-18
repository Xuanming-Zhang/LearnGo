package main
import (
	"fmt"
	"sort"
)

func main(){
	drinks := make(map[string]string,3)
	drinks["coke"]="可乐"
	drinks["sprite"]="雪碧"
	drinks["fenda"]="芬达"
	for k := range(drinks){
		fmt.Printf("饮料 %v/",k)
	}
	fmt.Println()
	for k := range(drinks){
		fmt.Printf("饮料 %v，翻译为%v/",k,drinks[k])
	}
	keys := make([]string,len(drinks))
	i := 0
	for k := range(drinks){
		keys[i]=k
		i+=1
	}
	sort.Strings(keys)
	fmt.Println()
	fmt.Println("After sort")
	for _,k := range(keys){
		fmt.Printf("饮料 %v，翻译为%v/",k,drinks[k])
	}
}
