/*
	chenshuo test
*/

package string4l

import (
	"fmt"
	"strings"
)

func Stringfunc() {
	str := "232 sda hels 2ddd safaf sda"
	arr := strings.Split(str, " ")
	fmt.Println(arr)
	fmt.Println(strings.Contains(str, "hels"))
	fmt.Println(strings.ToUpper(str))
	fmt.Println(strings.TrimPrefix(str, "232"))
	fmt.Println(strings.TrimRight(str, "safaf"))
	fmt.Println(strings.TrimSpace("  sad sdawww    s s s      \r \n")) //只删除最前和最后面的空格 换行符号
	fmt.Println(strings.TrimLeft(str, "232 "))
	fmt.Println(strings.TrimLeft(str, "sda"))
	fmt.Println(strings.Index(str, "sda"), strings.IndexAny(str, "sda"), strings.LastIndex(str, "sda"))
	fmt.Println(strings.Replace(str, "sda", "tihuan", -1))
}
