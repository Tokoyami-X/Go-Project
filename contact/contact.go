package contact

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Contact struct {
	Name  string
	Phone string
}

func ReadTrimmedInput(reader *bufio.Reader) string {
	codeStr, _ := reader.ReadString('\n')
	codeStr = strings.TrimSpace(codeStr)
	return codeStr
}

// setName 方法用于为 Struct 类型的实例设置名称
// 它接收一个字符串类型的参数 name，并将该值赋给结构体的 Name 字段
func setName(name string) *Contact {
	// 将传入的 name 参数赋值给结构体 c 的 Name 属性
	return &Contact{Name: name}
}

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
func SearchContacts() {
	// 创建一个联系人切片，包含三个联系人元素
	// 每个元素都是一个 Contact 结构体，包含 Name 和 Phone 字段
	// contacts = append(contacts, Contact{"David", "111-111-1111"})
	// 遍历 contacts 切片，使用 for-range 循环
	// 对于每个联系人 c，打印其姓名和电话号码
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
	finalContacts := make([]Contact, 0, len(contacts))
	// 记录是否找到并删除了匹配的联系人
	found := false
	// 遍历 contacts 切片，将姓名不等于 name 的联系人加入新切片
	for _, contact := range contacts {
		if contact.Name != name {
			finalContacts = append(finalContacts, contact)
		} else {
			found = true
		}
	}
	// 如果没有删除任何联系人，说明未找到匹配项
	if !found {
		fmt.Printf("未找到姓名为 %s 的联系人\n", name)
		return false
	}
	// 将删除后的切片赋值回全局 contacts
	contacts = finalContacts
	fmt.Printf("已成功删除姓名为 %s 的联系人\n", name)
	return true
}

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
func UpdateContacts(name string) {
	reader := bufio.NewReader(os.Stdin)
	if name == "" {
		fmt.Println("错误：输入内容不能为空！")
		return
	}
	for i := range contacts {
		if name == contacts[i].Name {
			fmt.Println("查找成功！请输入想要修改的联系姓名:")
			finalName := ReadTrimmedInput(reader)
			fmt.Println("请输入想要修改的联系电话号码:")
			finalPhone := ReadTrimmedInput(reader)
			contacts[i].Name = finalName
			contacts[i].Phone = finalPhone
			fmt.Printf("已成功修改联系人姓名为：%s, 电话号码：%s\n", finalName, finalPhone)
			return
		}
	}
	fmt.Printf("未找到姓名为 %s 的联系人\n", name)
}
