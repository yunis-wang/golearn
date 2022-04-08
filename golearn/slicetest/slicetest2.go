package main

import (
	"fmt"
)
func main(){
    // 第一种方式直接进行初始化
    myArray := []int{1, 2, 3, 4} // 动态数组,slice
    // 第二种方式，先不分配空间，后面在添加
    var slice2 []int
    if(slice2 == nil){
        fmt.Println("当前为空")
    }
    slice2 = make([]int, 3)
    // 第三种方式直接分配空间但是不进行初始化
    //slice3 := make([]int, 4)
    printfArray(myArray)
}
func printfArray(array []int){
    // 动态数组是一个引用传递，而固定数组是值传递
    // _ 表示一个匿名变量不关心
    for _, value := range array{
        fmt.Println("value: ", value)
    }
}
