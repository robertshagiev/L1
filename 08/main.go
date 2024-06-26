package main

import (
	"fmt"
)

func I1(n int64, i int) int64 { // Для того, чтобы установить i-й бит в 1, нужно использовать побитовое ИЛИ, при этом нижний операнд на i бите должен быть равен 1
	return n | 1<<(i-1)
}

func I0(n int64, i int) int64 { // Для того, чтобы установить i-й бит в 0, нужно использовать побитовое исключающее ИЛИ, при этом нижний операнд на i бите должен быть равен 1
	return n ^ (1 << (i - 1))
}

func main() {

	i1 := I1(0, 1)
	i0 := I0(-1, 64)

	fmt.Println(i1)
	fmt.Println(i0)

}
