package main

import (
	"example/contact"
	// "example/hello"
	"fmt"
)

func main() {
	// hello.Hello()
	fmt.Println("Hello, World!")
	fmt.Println("1、添加联系人 2、删除联系人 3、修改联系人 4、查询联系人 5、退出程序\n请输入你想使用的功能编号按Enter键确认")
	var opCode int
	fmt.Scanf("%d", &opCode)
	switch opCode {
	case 1:
		// fmt.Println("添加联系人")
	case 2:
		// contact.DeleteContacts()
	case 3:
		// fmt.Println("修改联系人")
	case 4:
		contact.SearchContacts()
	case 5:
		// fmt.Println("退出程序")
	}
}
