package main

import (
	"fmt"
)

func main() {
	myArray := [5]int{1, 2, 3, 4, 5}
	mySlice := myArray[1:3]
	// 切片用的比数组做，原因就如下所示，我的切片不用管长度，相加就加，由Go帮我来维护。
	// 和c++的vector差不多，但是操作简单很多
	mySlice1 := []int{}
	// 添加很容易。
	mySlice1 = append(mySlice1, 1)
	mySlice1 = append(mySlice1, 2)
	mySlice1 = append(mySlice1, 3)

	// 删除不是很容易，原因是没内置函数，自己写
	fmt.Printf("mySlice %+v\n", mySlice)
	fullSlice := myArray[:]
	remove3rfItem := deleteItem(fullSlice, 2)
	fmt.Printf("remove3rdItem %+v\n", remove3rfItem)

	//Map
	myMap := make(map[string]string, 10)
	myMap["a"] = "b" // key-value a-b
	myFuncMap := map[string]func() int{
		"funcA": func() int { return 1 }, //为什么这里有逗号
	} // Map里面的类型可以很复杂 string-function类型
	fmt.Println(myFuncMap)
	f := myFuncMap["funcA"]
	fmt.Println(f())
	value, exist := myMap["a"]
	if exist {
		println(value)
	}
	for k, v := range myMap {
		println(k, v)
	}
}

func deleteItem(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}

/*
变量的初始化
var i, j int = 1, 2
go可以对多个变量进行操作

短变量声明
在函数中，简洁的赋值语句 := 可在类型明确的地方代替var声明
函数外的每个语句都必须以关键字开始（var func等等），因此 := 结构不能在函数外使用
flag, num, str := true, 1, "no!"  //合法，他可以自己去看

类型转换
Go是一种强类型转换的语言 ,开始定了什么类型就是什么类型了，如果想换必须使用类型转换
i := 42
f := float64(i)
u := uint(f)

i := 1
经常会改这个i值，那这个就不能写成 i := 2
而是应该写成 i = 2, 不能多次声明

数组定义
var identifier [len] type

myArray := [3]int{1,2,3}

切片slice
切片是对数组一个连续片段的引用
切片在未初始化之前默认为nil,长度为0

Make 和 New
New返回指针地址
Make返回第一个元素，可预设内存空间，避免未来的内存拷贝

mySlice1 := new([]int)
mySlice1 := make([]int, 10)

关于切片的常见问题
1.切片是连续内存并可以动态拓展的，因为这个特性会出现问题
a := []int
b := []int{1, 2, 3}
c := a
a = append(b, 1)

这里的问题是，如果b后面没空间了，Go就要再找一块，地址就变了，此时c和a就不一样了，
一般的解决办法是用谁赋给谁，我想让c添加元素就直接用c

2.修改切片的值？
mySlice :=[]int {10, 20, 30, 40}
for _, value := range mySlice{
	value *= 2
}

?这里的mySlice中的value没有变，这个value是临时创建出来的，Go和Java一样都是值传递
真的想变mySlice中的值
for index, _ := range mySlice{
	mySlice[index] *= 2
}



fmt这个包的作用
1.打印
可以多输出
fmt.Println("aa", "bb", "cc")
fmt.Printf("%s%d\n", "a", 1)
%v ->strut
%+v  既打印值也打印类型

fmt.Sprintf("")输出字串，可以
2.error
fmt.Errorf()

包管理：
go只用最后一个包名就可以了


Map
声明方法：
	var map1 map[keytype]valuetype
myMap := make(map[string]string, 10)
myMap["a"] = "b"
myFuncMap := map[string]func() int{
	"funcA":func() int{return 1}
}
fmt.Println(myFuncMap)
f := myFuncMap["funcA"]
fmt.Println(f())



*/
