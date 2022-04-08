package main
import "fmt"
func Myfunc (arg interface{}){
    fmt.Println("Myfunc is run ...")
    fmt.Println(arg)
    value, ok := arg.(string)
    if ok {
        fmt.Println("arg is string")
        fmt.Println("value = ", value)
    } else {
        fmt.Println("arg is not string")
    }
    fmt.Println("===========")
}
type Book struct{
    Name string
}
func main(){
    book := Book{"hello"}
    Myfunc(book)
    Myfunc(100)
    Myfunc("abcd")
    Myfunc(3.14)
}
