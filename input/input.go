//Package get-input 使用bufio包获取输入,可以输入空格
package input

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var str string

//GetString 获取字符串
func GetString(input string) string {
	// for {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(input)
	data, _, _ := reader.ReadLine()
	return string(data)
	// }

}

//GetInt 获取整数
func GetInt(input string) int {
	var interger int
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println(input)
		data, _, _ := reader.ReadLine()
		inter, err := strconv.Atoi(string(data))
		if err == nil {
			interger = inter
			break
		}
	}
	return interger
}
