
// 需要加这个文件，不然c.go 中的没有方法体的函数就会编译报错
// 可以知道 在哪个 x.go文件有那种 无body 的函数，则需要对应的加上 x.s 文件
// 其中 x.go 和 x.s 的x可以不同名， 如: c.go 和 f.s, 只要f.s是个“空的”文件即可