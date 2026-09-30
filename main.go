package main

import (
	"example/contact"

	"github.com/eiannone/keyboard"

	// "example/hello"
	"fmt"
	"os"
)

// var name string

// func DeleteContactsbyName(name string) {

// }

func main() {
	// reader := bufio.NewReader(os.Stdin)
	// char, _ := reader.ReadByte()
	// hello.Hello()
	keyboard.Open()
	defer keyboard.Close()
	fmt.Println("Hello, World!")
	for {
		fmt.Println("1、添加联系人 2、删除联系人 3、修改联系人 4、查询联系人 5、退出程序\n请输入你想使用的功能编号按Enter键确认，按Esc键返回至主菜单")
		// opCodeStr, _ := reader.ReadString('\n')
		// opCodeStr = strings.TrimSpace(opCodeStr)
		opCodeStr, _ := contact.ReadTrimmedInput()
		// opCode, _ := strconv.Atoi(opCodeStr)
		switch opCodeStr {
		// case "ESC":
		// 	continue
		case "1":
			fmt.Println("添加联系人")
			fmt.Println("请输入联系人姓名：")
			name, Err := contact.ReadTrimmedInput()
			if Err != false {
				fmt.Println("请输入联系人电话号码：")
				phone, Err := contact.ReadTrimmedInput()
				if Err != false {
					contact.AddContacts(name, phone)
					contact.ShowContacts()
				}
			}
			continue
		case "2":
			fmt.Println("请输入你想删除的联系人姓名：")
			name, Err := contact.ReadTrimmedInput()
			if Err != false {
				if contact.DeleteContacts(name) {
					fmt.Println("删除成功！")
				}
				contact.ShowContacts()
			}
			continue
		case "3":
			fmt.Println("请输入你想修改的联系人姓名：")
			name, Err := contact.ReadTrimmedInput()
			if Err != false {
				contact.UpdateContacts(name)
				continue
			}
		case "4":
			contact.ShowContacts()
		case "5":
			os.Exit(0)
		}
	}
}
