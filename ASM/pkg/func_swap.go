package pkg

//go:nosplit  表示 跳过栈溢出检测
func Swap(a, b int) (int, int)
