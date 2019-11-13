// 该，文件名可以随便定义，只要汇编中的 变量和 go文件中的变量同名即可
//
// GLOBL 表示将变量 `Id` 导出
// 这里表示 Id 占用 8 个byte
// 而每个byte 上的值，由下面 DATA 中定义
GLOBL ·Id(SB),$8

// 下面是初始化各个 byte 上的初始值
// 注意，光标必须所在最后一行内容的下一行开始处， 不然会报 `EOF`
// DATA 后面的关键字之前的变量名的 `.` 表示该变量为当前包中Go语言定义的符号symbol
DATA ·Id+0(SB)/1,$0x37
DATA ·Id+1(SB)/1,$0x25
DATA ·Id+2(SB)/1,$0x00
DATA ·Id+3(SB)/1,$0x00
DATA ·Id+4(SB)/1,$0x00
DATA ·Id+5(SB)/1,$0x00
DATA ·Id+6(SB)/1,$0x00
DATA ·Id+7(SB)/1,$0x00
