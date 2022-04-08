/*
 * @Description:
 * @version:
 * @Author: Wang
 * @Date: 2022-04-03 21:57:11
 * @LastEditors: Andy
 * @LastEditTime: 2022-04-03 22:00:21
 */
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
