package main
import "fmt"
func main(){
    // 第一种方式, 中括号里面是key类型，外面是value类型
    var hash map[string]string
    hash = make(map[string]string, 10)
    hash["hello"] = "word"
    fmt.Println(hash)
    // 第二种创建方式
    hash2 := make(map[int]string)
    hash2[1] = "golang"
    fmt.Println(hash2)
    // 第三种创建方式
    hash3 := map[string]string{
        "one" : "java",
        "two" : "c++",
    }
    fmt.Println(hash3)
}
