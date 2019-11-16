package pkg

func If(ok int, a, b int) int

/**
想达到的效果为：

func If(ok bool, a, b int) int {
    if ok { return a } else { return b }
}

分解成汇编的思想为： 0 == false  true == 1

func If(ok int, a, b int) int {
    if ok == 0 { goto L }
    return a
L:
    return b
}

*/
