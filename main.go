package main

import (
	"bufio"
	"example/contact"

	// "example/hello"
	"fmt"
	"os"
)

// var name string

// func DeleteContactsbyName(name string) {

// }

func main() {
	reader := bufio.NewReader(os.Stdin)
	// hello.Hello()
	fmt.Println("Hello, World!")
	for {
		fmt.Println("1、添加联系人 2、删除联系人 3、修改联系人 4、查询联系人 5、退出程序\n请输入你想使用的功能编号按Enter键确认")
		// opCodeStr, _ := reader.ReadString('\n')
		// opCodeStr = strings.TrimSpace(opCodeStr)
		opCodeStr := contact.ReadTrimmedInput(reader)
		// opCode, _ := strconv.Atoi(opCodeStr)
		switch opCodeStr {
		case "1":
			fmt.Println("添加联系人")
			fmt.Println("请输入联系人姓名：")
			name := contact.ReadTrimmedInput(reader)
			fmt.Println("请输入联系人电话号码：")
			phone := contact.ReadTrimmedInput(reader)
			contact.AddContacts(name, phone)
			contact.SearchContacts()
		case "2":
			fmt.Println("请输入你想删除的联系人姓名：")
			name := contact.ReadTrimmedInput(reader)
			if contact.DeleteContacts(name) {
				fmt.Println("删除成功！")
			}
			contact.SearchContacts()
		case "3":
			fmt.Println("请输入你想修改的联系人姓名：")
			name := contact.ReadTrimmedInput(reader)
			contact.UpdateContacts(name)
		case "4":
			contact.SearchContacts()
		case "5":
			os.Exit(0)
		}
	}
}
