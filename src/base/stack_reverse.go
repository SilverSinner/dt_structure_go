//实现栈的所有元素翻转
package main

import (
	st "../structs/stack"
)

func main() {
	stack := &st.SliceStack{Slice: make([]int, 0)}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	stack.PrintStack()
}
