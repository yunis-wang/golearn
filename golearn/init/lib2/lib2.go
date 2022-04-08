/*
 * @Description:
 * @version:
 * @Author: Wang
 * @Date: 2022-04-03 18:21:29
 * @LastEditors: Andy
 * @LastEditTime: 2022-04-03 21:02:07
 */
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
