// 初始化
// 汇编定义变量时初始化数据并不区分整数是否有符号.
// 只有在CPU指令处理该寄存器数据时,才会根据指令的类型来取分数据的类型或者是否带有符号位.
GLOBL ·Int32Value(SB),$4
DATA ·Int32Value+0(SB)/1,$0x01  // 第0字节
DATA ·Int32Value+1(SB)/1,$0x02  // 第1字节
DATA ·Int32Value+2(SB)/2,$0x03  // 第3-4字节

GLOBL ·Uint32Value(SB),$4
DATA ·Uint32Value(SB)/4,$0x01020304 // 第1-4字节
