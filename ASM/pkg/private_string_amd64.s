
// amd64环境中StringHeader有16个字节大小，
// 因此我们先在Go代码声明字符串变量，然后在汇编中定义一个16字节大小的变量
// 定义了一个text当前文件内的私有变量（以<>为后缀名），内容为“Hello World!”

// text<>私有变量表示的字符串(真实)只有12个字符长度，
// 但是我们依然需要将变量的长度扩展为2的指数倍数，
// 这里也就是16个字节的长度。其中NOPTR表示text<>不包含指针数据


#include "textflag.h" // 貌似 导出定义 中有 NOPTR 则就需要有这一行
GLOBL text<>(SB),NOPTR,$16
DATA text<>+0(SB)/8,$"Hello Wo"
DATA text<>+8(SB)/8,$"rld!"


GLOBL ·Helloworld(SB),$16
// 然后使用text私有变量对应的内存地址对应的常量来初始化字符串头结构体中的Data部分，
// 并且手工指定Len部分为字符串的长度
DATA ·Helloworld+0(SB)/8,$text<>(SB) // StringHeader.Data
DATA ·Helloworld+8(SB)/8,$12         // StringHeader.Len
