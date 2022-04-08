/*
 * @Description: 变量学习
 * @version:
 * @Author: Wang
 * @Date: 2022-04-03 11:22:45
 * @LastEditors: Andy
 * @LastEditTime: 2022-04-03 15:02:52
 */
package main

import "fmt"

// 声明全局变量时，只能使用前两种方法，最后一种方法会报错

func main() {
	// 方法一：声明一个变量，默认值是0,同时可以指定一个初始值
	var a int = 10
	fmt.Println("a = ", a)
	// 方法二： 在初始化的时候省略数据类型，自动能够根据数据进行类型的匹配
	var b = "hello"
	fmt.Println("b = ", b)
	// 方法三：（常用） 省去var关键字直接进行自动匹配
	c := 10000
	fmt.Println("c = ", c)
	fmt.Printf("c = %T\n", c)

	// 声明多个变量
	var d, e = 123, "abc"
	var (
		f = 123
		g = "ABC"
	)
	fmt.Println("d = ", d, "e = ", e)
	fmt.Println("f = ", f, "g = ", g)
}
