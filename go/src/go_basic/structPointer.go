package main

import "reflect"

type IF interface {
	getName() string
}

type Human struct {
	firstName, lastName string
}

type MyType struct {
	Name string `json:"name"` //注意这个标点是``不是''
	// 这个就是结构体标签
}

func (h *Human) getName() string {
	return h.firstName + "," + h.lastName
}

//Go里面就是两种，一种就是结构体，只有属性
// 接口 一组函数的抽象，定义行为

type Car struct {
	factory, model string
}

func (c *Car) getName() string {
	return c.factory + "-" + c.model
}
func main() {
	// h1 := Human{}
	// p1 := &h1
	interfaces := []IF{}
	h := new(Human)
	// // 通过new来创建struct，这个h就是这个对象的指针
	h.firstName = "first"
	h.lastName = "last"
	interfaces = append(interfaces, h)
	c := new(Car)
	c.factory = "benz"
	c.model = "s"
	interfaces = append(interfaces, c)

	mt := MyType{Name: "test"}
	myType := reflect.TypeOf(mt)
	name := myType.Field(0)
	tag := name.Tag.Get("json")
	println(tag)
}

// func printMyType(t *MyType) {
// 	println(t.Name)
// }

/*
结构体和指针
 通过type ... struct关键字自定义结构体
 Go语言支持指针，但不支持指针运算
 	指针变量的值为内存地址
 	为负值的指针为nil

结构体标签
	结构体还可以有一个可选的标签（tag）
	使用场景：Kubernetes APIServer 对所有资源的定义都用Json tag和protoBuff tag

*/
