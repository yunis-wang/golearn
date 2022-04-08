<!--
 * @Description: 
 * @version: 
 * @Author: Wang
 * @Date: 2022-04-03 15:57:19
 * @LastEditors: Andy
 * @LastEditTime: 2022-04-06 19:42:02
-->
# Go基础

## hello word

```go
package main

import "fmt"

func main() {
    fmt.Println("hello go")
}
```

## 变量的声明

1. 声明一个变量指定数据类型 `var a int = 10`
2. 在初始化是省略数据类型 `var b = "abc"`
3. (常用)省略var关键字 `c := 10`

> 注：声明全局变量时不能使用第三种方法

扩展：
声明多个变量

```go
// 在一行
var d, e = 123, "abc"
// 多行
var (
    f = 123
    g = "ABC"
)
```

## 常量的定义

==const==关键字，用来表明这是一个常量 `const pai = 3.14`
同时==const(...)==可以用于声明多个常量

```go
// const 定义枚举类型
const (
    a = 1
    b = 2
    // 可以在其中添加一个iota关键字，每行的iota都会进行加1，第一行默认为0
    //从const的第一行计数
    // iota只能配合const()一起使用
    c = iota * 2
    d
    e
 /*
    a =  1
    b =  2
    c =  4
    d =  6
    e =  8
 */
)
```

## 函数调用以及有多个返回值的情况

```go
package main

import "fmt"

func main() {
    c := add(1, 2)
    fmt.Println(c)
    x, y := swap(5, 6)
    fmt.Println("x = ", x, " y = ", y)
    x2, y2 := swap(5, 6)
    fmt.Println("x2 = ", x2, " y2 = ", y2)
}

/**
 * @Author: wang
 * @description: 相加函数
 * @param {int} a
 * @param {int} b
 * @return {*}
 */
func add(a int, b int) int {
    return a + b
}

/**
 * @Author: wang
 * @description: 多个返回值的函数
 * @param {*} a
 * @param {int} b
 * @return {*}
 */
func swap(a, b int) (int, int) {
    return b, a
}

/**
 * @Author: wang
 * @description: 返回多个有行参名称的
 * @param {*} a
 * @param {int} b
 * @return {*}
 */
//func swap2(a, b int) (c, d int) 这样写也是可以的
func swap2(a, b int) (c int, d int) {
    // 给有名字的返回值赋值
    // 两个返回变量有默认的初始值
    c = b
    d = a
    return
    // return b, a 也可以
}

```

## 导包以及函数的调用

首先包的结构

| ---lib1 |
| ------- ||lib1
| ---lib2 |
| ------- ||lib2
|---main

> lib2

```go
package lib2

import "fmt"

// 函数名小写只能在当前包内调用
func init() {
    fmt.Println("lib2...")
}

//如果函数名大写的话代表是一个对外开放的函数
func Hello() {
    fmt.Println("welcome to lib2")
}
```

> lib1

```go
package lib1

import "fmt"

func init() {
    fmt.Println("lib1...")
}

func Hello() {
    fmt.Println("welcome to lib1")
}
```

> main

```go
package main

import (
    // 导入相关的包
    // 再导入过程中会执行，包中的init方法
    "./lib1"
    "./lib2"
)

func main() {
    // 调用函数
    lib1.Hello()
    lib2.Hello()
}
```

> 注意：
>
> 1. 函数名大小写问题（函数名首字母大写代表是对外开放的函数，首字母小写只能在当前包内调用）
>
> 2. 调用过程中首先会执行init方法

导包方式

1. 下划线空格包，这种形式是给包一个匿名，无法使用包中的方法，但是可以执行init方法 `"./lib1"`
2. 别名空格包，这种形式是给包起一个别名然后使用别名来使用对应的包 `mylib2 "./lib2"`
3. 点空格包，将包中的所有方法全部都导入到当前包中，可以直接使用，不需要进行点 `. "./lib2"`

## 指针

```go
package main

import "fmt"

func main() {
    a := 10
    b := 20
    swap(&a, &b)
    fmt.Println("a: ", a, "b: ", b)
}

func swap(a, b *int) {
    c := *a
    *a = *b
    *b = c
}
```

## defer关键字

```go
package main
import "fmt"
func main(){
    // defer 类似于c++析构函数，在函数的最后执行
    // 使用的是一种栈的结构，后进先出（先打印2在打印1）
    defer fmt.Println("main end")
    defer fmt.Println("main end2")
    fmt.Println("first")
    fmt.Println("next")
}
```

如果return和defer同时存在的话先执行return后执行return

```go
package main
import "fmt"
func main(){
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

```

## 数组

### 固定数组

1. 固定数组需要自己去指定一个数组的大小
2. 函数传递时是一种拷贝的方式进行值传递
3. 传递时数组的大小也必须一致

```go
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

```

### 动态数组

1. 不需要指定大小
2. 传递时使用的是引用传递的方式

```go
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
    // 通过make来添加空间
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

```

==注：使用nil来表示空（类似于null）==

### slice的追加以及截取操作

1. 有长度和容量两个属性，不指定容量是默认容量和长度一致
2. 当向后追加超过容量是进行扩容，每次扩容的大小就是原容量的大小
3. 截取数组时，实际上指向的是同一片内存空间，**修改新数组原数组也会发生改变**
4. 使用copy函数进行深拷贝

```go
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
```

## map

```go
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
```

### map的增删改查， 以及传递

```go
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
```

## 结构体

使用关键字struct定义 , 然后使用type起一个别名

```go
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
```

## 对象

go对象是由结构体以及与这个结构体绑定的函数来实现的

同时首字母大写代表外界能够访问，首字母小写代表外界无法访问

### 封装

```go
package main
import "fmt"
// 首字母大写表示其他包中的内容也能访问
type Hero struct{
    // 首字母大写表示其他包中的内容也能访问这个属性，相当于public
    Name string
    Level int
}
/*
// 通过括号以及传递相关的结构体的方式实现方法与结构的绑定
func (this Hero) GetName() string{
    return this.Name
}
// 这种绑定方式，由于传递的是值所有没有使用set方法不会发生改变
func (this Hero) SetName(name string){
    this.Name = name
}
func (this Hero) Show(){
    fmt.Println("name: ", this.Name)
    fmt.Println("level: ", this.Level)
}*/ 
func (this *Hero) GetName() string{
    return this.Name
}
// 使用指针的绑定方式这样就可以实现修改内部的值
func (this *Hero) SetName(name string){
    this.Name = name
}
func (this *Hero) Show(){
    fmt.Println("name: ", this.Name)
    fmt.Println("level: ", this.Level)
}
 
func main(){
    hero := Hero{Name: "zhangsan", Level: 1}
    hero.Show()
    hero.SetName("lisi")
    hero.Show()
}
```

### 继承

```go
package main
import "fmt"
type Human struct{
    Name string
    Sex string
}
func (this *Human) Work(){
    fmt.Println("Human work()...")
}
func (this *Human) Eat(){
    fmt.Println("Human eat()...")
}
type Superman struct{
    // Supenman继承了human中的方法
    Human
    Level int
}
// 重写父类方法
func (this *Superman) Eat(){
    fmt.Println("superman eat()....")
}
// 添加子类的新方法
func (this *Superman) Fly(){
    fmt.Println("superman fly()...")
}
func main(){
    // 父类对象，注意传值的时候是打括号
    h := Human{"zhagsan", "nan"}
    h.Work()
    h.Eat()
    // 定义一个子类对象， 先定义父类然后定义自己的方法
    s := Superman{Human{"lisi", "nv"}, 10}
    s.Work()
    s.Eat()
    s.Fly()
    // 也可以直接传递已有的父类
    s2 := Superman{h, 9}
    s2.Work()
}
```

### 多态

1. 有一个父类（接口）
2. 有子类（实现了父类的全部接口方法）
3. 父类类型的变量实际上是一个指针，需要指向子类的具体数据变量

```go
package main

import "fmt"

// 定义一个动物的接口
// 实现接口时只需要具体类里面有接口中的所有方法就默认为实现了
type Animal interface{
    Sleep()
    GetColor() string
    GetType() string
}
// 定义一只具体的动物
type Cat struct{
    Color string
}
// 实现接口中的三个方法
func (this *Cat) Sleep(){
    fmt.Println("cat sleep()...")
}
func (this *Cat) GetColor() string{
    return this.Color
}
func (this *Cat) GetType() string{
    return "cat"
}
type Dog struct{
    Color string
}
func (this *Dog) Sleep(){
    fmt.Println("dog sleep()...")
}
func (this *Dog) GetColor() string{
    return this.Color
}
func (this *Dog) GetType() string{
    return "dog"
}
func ShowAnimal(animal Animal){
    animal.Sleep()
    fmt.Println("color: ", animal.GetColor())
    fmt.Println("type: ", animal.GetType())
}
func main(){
    // 父类变量实际上是一个指针
    var animal Animal
    // 传递一个具体子类的地址
    animal = &Cat{"Green"}
    animal.sleep()
    cat := Cat{"yellow"}
    dog := Dog{"green"}
    ShowAnimal(&cat)
    ShowAnimal(&dog)
}
```

### interface

首先是可以用作是一个接口

为空时就是一个所有类型的父类，接受任意类型的参数

通过断言来判断他的具体类型

```go
package main
import "fmt"
// 所有的类型都继承与interface{}相当于Java中的object类
// 可以接受任意类型的参数
func Myfunc (arg interface{}){
    fmt.Println("Myfunc is run ...")
    fmt.Println(arg)
    // 断言机制，通过向下转型的方式判断传入的类型是什么，有两个返回值
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
```
