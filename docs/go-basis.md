# Go-basis

### 数据类型

- 布尔类型

- 数字类型

  - 整数类型

  - 浮点数 float32 float64

  - 复数

    ```
    uint8  无符号8位整型   0-255
    uint16 无符号16位整型  0-65535
    uint32 无符号32位整型  0-42亿
    uint64 无符号64位整型  0-18446744073709551615
    int8   有符号8位整型   -128-127
    int16  有符号16位整型 -32768-32767
    int32  有符号32位整型 -21亿-21亿
    int64。有符号64位整型  -9223372036854775808- 9223372036854775807
    float32 IEEE-754 32位浮点型数
    float64 IEEE-754 64位浮点型数
    complex32  32位实数和虚数
    complex64  64位实数和虚数
    ```

    ```
    byte 类似uint8
    rune 类似int32
    uint 32/64
    int =uint
    uintptr 无符号整数 存放指针
    ```

    

- 字符串类型  

  ```
  单个字节连接起来的字符串 UTF-8标识
  ```

- 派生类型

  - 指针类型
  - 数组类型
  - 结构化类型
  - Channel类型
  - 函数类型
  - 切片类型
  - 接口类型
  - Map类型

### 变量

#### 定义

```go
var a int=1
var a,b int =1,2
var (
  a int=1
  b int=2
)
类型推断
var a="jasmine"
name必须为新的变量否则编译错误
name :="jasmine"
```

### 条件语句

##### if

##### switch 支持多条件匹配

```
switch var1{
   case val1:
  	 ...
   case val2:
  	 ...
   default:
  	 ...
}
```

```
判断某个 interface 变量中实际存储的变量类型
switch x.(type){
	 case type:
 		 ...
	 case type:
		 ...
	 defaluct:
		 ...
}
```

```
failthuough 不会去判断下一个是否成立
switch{
	case true:
		...
		failthrough
	case false:
		...
		
}
```

##### Select

```
随机执行一个可运行的case，没有可运行的就阻塞，知道有case可运行
```

