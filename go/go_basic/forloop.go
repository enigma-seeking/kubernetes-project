package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}
	fullString := "hello world"
	fmt.Println(fullString)
	for i, c := range fullString {
		fmt.Println(i, string(c))
	}

}

// ! 循环，Go只有一种循环：for循环
// for 初始化; 条件语句; 修饰语句｛｝
// for i := 0; i < 10; i++{

// }

// 初始化语句和后置语句是可选的，此场景与while等价
// for ; sum < 1000;{
// 	sum += sum;
// }

// 无限循环
// for{
// 	if condition1{
// 		break;
// 	}
// }

// for-range
// 遍历数组，切片，字符串，Map等
// for index, char := range myString{

// }

// for key, value := range myMap{

// }

// for index, value := range MyArray{

// }
// 需要注意的是：如果for range遍历指针数组，则value取出的指针地址为原指针地址的拷贝
