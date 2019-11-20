
//
//  汇编实现这个 函数
//  func LoopAdd(cnt, v0, step int) int {
//      result := v0
//      for i := 0; i < cnt; i++ {
//          result += step
//      }
//      return result
//  }
//
//
//  先转成 汇编的思想：
//
//
//  func LoopAdd(cnt, v0, step int) int {
//      var i = 0
//      var result = 0
//
//  LOOP_BEGIN:
//      result = v0
//
//  LOOP_IF:
//      if i < cnt { goto LOOP_BODY }
//      goto LOOP_END
//
//  LOOP_BODY
//      i = i+1
//      result = result + step
//      goto LOOP_IF
//
//  LOOP_END:
//
//      return result
//  }
//
//
//
//
//
//
//
//



#include "textflag.h"

// func LoopAdd(cnt, v0, step int) int
TEXT ·LoopAdd(SB), NOSPLIT,  $0-32
    MOVQ cnt+0(FP), AX   // cnt
    MOVQ v0+8(FP), BX    // v0/result
    MOVQ step+16(FP), CX // step

LOOP_BEGIN:
    MOVQ $0, DX          // i

LOOP_IF:
    CMPQ DX, AX          // compare i, cnt
    JL   LOOP_BODY       // if i < cnt: goto LOOP_BODY
    JMP LOOP_END

LOOP_BODY:
    ADDQ $1, DX          // i++
    ADDQ CX, BX          // result += step
    JMP LOOP_IF

LOOP_END:

    MOVQ BX, ret+24(FP)  // return result
    RET
