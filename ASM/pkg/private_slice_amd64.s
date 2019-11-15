
// slice变量和string变量相似，只不过是对应的是切片头结构体而已
//
// type reflect.SliceHeader struct {
//     Data uintptr
//     Len  int
//     Cap  int
// }
//

// 切片的头的前2个成员字符串是一样的。
// 因此我们可以在前面字符串变量的基础上，
// 再扩展一个Cap成员就成了切片类型了：

//
// 注意: 因为切片和字符串的相容性，我们可以将切片头的前16个字节临时作为字符串使用，这样可以省去不必要的转换

GLOBL ·HelloGavin(SB),$24            // var helloworld []byte("Hello World!")
DATA ·HelloGavin+0(SB)/8,$text2<>(SB) // StringHeader.Data
DATA ·HelloGavin+8(SB)/8,$12         // StringHeader.Len
DATA ·HelloGavin+16(SB)/8,$16        // StringHeader.Cap

#include "textflag.h"
GLOBL text2<>(SB), NOPTR, $16
DATA text2<>+0(SB)/8,$"Hello Wo"      // ...string data...
DATA text2<>+8(SB)/8,$"rld!"          // ...string data...
