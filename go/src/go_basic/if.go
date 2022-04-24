// 条件：
/*
if condition1{

}else if condition2{

}else{

}
格式限制的很死，连括号的位置都是固定的。

也有简写法
if v := x - 100; v < 0{
	return v
}

switch var1{
	case val1: //空分只
	case val2:
		fallthrough // 执行case3的内容
	case val3:
		f()
	default:
}



*/
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
}

//如果运行的时候是./if  那么
// os.Args就是操作系统的参数，go语言的参数怎么来的呢，都是通过os.Args拿进来的，他不需要main来传
// 有了flag这个包，可以通过 --的方式传参数，flag可以捕捉
// 因为格式十分重要，package这句话应该是在第一行去写，放在第二行就报错
