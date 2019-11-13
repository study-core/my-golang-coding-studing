// 第二中 字符串的定义
// 消除 Name2Data字段
// 将底层的字符串数据和字符串头结构体定义在一起，这样可以避免引入NameData符号：
GLOBL ·Name2(SB),$24

DATA ·Name2+0(SB)/8,$·Name2+16(SB)
DATA ·Name2+8(SB)/8,$6
DATA ·Name2+16(SB)/8,$"GKssddd"
