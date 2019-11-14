package main

import (
	"fmt"
	"regexp"
	"runtime"
	"runtime/debug"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(CurrentGoRoutineID())
	fmt.Println(GoId())
	fmt.Println(GoID())
	fmt.Println(GetGoid())
}

// 现在流行用这个, 但学习了Plan9汇编之后觉得这个不是最强的
func CurrentGoRoutineID() string {
	bytes := debug.Stack()
	for i, ch := range bytes {
		if ch == '\n' || ch == '\r' {
			bytes = bytes[0:i]
			break
		}
	}
	line := string(bytes)
	var valid = regexp.MustCompile(`goroutine\s(\d+)\s+\[`)

	if params := valid.FindAllStringSubmatch(line, -1); params != nil {
		return params[0][1]
	} else {
		return ""
	}
}

func GoId() string {
	var buf = make([]byte, 64)
	var stk = buf[:runtime.Stack(buf, false)]
	return string(stk)
}

func GetGoid() int64 {
	var (
		buf [64]byte
		n   = runtime.Stack(buf[:], false)
		stk = strings.TrimPrefix(string(buf[:n]), "goroutine ")
	)

	idField := strings.Fields(stk)[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Errorf("can not get goroutine id: %v", err))
	}

	return int64(id)
}

func GoID() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}
