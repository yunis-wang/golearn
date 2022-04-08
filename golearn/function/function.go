/*
 * @Description:
 * @version:
 * @Author: Wang
 * @Date: 2022-04-03 16:21:13
 * @LastEditors: Andy
 * @LastEditTime: 2022-04-03 16:36:58
 */
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
