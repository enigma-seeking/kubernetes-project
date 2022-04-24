package main

/*
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

*/
