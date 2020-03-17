/*
	chenshuo test
*/

package unicode4l

import (
	"fmt"
	"unicode"
)

func Unicodefunc() {
	str := "hello wor\nld 123\r2"
	strrune := []rune(str)
	fmt.Println(strrune)
	for i, v := range strrune {
		if unicode.IsDigit(strrune[i]) {
			fmt.Println("digit:", i, v)
		}
		if unicode.IsLetter(strrune[i]) {
			fmt.Println("letter:", i, v)
			fmt.Println("upletter", unicode.ToUpper(v))
		}
		if unicode.IsSpace(strrune[i]) {
			fmt.Println("space:", i, v)
		}
	}
	fmt.Println((string(strrune)))
}
