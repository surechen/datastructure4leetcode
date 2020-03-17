package bigmath4l

import (
	"fmt"
	"math"
	"math/big"
)

func BigMathFunc() {
	s := big.NewInt(99)
	k := big.NewInt(math.MaxInt64)
	fmt.Println(math.MaxInt64)

	m := s.Add(s, k)
	fmt.Println(m)
	fmt.Println(s)

	fmt.Println(s.Mul(s, k))

}
