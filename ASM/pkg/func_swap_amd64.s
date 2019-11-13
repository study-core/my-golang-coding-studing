// 下面两种方式
// 目前可能遇到的函数标志有NOSPLIT、WRAPPER和NEEDCTXT几个。
// 【NOSPLIT】不会生成或包含栈分裂代码，这一般用于没有任何其它函数调用的叶子函数，这样可以适当提高性能。
// 【WRAPPER】标志则表示这个是一个包装函数，在panic或runtime.caller等某些处理函数帧的地方不会增加函数帧计数。
// 【NEEDCTXT】表示需要一个上下文参数，一般用于闭包函数。



//
//  需要注意的是函数也没有类型，这里定义的Swap函数签名可以下面任意一种格式：
//
//  func Swap(a, b, c int) int
//  func Swap(a, b, c, d int)
//  func Swap() (a, b, c, d int)
//  func Swap() (a []int, d int)
//  // ...
//
//  注意： 对于汇编函数来说，只要是函数的名字和参数大小一致就可以是相同的函数了。
//  而且在Go汇编语言中，输入参数和返回值参数是没有任何的区别的。



// func Swap(a, b int) (int, int)
//
// 这是 完整写法， 函数名部分包含了当前包的路径，同时指明了函数的参数大小为32个字节（对应参数和返回值的4个int类型）。
TEXT ·Swap(SB), NOSPLIT, $0-32

// func Swap(a, b int) (int, int)
//
// 这是 简洁写法， 省略了当前包的路径和参数的大小。如果有NOSPLIT标注，会禁止汇编器为汇编函数插入栈分裂的代码 (即： 溢出时，做分裂栈拓展)。NOSPLIT对应Go语言中的//go:nosplit注释。
// TEXT ·Swap(SB), NOSPLIT, $0