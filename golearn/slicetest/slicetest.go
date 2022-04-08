package main
import "fmt"
// 传递参数时大小也必须相同
func prinfArray(p [4]int){
    for index, value := range p{
        fmt.Println("index: ", index, ", value: ", value)
    }
}
func main(){
    // 固定长度的数组
    a := [10]int{1, 2, 3, 4}
    for i := 0; i < len(a); i++{
        fmt.Println(a[i])
    }
    b := [10]int{2, 3, 4, 5}
    // range 是切片 有两个返回值返回对应数组的下标以及元素值
    for index, value := range b{
        fmt.Println("index: ", index, ", value: ", value)
    }
    c := [4]int{11, 22, 33, 44}
    prinfArray(c)
}
