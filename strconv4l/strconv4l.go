package strconv4l

import (
	"fmt"
	"strconv"
)

func StrconvFunc() {
	strInt := "-9732"
	fmt.Print("Atoi:")
	fmt.Println(strconv.Atoi(strInt))
	fmt.Print("ParseInt:")
	fmt.Println(strconv.ParseInt(strInt, 10, 0))

	intval := 342424
	fmt.Print("Itoa:")
	fmt.Println(strconv.Itoa(intval))
	fmt.Print("FormatInt:")
	fmt.Println(strconv.FormatInt(int64(intval), 10))

	strFloat := "-93.2322"
	fmt.Print("ParseFloat:")
	fmt.Println(strconv.ParseFloat(strFloat, 32))

	floatval := -93.2322522
	fmt.Print("FormatFloat:")
	//末尾四舍五入
	fmt.Println(strconv.FormatFloat(floatval, 'f', 4, 32))
}
