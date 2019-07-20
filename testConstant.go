package main

import "fmt"

func main() {
	//fmt.Println(CALL, CALL2, CALL3, SEND, SEND2, SEND3)
	fmt.Println(STOP,
	ADD,
	MUL,
	SUB,
	DIV,
	SDIV,
	MOD,
	SMOD,
	ADDMOD,
	MULMOD,
	EXP,
	SIGNEXTEND,
	LT ,
	GT,
	SLT,
	SGT,
	EQ,
	ISZERO,
	AND,
	OR,
	XOR,
	NOT,
	BYTE,
	SHL,
	SHR,
	SAR,
	SHA3)


	fmt.Println(PUSH1,
	PUSH2,
	PUSH3,
	PUSH4,
	PUSH5,
	PUSH6,
	PUSH7,
	PUSH8,
	PUSH9,
	PUSH10,
	PUSH11,
	PUSH12,
	PUSH13,
	PUSH14,
	PUSH15,
	PUSH16,
	PUSH17,
	PUSH18,
	PUSH19,
	PUSH20,
	PUSH21,
	PUSH22,
	PUSH23,
	PUSH24,
	PUSH25,
	PUSH26,
	PUSH27,
	PUSH28,
	PUSH29,
	PUSH30,
	PUSH31,
	PUSH32,
	DUP1,
	DUP2,
	DUP3,
	DUP4,
	DUP5,
	DUP6,
	DUP7,
	DUP8,
	DUP9,
	DUP10,
	DUP11,
	DUP12,
	DUP13,
	DUP14,
	DUP15,
	DUP16,
	SWAP1,
	SWAP2,
	SWAP3,
	SWAP4,
	SWAP5,
	SWAP6,
	SWAP7,
	SWAP8,
	SWAP9,
	SWAP10,
	SWAP11,
	SWAP12,
	SWAP13,
	SWAP14,
	SWAP15,
	SWAP16)
}

const (
	CALL = iota
	CALL2
	CALL3

	/*SEND = iota
	SEND2
	SEND3*/
)

const (
	SEND = iota
	SEND2
	SEND3
)


const (
	STOP  = iota 	// 0 == 0x00
	ADD					// 1
	MUL					// 2
	SUB					// 3
	DIV					// 4
	SDIV				// 5
	MOD					// 6
	SMOD				// 7
	ADDMOD				// 8
	MULMOD				// 9
	EXP					// 10
	SIGNEXTEND			// 11

)

const (
	LT  = iota + 0x10
	GT
	SLT
	SGT
	EQ
	ISZERO
	AND
	OR
	XOR
	NOT
	BYTE
	SHL
	SHR
	SAR
	SHA3 = 0x20
)

const (
	PUSH1     = 0x60 + iota
	PUSH2
	PUSH3
	PUSH4
	PUSH5
	PUSH6
	PUSH7
	PUSH8
	PUSH9
	PUSH10
	PUSH11
	PUSH12
	PUSH13
	PUSH14
	PUSH15
	PUSH16
	PUSH17
	PUSH18
	PUSH19
	PUSH20
	PUSH21
	PUSH22
	PUSH23
	PUSH24
	PUSH25
	PUSH26
	PUSH27
	PUSH28
	PUSH29
	PUSH30
	PUSH31
	PUSH32
	DUP1
	DUP2
	DUP3
	DUP4
	DUP5
	DUP6
	DUP7
	DUP8
	DUP9
	DUP10
	DUP11
	DUP12
	DUP13
	DUP14
	DUP15
	DUP16
	SWAP1
	SWAP2
	SWAP3
	SWAP4
	SWAP5
	SWAP6
	SWAP7
	SWAP8
	SWAP9
	SWAP10
	SWAP11
	SWAP12
	SWAP13
	SWAP14
	SWAP15
	SWAP16
)