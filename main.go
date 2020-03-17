/*
	chenshuo test
*/

package main

import (
	"./bigmath4l"
	"./heap4l"
	"./list4l"
	"./math4l"
	"./sort4l"
	"./stack4l"
	"./string4l"
	"./unicode4l"
)

func testPacString() {
	string4l.Stringfunc()
}

func testPacUnicode() {
	unicode4l.Unicodefunc()
}

func testPacSort() {
	sort4l.Sort4lFunc()
}

func testBigMathFunc() {
	bigmath4l.BigMathFunc()
}

func testHeapFunc() {
	heap4l.HeapFunc()
}

func testListFunc() {
	list4l.ListFunc()
}

func testMathFunc() {
	math4l.MathFunc()
}

func testStackFunc() {
	stack4l.StackFunc()
}

func main() {
	//testPacString()
	//testPacUnicode()
	testPacSort()
	//testBigMathFunc()
	//testHeapFunc()
	//testListFunc()
	//testMathFunc()
	//testStackFunc()
}
