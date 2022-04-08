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
