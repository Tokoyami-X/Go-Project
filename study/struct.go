// package main

// import "fmt"

// type Person struct {
// 	Name string
// 	Age  int
// }

// func (p Person) Error() string {
// 	return "姓名不能为空"
// }

// func (p *Person) setName(name string) error {
// 	if name == "" {
// 		fmt.Println(p.Error())
// 	}
// 	p.Name = name
// 	return nil
// }
// func showStruct() {
// 	p := &Person{"张三", 18}
// 	p.setName("李四")
// 	// fmt.Println("%p\n", p)
// 	fmt.Println(p)
// 	fmt.Printf("%p\n", p)

// }

package main

import (
	"errors"
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

// Person 不再实现 error 接口，规避所有陷阱

func (p *Person) setName(name string) error {
	if name == "" {
		// 出错时，返回标准的 error
		return errors.New("姓名不能为空")
	}
	// 成功时，修改字段，并返回 nil
	p.Name = name
	return nil
}

func showStruct() {
	// 创建对象
	p := &Person{"张三", 18}

	// 调用方法，传入 "李四"
	err := p.setName("李四")

	// 判断是否有错误
	if err != nil {
		fmt.Println("修改失败:", err)
	} else {
		// 成功：打印结构体和地址
		fmt.Println("修改成功:", p)
		fmt.Printf("内存地址: %p\n", p)
	}
}
