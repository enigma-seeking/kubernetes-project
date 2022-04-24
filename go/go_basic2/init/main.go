/*
init函数：会在包初始化时运行
谨慎使用init函数
	当多个依赖项目引用统一项目，且被引用项目的初始化在init中完成，并且不可重负运行时，会导致启动错误

kubernetes->glog->init->flag parse parameter
依赖glog。glog先运行他的init，而init用flag parse parameter接受参数
正常下没问题

问题是kubernetes->a->vendor->glog->init->flag parse parameter
引用了项目a，a的vendor目录，引用了自己的glog，glog存在多个地方，kubernets不认为是一个包，所以又做init
很可能项目跑不起来了
现在管理是a如果还有glog就用最顶的glog。现在放弃glog了用的自己的klog

*/
package main

import (
	"fmt"
	// _ "github.com/cncamp/golang/examples/init/a"
	// _ "github.com/cncamp/golang/examples/init/b"
)

func init() {
	fmt.Println("main init")
} //init是在所有函数之前执行的

func main() {

}

// 打印顺序是 b-a-main
