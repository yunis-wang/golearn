package main
import "fmt"
func main(){
    // 动态数组的长度为3，整体的容量为5
    num := make([]int, 3, 5)
    fmt.Printf("len: %d, cap: %d, slice: %v\n", len(num), cap(num), num)
    // 向末尾添加元素
    num = append(num, 1)
    num = append(num, 1)
    fmt.Printf("len: %d, cap: %d, slice: %v\n", len(num), cap(num), num)
    // 超过容量以后继续添加，此时cap会进行扩容
    num = append(num, 1)
    // 扩容的大小就是cap的大小
    fmt.Printf("len: %d, cap: %d, slice: %v\n", len(num), cap(num), num)

    // 当不指定容量大小时，就会默认的将长度大小设置为容量大小
    num2 := make([]int, 3)
    fmt.Printf("len: %d, cap: %d, slice: %v\n", len(num2), cap(num2), num2)

    // 表示截取 ，左闭右开，[2 : 4] 包括2，3两个位置上的元素，如果前面或后面不写的话表示0号位置，或者最后一个位置
    // 此时底层num3 和num 实际上指向的是同一片内存空间（类似于浅拷贝）修改num3时num也会发生改变
    num3 := num[2 : 4]
    fmt.Println(num3)

    // 使用拷贝函数进行深拷贝
    num4 := make([]int, 2)
    copy(num4, num[3:5])
    fmt.Println(num4)
}
