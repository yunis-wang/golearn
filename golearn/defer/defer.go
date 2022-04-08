package main
import "fmt"
func main(){
    // defer 类似于c++析构函数，在函数的最后执行
    // 使用的是一种栈的结构，后进先出（先打印2在打印1）
    //defer fmt.Println("main end")
    //defer fmt.Println("main end2")
    //fmt.Println("first")
    //fmt.Println("next")
    text()
}
func text() int{
    defer deferFunc()
    return returnFunc()
}
func deferFunc() int{
    fmt.Println("defer is run....")
    return 0
}
func returnFunc() int{
    fmt.Println("return is run.....")
    return 0
}
