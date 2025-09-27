package main

import "fmt"

// 定义一个结构体用于演示
type User struct {
	ID   int
	Name string
}

func main() {
	// 创建一个结构体实例
	user := User{ID: 101, Name: "Alice"}

	// 创建一个简单的字符串
	message := "success"

	fmt.Println("--- 演示开始 ---")

	// 1. 使用 %v (默认格式)
	// 对于结构体，它会输出用花括号包裹的值
	// 对于字符串，它会输出字符串内容
	fmt.Printf("%%v (默认格式)    -> 结构体: %v, 字符串: %v\n", user, message)

	// 2. 使用 %+v (带字段名的结构体)
	// 对于结构体，它会额外输出每个字段的名称
	// 对于字符串，效果和 %v 一样
	fmt.Printf("%%+v (带字段名)   -> 结构体: %+v, 字符串: %+v\n", user, message)

	// 3. 使用 %#v (Go 语法表示)
	// 对于结构体，它会输出一个完整的、可以在代码中使用的结构体实例语法
	// 对于字符串，它会加上双引号
	fmt.Printf("%%#v (Go 语法表示) -> 结构体: %#v, 字符串: %#v\n", user, message)

	// 4. 使用 %T (值的类型)
	// 输出变量的类型
	fmt.Printf("%%T (类型)          -> 结构体类型: %T, 字符串类型: %T\n", user, message)

	// 5. 使用 %% (输出百分号)
	// 用于在格式化字符串中输出一个 '%' 字符本身
	progress := 99.9
	fmt.Printf("当前进度是: %.1f%%\n", progress)

	fmt.Println("--- 演示结束 ---")
}