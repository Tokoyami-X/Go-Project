package main

import "fmt"

type Struct struct {
	Name  string
	Phone string
}

func (c *Struct) setName(name string) {
	c.Name = name
}

// func (c *Struct) setPhone(phone string) {
// 	c.Phone = phone
// }

// main 函数是程序的入口点
/*
这是一个主函数，用于演示如何创建和遍历联系人切片
每个联系人包含姓名和电话号码两个属性
程序会打印出所有联系人的姓名和电话号码
*/
func main() {
	fmt.Println("1、添加联系人 2、删除联系人 3、修改联系人 4、查询联系人 5、退出程序\n请输入你想使用的功能编号按Enter键确认：")
	// 创建一个联系人切片，包含三个联系人元素
	// 每个元素都是一个 Contact 结构体，包含 Name 和 Phone 字段
	contacts := []Struct{
		{"Alice", "123-456-7890"},   // 第一个联系人：Alice，电话号码：123-456-7890
		{"Bob", "987-654-3210"},     // 第二个联系人：Bob，电话号码：987-654-3210
		{"Charlie", "555-555-5555"}, // 第三个联系人：Charlie，电话号码：555-555-5555
	}
	contacts = append(contacts, Struct{"David", "111-111-1111"})
	// 遍历 contacts 切片，使用 for-range 循环
	// 对于每个联系人 c，打印其姓名和电话号码
	for i, c := range contacts {
		fmt.Printf("第%d个联系人：%s, 电话号码：%s\n", i+1, c.Name, c.Phone)
	}
	// // 辅助函数：将字符串转换为字符串指针
	// strPtr := func(s string) *string { return &s }

	// 如果要将Name转成指针类型，可以这样做：
	// // 创建一个联系人切片，包含三个联系人元素
	// // 每个元素都是一个 Contact 结构体，包含 Name 和 Phone 字段
	// contacts := []Contact{
	// 	{strPtr("Alice"), "123-456-7890"},   // 第一个联系人：Alice，电话号码：123-456-7890
	// 	{strPtr("Bob"), "987-654-3210"},     // 第二个联系人：Bob，电话号码：987-654-3210
	// 	{strPtr("Charlie"), "555-555-5555"}, // 第三个联系人：Charlie，电话号码：555-555-5555
	// }
	// contacts = append(contacts, Contact{strPtr("David"), "111-111-1111"})
	// // 遍历 contacts 切片，使用 for-range 循环
	// // 对于每个联系人 c，打印其姓名和电话号码
	// for i, c := range contacts {
	// 	name := ""
	// 	if c.Name != nil {
	// 		name = *c.Name
	// 	}
	// 	fmt.Printf("第%d个联系人：%s, 电话号码：%s\n", i+1, name, c.Phone)
	// }
}
