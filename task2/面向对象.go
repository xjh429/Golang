package main

import (
	"fmt"
	"math"
)

// Shape 定义 Shape 接口，包含 Area 和 Perimeter 方法
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Rectangle 定义矩形结构体
type Rectangle struct {
	Width  float64
	Height float64
}

// Area 实现 Shape 接口的 Area 方法，计算矩形面积
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter 实现 Shape 接口的 Perimeter 方法，计算矩形周长
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Circle 定义圆形结构体
type Circle struct {
	Radius float64
}

// Area 实现 Shape 接口的 Area 方法，计算圆形面积
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Perimeter 实现 Shape 接口的 Perimeter 方法，计算圆形周长
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Person 定义 Person 结构体，包含姓名和年龄字段
type Person struct {
	Name string
	Age  int
}

// Employee 定义 Employee 结构体，组合 Person 结构体并添加员工 ID 字段
type Employee struct {
	Person
	EmployeeID int
}

// PrintInfo 为 Employee 结构体实现 PrintInfo 方法，输出员工信息
func (e Employee) PrintInfo() {
	fmt.Printf("姓名: %s, 年龄: %d, 员工 ID: %d\n", e.Name, e.Age, e.EmployeeID)
}

func main() {
	// 题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
	// 考察点 ：接口的定义与实现、面向对象编程风格。

	// 创建矩形和圆形实例
	rect := Rectangle{5, 3}
	circ := Circle{4}

	// 调用 Area 和 Perimeter 方法并输出结果
	fmt.Printf("矩形面积: %.2f\n", rect.Area())
	fmt.Printf("矩形周长: %.2f\n", rect.Perimeter())
	fmt.Printf("圆形面积: %.2f\n", circ.Area())
	fmt.Printf("圆形周长: %.2f\n", circ.Perimeter())

	// 	题目 ：使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
	// 考察点 ：组合的使用、方法接收者。

	// 创建 Employee 实例
	employee := Employee{
		Person: Person{
			Name: "张三",
			Age:  30,
		},
		EmployeeID: 12345,
	}

	// 调用 PrintInfo 方法输出员工信息
	employee.PrintInfo()

}
