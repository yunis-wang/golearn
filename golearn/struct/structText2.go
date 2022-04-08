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
