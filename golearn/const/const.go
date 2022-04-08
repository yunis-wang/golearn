/*
 * @Description:
 * @version:
 * @Author: Wang
 * @Date: 2022-04-03 15:02:06
 * @LastEditors: Andy
 * @LastEditTime: 2022-04-03 16:15:59
 */
package main

import "fmt"

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
)

func main() {
	// 常量（只读属性）
	const pai = 3.14
	fmt.Println("pai = ", pai)
	// pai = 4.13 // cannot assign to pai (declared const)
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	fmt.Println("c = ", c)
	fmt.Println("d = ", d)
	fmt.Println("e = ", e)
}
