package linknameC

import _ "unsafe"

//// 这个会报错， 因为 my-study/linknameB.hello 函数已经在自己的签名上表明了
//// 它将自己导出给 my-study/linknameA.Hello
//// 所以在 c中不能再使用该函数 导入了
////go:linkname Hello my-study/linknameB.hello
//func Hello() string

//go:linkname IsSpace fmt.isSpace
func IsSpace(r rune) bool

//go:linkname Say my-study/linknameB.say
func Say()
