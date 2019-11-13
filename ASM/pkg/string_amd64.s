
// 初始化 字符串
// 字符串其实是 只读引用，而不是常亮， 在go中对应的类型如下:
// 在Go汇编语言中，go.string."gopher"不是一个合法的符号，因此我们无法通过手工创建（这是给编译器保留的部分特权，因为手工创建类似符号可能打破编译器输出代码的某些规则）
//
// type reflect.StringHeader struct {
//     Data uintptr
//     Len  int
// }
//
// 注意： 如果这里谢忱过这样就会报错 `GLOBL ·NameData(SB),$8` 如下：
//  missing Go type information for global symbol: size 8
// 而 汇编中根本没有具体的类型可言，真正的原因是： 当Go语言的垃圾回收器在扫描到NameData变量的时候，无法知晓该变量内部是否包含指针
// 所以我们应该携程下面所示， 加上 `NOPTR` 表示其中不会包含指针数据
//
// 通过给·NameData增加NOPTR标志的方式表示其中不含指针数据。
// 或者在Go文件中我们也可以通过给·NameData变量在Go语言中增加一个不含指针并且大小为8个字节的类型来修改该错误：
//
// var NameData [8]byte
// var Name string
//
// 注意： `#include "textflag.h"` 这个必须加上哦
// 这个是字符串数据
#include "textflag.h"
GLOBL ·NameData(SB),NOPTR,$8
DATA  ·NameData(SB)/8,$"Gavinss"

// 这个是定义 字符串头结构体
GLOBL ·Name(SB),$16
DATA  ·Name+0(SB)/8,$·NameData(SB)
DATA  ·Name+8(SB)/8,$6
