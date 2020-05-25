#include "textflag.h"
TEXT ·If(SB), NOSPLIT, $0-32
    MOVQ ok+8*0(FP), CX // ok
    MOVQ a+8*1(FP), AX  // a
    MOVQ b+8*2(FP), BX  // b

    CMPQ CX, $0         // test ok  使用CMPQ比较指令将CX寄存器和常数0进行比较
    JZ   L              // if ok == 0, goto L   比较的结果为0，那么下一条JZ为0时跳转指令将跳转到L标号对应的语句，也就是返回变量b的值。
    MOVQ AX, ret+24(FP) // return a    如果比较的结果不为0，那么JZ指令将没有效果，继续执行后面的指令，也就是返回变量a的值。
    RET

L:
    MOVQ BX, ret+24(FP) // return b
    RET


// 在跳转指令中，跳转的目标一般是通过一个标号表示。
// 不过在有些通过宏实现的函数中，更希望通过相对位置跳转，
// 这时候可以通过PC寄存器的偏移量来计算临近跳转的位置。