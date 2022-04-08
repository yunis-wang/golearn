/*
 * @Description:
 * @version:
 * @Author: Wang
 * @Date: 2022-04-03 18:20:49
 * @LastEditors: Andy
 * @LastEditTime: 2022-04-03 21:22:12
 */
package main

import (
	// 导入相关的包
	// 再导入过程中会执行，包中的init方法
	// 下划线空格包，这种形式是给包一个匿名，无法使用包中的方法，但是可以执行init方法
	_ "./lib1"
	// 别名空格包，这种形式是给包起一个别名然后使用别名来使用对应的包
	mylib2 "./lib2"
	// 点空格包，将包中的所有方法全部都导入到当前包中，可以直接使用，不需要进行点
	//. "./lib2"
)

func main() {
	// 调用函数
	mylib2.Hello()
}
