//Package get-input 用fmt包获取输入的字符或数字,字符不能有空格
package input_by_fmt

import (
	"fmt"
	"strconv"
)

var str string

//GetString 获取字符串
func GetString(input string) string {
	n := 1
	for n == 1 {
		fmt.Println(input)
		n, err := fmt.Scanln(&str)
		fmt.Println(n, str)
		if err == nil {
			break
		}

	}
	return str
}

//GetInt 获取整数
func GetInt(input string) int {
	var interger int
	for {
		fmt.Println(input)
		_, err := fmt.Scanln(&str)
		if err == nil {
			interger, err = strconv.Atoi(str)
			if err == nil {
				break
			}
			println("请输入数字")
		}
	}
	return interger
}
