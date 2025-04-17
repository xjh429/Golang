-- 题目1：基本CRUD操作
-- 假设有一个名为 students 的表，包含字段 id （主键，自增）、 name （学生姓名，字符串类型）、 age （学生年龄，整数类型）、 grade （学生年级，字符串类型）。
  create table students
(
    id    int auto_increment
        primary key,
    name  varchar(50) not null,
    age   int         null,
    grade varchar(20) null
);
-- 编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
  INSERT INTO students (name, age, grade) VALUES ('张三', 20, '三年级');
-- 编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
  SELECT * FROM students WHERE age > 18;
-- 编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
  UPDATE students SET grade = '四年级' WHERE name = '张三';
-- 编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
  DELETE FROM students WHERE age < 15;

-- 题目2：事务语句
-- 假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）和 transactions 表（包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。
CREATE TABLE `accounts` (
  `id` INT NOT NULL AUTO_INCREMENT COMMENT '账户ID',
  `balance` DECIMAL(10, 2) NOT NULL DEFAULT 0.00 COMMENT '账户余额',
  PRIMARY KEY (`id`)
);
CREATE TABLE `transactions` (
  `id` INT NOT NULL AUTO_INCREMENT COMMENT '转账ID',
  `from_account_id` INT NOT NULL COMMENT '转出账户ID',
  `to_account_id` INT NOT NULL COMMENT '转入账户ID',
  `amount` DECIMAL(10, 2) NOT NULL COMMENT '转账金额',
  PRIMARY KEY (`id`)
);
-- 编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。在事务中，需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。
-- 账户A
INSERT INTO accounts (balance) VALUES (200);
-- 账户B
INSERT INTO accounts (balance) VALUES (200);

-- 开始转账事务
START TRANSACTION;

-- 设置变量
SET @accountA_id = 1;
SET @accountB_id = 2;
SET @transfer_amount = 100.00;

-- 获取并锁定账户A的余额
SELECT balance INTO @current_balance FROM accounts WHERE id = @accountA_id FOR UPDATE;

-- 使用条件判断（MySQL特有语法）
SET @success = IF(@current_balance >= @transfer_amount, 1, 0);

-- 如果余额足够，执行转账操作
UPDATE accounts SET balance = balance - IF(@success, @transfer_amount, 0)
WHERE id = @accountA_id;

UPDATE accounts SET balance = balance + IF(@success, @transfer_amount, 0)
WHERE id = @accountB_id;

-- 记录交易（仅在成功时）
INSERT INTO transactions (from_account_id, to_account_id, amount)
SELECT @accountA_id, @accountB_id, @transfer_amount
WHERE @success = 1;

-- 根据结果提交或回滚
COMMIT;

-- 返回结果
SELECT IF(@success, '转账成功', '转账失败：余额不足') AS result;
