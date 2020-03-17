package math4l

import (
	"fmt"
	"math"
)

func MathFunc() {
	fmt.Println(math.Sqrt(9.32))
	fmt.Println(math.Hypot(5, 6))
	fmt.Println(math.Max(9.2, 2.7))
	fmt.Println(math.Min(9.2, 2.7))
	fmt.Println(math.MaxInt32)
	fmt.Println(math.MaxFloat32)
	fmt.Println(math.Ceil(9.4))   //向上取整
	fmt.Println(math.Floor(9.78)) //向下取整
	fmt.Println(math.Pow(9, 4))
	fmt.Println(math.Log10(100))               //以10为底
	fmt.Println(math.Log(math.Pow(math.E, 4))) //自然对数
	fmt.Println(math.Abs(-9.24))               //自然对数
}
