package pkg

// 使用Plan9汇编实现的自定义 互换函数
//go:nosplit  表示 跳过栈溢出检测
func GavinSwap(a, b int) (int, int)

func GavinSwap2(a, b int) (int, int)
