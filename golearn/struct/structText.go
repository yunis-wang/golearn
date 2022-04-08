package main
import "fmt"
// type关键字是用来起别名的
type myint int
// 定义一个结构体，名字是book， 创建结构体的关键字struct
type Book struct{
    name string
    auto string
}
func main(){
    var a myint = 0
    fmt.Printf("a: %d, %T\n", a, a)
    var book Book
    book.auto = "zhangsan"
    book.name = "golang"
    fmt.Println(book)
    chage1(book)
    fmt.Println(book)
    chage2(&book)
    fmt.Println(book)
}
func chage1(book Book){
    // 一般是值传递修改不影响原来
    book.name = "lisi"
}
func chage2(book *Book){
    // 指针传递传过来的是一个地址
    book.auto = "lisi"
}
