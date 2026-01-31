# Day 2 学习笔记 - Go基础语法

## 变量声明

### 4种方式
1. `var name string = "value"` - 完整声明
2. `var name = "value"` - 类型推导
3. `name := "value"` - 简短声明（只能在函数内）
4. 批量声明（使用var()）

### 零值
- int: 0
- float64: 0.0
- string: ""
- bool: false

## 常量

### 定义
- 使用 `const` 关键字
- 不能用 `:=`
- 编译时确定，不可修改

### iota
- 常量计数器，从0开始
- 每增加一行，iota自动+1
- 常用于枚举

## 类型系统

### 基本类型
- 整型：int8, int16, int32, int64, uint8, uint16, uint32, uint64, int, uint
- 浮点：float32, float64
- 布尔：bool
- 字符串：string
- 字符：rune(int32), byte(uint8)

### 类型转换
- **必须显式转换**，不支持隐式转换
- 语法：`目标类型(值)`
- 示例：`float64(100)`, `int(3.14)`

## 重要提醒
- Go的int类型大小取决于平台（32位或64位）
- string的len()返回字节数，不是字符数（中文1个字符=3字节）
- := 只能在函数内使用，不能用于包级别变量

## 控制流

### if语句
- 基本if、if-else、if-else if
- **if的简短语句**（重要）：`if num := 10; num > 0 { ... }`
- 变量作用域仅在if块内

### switch语句
- 自动break，不需要手动写
- 可以有简短语句
- 无条件switch：`switch { case xxx: ... }`
- fallthrough：强制执行下一个case

### for循环
- Go唯一的循环语句
- 3种形式：标准for、while式、无限循环
- break、continue
- range遍历：`for index, value := range slice { ... }`

## 数组、切片、字典

### 数组（定长）
- 声明：`var arr [5]int` 或 `arr := [5]int{1,2,3,4,5}`
- 长度固定，是值类型
- 传递会复制整个数组

### 切片（动态数组）
- 创建：`[]int{1,2,3}` 或 `make([]int, len, cap)`
- 动态长度，是引用类型
- **append**：追加元素，容量不足时扩容
- **copy**：复制切片
- 扩容机制：通常扩容为原来的2倍

### map（字典）
- 创建：`make(map[K]V)` 或 `map[K]V{k:v}`
- 增：`m[k] = v`
- 查：`v, ok := m[k]`
- 删：`delete(m, k)`
- 遍历：`for k, v := range m { ... }`
- **无序**、引用类型

## 重要提醒
- 切片和map都是引用类型，修改会影响原数据
- 数组是值类型，传递会复制
- map遍历是无序的
- range遍历时，index和value都是副本