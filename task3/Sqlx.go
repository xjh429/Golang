// 题目1：使用SQL扩展库进行查询
// 假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。
// -- 创建一个 employees 表，包含字段 id 、 name 、 department 、 salary 。
// CREATE TABLE employees (
//     id INT PRIMARY KEY,
//     name VARCHAR(255),
//     department VARCHAR(255),
//     salary DECIMAL(10, 2)
// );
// -- 插入一些测试数据
// INSERT INTO employees (id, name, department, salary) VALUES
//     (1, '张三', '技术部', 80000.00),
//     (2, '李四', '技术部', 60000.00),
//     (3, '王五', '销售部', 75000.00);
要求 ：
编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。
package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// 定义Employee结构体
type Employee struct {
	ID         int     `db:"id"`
	Name       string  `db:"name"`
	Department string  `db:"department"`
	Salary     float64 `db:"salary"`
}

// 查询技术部员工
func getTechEmployees(db *sqlx.DB) ([]Employee, error) {
	var employees []Employee
	err := db.Select(&employees, "SELECT * FROM employees WHERE department = ?", "技术部")
	return employees, err
}

// 查询最高工资员工
func getHighestPaidEmployee(db *sqlx.DB) (Employee, error) {
	var employee Employee
	err := db.Get(&employee, "SELECT * FROM employees ORDER BY salary DESC LIMIT 1")
	return employee, err
}

func main() {
	// 连接数据库
	db, err := sqlx.Connect("mysql", "root:1234@tcp(localhost:3306)/db_shard_0")
	if err != nil {
		fmt.Println("数据库连接失败:", err)
		return
	}
	defer db.Close()

	// 查询技术部员工
	techEmployees, err := getTechEmployees(db)
	if err != nil {
		fmt.Println("查询技术部员工失败:", err)
		return
	}
	fmt.Println("技术部员工:", techEmployees)

	// 查询最高工资员工
	highestPaid, err := getHighestPaidEmployee(db)
	if err != nil {
		fmt.Println("查询最高工资员工失败:", err)
		return
	}
	fmt.Println("最高工资员工:", highestPaid)

}
// 题目2：实现类型安全映射
// 假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。
// # 1. 创建 books 表
// CREATE TABLE books (
//     id INT PRIMARY KEY,
//     title VARCHAR(255),
//     author VARCHAR(255),
//     price DECIMAL(10, 2)
// );
// # 2. 插入一些示例数据
// INSERT INTO books (id, title, author, price) VALUES
//     (1, '《论语》', '孔子', 39.99),
//     (2, '《资治通鉴》', '司马光', 29.99),
//     (3, '《史记》', '司马迁', 49.99),
//     (4, '《论语2》', '孔子', 59.99),
//     (5, '《资治通鉴2》', '司马光', 69.99)

// 要求 ：
// 定义一个 Book 结构体，包含与 books 表对应的字段。
// 编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。
package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)
type Book struct {
	ID     int     `db:"id"`
	Title  string  `db:"title"`
	Author string  `db:"author"`
	Price  float64 `db:"price"`
}

// 查询价格大于50元的书籍
func getExpensiveBooks(db *sqlx.DB) ([]Book, error) {
	var books []Book
	err := db.Select(&books, "SELECT * FROM books WHERE price > ?", 50.0)
	return books, err
}

func main() {
	// 连接数据库
	db, err := sqlx.Connect("mysql", "root:1234@tcp(localhost:3306)/db_shard_0")
	if err != nil {
		fmt.Println("数据库连接失败:", err)
		return
	}
	defer db.Close()

	// 查询价格大于50元的书籍
	expensiveBooks, err := getExpensiveBooks(db)
	if err != nil {
		fmt.Println("查询高价书籍失败:", err)
		return
	}
	fmt.Println("价格大于50元的书籍:", expensiveBooks)
}

