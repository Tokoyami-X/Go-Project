package contact

import (
	"fmt"
	"strings"

	"github.com/eiannone/keyboard"
)

type Contact struct {
	Name  string
	Phone string
}

func findContactIndex(name string) int {
	for i, contact := range contacts {
		if contact.Name == name {
			return i
		}
	}
	return -1
}

func ReadTrimmedInput() (string, bool) {
	var input []rune
	for {
		char, code, err := keyboard.GetKey()
		if err != nil {
			return "", false
		}
		if code == keyboard.KeyEsc {
			fmt.Println()
			return "ESC", false
		}
		if code == keyboard.KeyEnter {
			fmt.Println()
			return strings.TrimSpace(string(input)), true
		}
		if code == keyboard.KeyBackspace || code == keyboard.KeyBackspace2 {
			if len(input) > 0 {
				input = input[:len(input)-1]
				// \r 回到当前行首，\033[K 清掉光标到行尾的所有内容
				// 然后重新打印当前 input（此时最后一个字符已从数据层删掉）
				// 这种写法不依赖字符宽度，中文/emoji/ASCII 都能正确处理
				fmt.Printf("\r\033[K%s", string(input))
			}
			continue
		}
		// 只打印当前键入的字符，而不是整个 input 切片
		input = append(input, char)
		fmt.Print(string(char))
	}
}

// findContactIndex 按姓名查找联系人，返回下标；未找到返回 -1
// func findContactIndex(name string) int {
// 	for i, c := range contacts {
// 		if c.Name == name {
// 			return i
// 		}
// 	}
// 	return -1
// }

// setName 方法用于为 Struct 类型的实例设置名称
// 它接收一个字符串类型的参数 name，并将该值赋给结构体的 Name 字段
// func setName(name string) *Contact {
// 将传入的 name 参数赋值给结构体 c 的 Name 属性
// return &Contact{Name: name}
// }

//	func (c *Struct) setPhone(phone string) {
//		c.Phone = phone
//	}
var contacts = []Contact{
	{"Alice", "123-456-7890"},   // 第一个联系人：Alice，电话号码：123-456-7890
	{"Bob", "987-654-3210"},     // 第二个联系人：Bob，电话号码：987-654-3210
	{"Charlie", "555-555-5555"}, // 第三个联系人：Charlie，电话号码：555-555-5555
}

// main 函数是程序的入口点
/*
这是一个主函数，用于演示如何创建和遍历联系人切片
每个联系人包含姓名和电话号码两个属性
程序会打印出所有联系人的姓名和电话号码
*/
// 一、ShowContact 函数用于展示联系人列表
// 该函数创建一个联系人切片，并遍历打印每个联系人的姓名和电话号码
func ShowContacts() {
	for i, c := range contacts {
		fmt.Printf("第%d个联系人：%s, 电话号码：%s\n", i+1, c.Name, c.Phone)
	}

}

// 二、DeleteContacts 函数用于删除联系人列表
func DeleteContacts(name string) bool {
	// 校验参数：姓名不能为空
	if name == "" {
		fmt.Println("错误：要删除的联系人姓名不能为空！")
		return false
	}
	// 创建一个新切片用于存储删除后的联系人
	// finalContacts := make([]Contact, 0, len(contacts))
	// 记录是否找到并删除了匹配的联系人
	// found := false
	// 遍历 contacts 切片，将姓名不等于 name 的联系人加入新切片
	idx := findContactIndex(name)
	if idx != -1 {
		contacts = append(contacts[:idx], contacts[idx+1:]...)
		fmt.Printf("已成功删除姓名为 %s 的联系人\n", name)
		return true
	}
	fmt.Printf("未找到姓名为 %s 的联系人\n", name)
	return false
	// for _, contact := range contacts {
	// 	if contact.Name != name {
	// 		finalContacts = append(finalContacts, contact)
	// 	} else {
	// 		found = true
	// 	}
	// }
	// 如果没有删除任何联系人，说明未找到匹配项
	// if !found {
	// 	fmt.Printf("未找到姓名为 %s 的联系人\n", name)
	// 	return false
	// }
	// 将删除后的切片赋值回全局 contacts
	// contacts = finalContacts
	// fmt.Printf("已成功删除姓名为 %s 的联系人\n", name)
}

// DeleteContacts 删除指定姓名的联系人
// func DeleteContacts(name string) bool {
// 	if name == "" {
// 		fmt.Println("错误：要删除的联系人姓名不能为空！")
// 		return false
// 	}
// 	idx := findContactIndex(name)
// 	if idx == -1 {
// 		fmt.Printf("未找到姓名为 %s 的联系人\n", name)
// 		return false
// 	}
// 	// 原地拼接，不分配新切片
// 	contacts = append(contacts[:idx], contacts[idx+1:]...)
// 	fmt.Printf("已成功删除姓名为 %s 的联系人\n", name)
// 	return true
// }

// 三、AddContacts 函数用于添加联系人列表
func AddContacts(name, phone string) {
	// 校验参数：姓名和电话号码不能为空
	if name == "" || phone == "" {
		fmt.Println("错误：联系人姓名和电话号码不能为空！")
		return
	}
	// 创建一个新的联系人实例
	newContact := Contact{Name: name, Phone: phone}
	// 将新联系人添加到 contacts 切片中
	contacts = append(contacts, newContact)
	fmt.Printf("已成功添加联系人：%s, 电话号码：%s\n", name, phone)
}

// 四、UpdateContacts 函数用于修改联系人列表
func UpdateContacts(name string) bool {
	if name == "" {
		fmt.Println("错误：输入内容不能为空！")
		return true
	}
	idx := findContactIndex(name)
	if idx == -1 {
		fmt.Printf("未找到姓名为 %s 的联系人\n", name)
		return true
	}

	fmt.Println("查找成功！请输入想要修改的联系姓名:")
	newName, ok := ReadTrimmedInput()
	if !ok {
		return false
	}
	newName = strings.TrimSpace(newName)
	if newName == "" {
		fmt.Println("错误：输入内容不能为空！请重新输入")
		return false
	}

	fmt.Println("请输入想要修改的联系电话号码:")
	newPhone, ok := ReadTrimmedInput()
	if !ok {
		return false
	}
	newPhone = strings.TrimSpace(newPhone)
	if newPhone == "" {
		fmt.Println("错误：修改后的电话号码不能为空！请重新输入")
		return false
	}

	contacts[idx].Name = newName
	contacts[idx].Phone = newPhone
	fmt.Printf("已成功修改联系人姓名为：%s, 电话号码：%s\n", newName, newPhone)
	return true
	// reader := bufio.NewReader(os.Stdin)
	// if name == "" {
	// 	fmt.Println("错误：输入内容不能为空！")
	// 	return false
	// }
	// idx := findContactIndex(name)
	// if idx == -1 {
	// 	fmt.Printf("未找到姓名为 %s 的联系人\n", name)
	// 	return false
	// }
	// fmt.Println("查找成功！请输入想要修改的联系姓名:")
	// for {
	// 	newName, Err := ReadTrimmedInput()
	// 	if Err != false {
	// 		for {
	// 			if newName == "" {
	// 				fmt.Println("错误：输入内容不能为空！请重新输入")
	// 				break
	// 			}
	// 			fmt.Println("请输入想要修改的联系电话号码:")
	// 			for {
	// 				newPhone, Err := ReadTrimmedInput()
	// 				if Err != false {
	// 					for {
	// 						if newPhone == "" {
	// 							fmt.Println("错误：修改后的电话号码不能为空！请重新输入")
	// 							break
	// 						}
	// 						contacts[idx].Name = newName
	// 						contacts[idx].Phone = newPhone
	// 						fmt.Printf("已成功修改联系人姓名为：%s, 电话号码：%s\n", newName, newPhone)
	// 						return true
	// 					}
	// 				}
	// 				return false
	// 			}
	// 		}
	// 	}
	// 	return false
	//ai写的updateContacts代码

	// }
	// 这行只是为了让 Go 编译器消除 "function ends without a return statement" 报错
	// 实际上上面的无限 for 内部已经覆盖了所有返回路径，运行时永远不会走到这里
	// return false
}
