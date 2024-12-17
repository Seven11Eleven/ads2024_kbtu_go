package main

import (
	"fmt"
	"github.com/Seven11Eleven/ads2024_kbtu_go/lecture1"
)

func main() {
	stakke := lecture1.NewStack()
	stakkelan := new(lecture1.Stack)
	stakke.Push("1")
	stakke.Push("1337")
	fmt.Println(stakke.Peek())
	stakke.Pop()
	fmt.Println(stakke.Peek())
	fmt.Println("pornography is death")
	stakkelan.Push(1)
	stakkelan.Push(1337)
	fmt.Println(stakkelan.Peek())
	stakkelan.Pop()
	fmt.Println(stakkelan.Peek())
}
