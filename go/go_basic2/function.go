package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	name := flag.String("name", "world", "specify the name you want to say hi")
	// 据这话的意思是name的默认值是world
	flag.Parse()
	fmt.Println("os args is:", os.Args)
	fmt.Println("input parameter is:", *name)
	fullString := fmt.Sprintf("Hello %s from Go\n", *name)
	fmt.Println(fullString)
	err, result := DuplicateString("aaa")
	if err == nil {
		fmt.Println(result)
	} else {
		fmt.Println(err)
	}

	// 回调函数示例

	DoOperation(1, increase)
	DoOperation(1, decrease)

	//匿名函数 意味着这里面的逻辑我希望是独立运行的
	defer func() {
		if r := recover(); r != nil {
			println("recover in FuncX")
		}
	}()
}

func increase(a, b int) {
	println("increase result is:", a+b)
}

func DoOperation(y int, f func(int, int)) {
	// 给了值，给了函数，在这里用你给我的值调用了给我的方法
	f(y, 1)
}

func decrease(a, b int) {
	println("decrease result is", a-b)
}

//Go语言函数可以用多返回 .这个和Java是不一样的
// func DuplicateString(input string) (error, string) {
// 	if input == "aaa" {
// 		return fmt.Errorf("aaa is not alowed"), ""
// 	}
// 	return nil, input + input
// }
func DuplicateString(input string) (err error, result string) {
	if input == "aaa" {
		err = fmt.Errorf("aaa is not alowed")
		return
	}
	result = input + input
	return
}

/*
Go语言和Java不同，没有类似Java的[]string args参数
Go语言如何传入参数呢？
	方法1：
		fmt.Println("os args is: ", os.Args)

	方法2：
		name := flag.String("name", "world", "specify the name you want to say hi")
		据这话的意思是name的默认值是world, 最后一个参数是name的注解
		flag.Parse()

返回值
	多值返回
		函数可以返回任意数量的返回值
		多返回值的应用场景？错误处理
	其他语言是exception和返回值，是不同的关键字，但是在
	Go语言里引入了error的接口，error就是一个message，当作一个返回值一样返回的
	如果错了，就是返回error和一个空的返回值
	如果没错，就是返回一个空的错误和一个真实的返回值。
	看有没有错都是看error是不是为空error == nil

	返回值是可以被命名的
	命名的意思是

	func DuplicateString(input string) (error, string) {
	if input == "aaa" {
		return fmt.Errorf("aaa is not alowed"), ""
	}
	return nil, input + input
}
等价于
func DuplicateString(input string) (err error, result string) {
	if input == "aaa" {
		err = fmt.Errorf("aaa is not alowed")
		return
	}
	result = input + input
	return
}

Go里几乎都是多返回值，很少有单返回值的情况

Go语言中的可长参数允许调用方传递任意多个相同类型的参数
函数定义
func append(slice []Type elems ...Type) []Type
在类型前加三个点意味着elems这参数是可变参数
myarray := []string{}
myarray = append(myArray, "a", "b", "c")

内置函数
close 管道关闭
len,cap 返回数组、切片，Map的长度或者容量
new, make 内存分配 new返回地址， make返回
copy append 操作切片
panic recover 错误处理
complex real imag 操作复数

回调函数(Callback)
函数作为参数传入其他函数，并在其他函数内部调用执行
	strings.IndexFunc(line, unicode.isSpace)
	Kubernets controller的leaderelection

什么场景会用呢，一般是你用一个通用的framework例如树遍历
针对树的某个节点做某一些操作，而操作在不同的情形下，调用的函数是不一样的。

函数的存在的意义是为了逻辑抽象，很多时候我有一个逻辑被多个地方调用，或者是一个完整的逻辑，为了重用

如果有逻辑我希望独立执行，就有匿名函数
闭包
	匿名函数
		不能独立存在
		可以赋值给其他变量
		x := func(){} 这是一个函数声明
		可以直接调用
		func(x,y,int){
			println(x + y)
		}(1,2)  //加上圆括号
		可以作为函数的返回值。
			func Add() (func(b int) int)

*/
