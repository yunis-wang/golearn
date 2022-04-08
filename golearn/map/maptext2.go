package main
import "fmt"
func main(){
    hash := make(map[string]string)

    // 添加
    hash["one"] = "golong"
    hash["two"] = "java"
    hash["three"] = "php"

    // 遍历
    for key, value := range hash{
        fmt.Println("key: ", key, ", value: ", value)
    }

    fmt.Println("================")

    // 删除
    delete(hash, "three")

    // 修改
    hash["two"] = "c++"

    // 遍历
    for key, value := range hash{
        fmt.Println("key: ", key, ", value: ", value)
    }

    printfmap(hash)

    fmt.Println("================")

    // 遍历
    for key, value := range hash{
        fmt.Println("key: ", key, ", value: ", value)
    }

}
func printfmap(myMap map[string]string){
    // 应用传递，传过来的是一个指针类型
    // 如果在函数中进行修改的话，原map也会被修改
    myMap["two"] = "c"
}
