package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

/*
	读取文件一共有三种方法: 直接通过os来读取、通过bufio来读取、通过ioutil读取
*/

// 通过os直接来读取
func readByOS(filePathAb string) {
	var result []byte
	buffer := make([]byte, 16)

	file, err := os.Open(filePathAb)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			return
		}
		result = append(result, buffer[:n]...)
	}

	fmt.Println(string(result))
}

// 通过buffer来进行读取文件
// 1.打开文件 os.Open()
// 2.创建bufio.NewReader(file)对象
// 3.通过ReadString('\n')函数来界定分隔符，这里使用的是\n，每次读取一行
func readByBuffer(filePathAb string) {
	var data string
	file, err := os.Open(filePathAb)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err1 := reader.ReadString('\n')

		// 由于空行的存在，所以即使遇到终止符，也有可能还会有文件流输出，所以需要在break前面再次读取一次文本流
		if err1 == io.EOF {
			data += line
			break
		}
		if err1 != nil {
			fmt.Println(err)
			return
		}
		data += line
	}
	fmt.Println(data)
}

func readByUtil(filePathAb string) {
	file, err := os.Open(filePathAb)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	io.ReadAll(file)
}

func main() {
	readByOS("./main.go")
	readByBuffer("./main.go")
}
