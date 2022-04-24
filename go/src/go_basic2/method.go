/*
在Go语言中除了函数还有方法method
那他和function有什么区别呢---mathod有一个receiver
func functionname()
方法是
func (recv receiver_type) methodName(parameter_list) (return_value_list)
你定义一个实体 不管是接口还是数据结构，你可以为这个接口和数据结构定义一个方法
定义方法：把你的接口或者是数据结构放在receiver的位置上

使用时的区别，你想使用你定义的函数是通过包名加函数的
使用场景
很多场景下，函数需要的上下文可以保存在receiver属性中，

*/