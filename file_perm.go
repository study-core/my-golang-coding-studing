package main

import (
	"fmt"
	"os"
)

func main() {


	fmt.Println("ModeDir:", os.ModeDir)
	fmt.Println("ModeAppend:", os.ModeAppend)
	fmt.Println("ModeExclusive:", os.ModeExclusive)
	fmt.Println("ModeTemporary:", os.ModeTemporary)
	fmt.Println("ModeDevice:", os.ModeDevice)
	fmt.Println("ModeNamedPipe:", os.ModeNamedPipe)
	fmt.Println("ModeSocket:", os.ModeSocket)
	fmt.Println("ModeSetuid:", os.ModeSetuid)
	fmt.Println("ModeSetgid:", os.ModeSetgid)
	fmt.Println("ModeCharDevice:", os.ModeCharDevice)
	fmt.Println("ModeSticky:", os.ModeSticky)
	fmt.Println("ModeIrregular:", os.ModeIrregular)
	fmt.Println("ModeType:", os.ModeType)
	fmt.Println("ModePerm:", os.ModePerm)

}
