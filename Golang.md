### init 函数

- 在导入包时自动执行其中的 `init` 函数内容。
- 每个包可以有多个 `init` 函数，它们会按照声明顺序执行。
- `init` 函数没有参数，也不返回值。
- 主要用于包的初始化工作，例如：
  - 初始化全局变量
  - 注册回调函数
  - 连接数据库







### new 定义函数

- 用于分配内存并返回指向该内存地址的指针。
- 返回值类型为 `*T`，其中 `T` 是你想要分配的类型。
- 分配的内存被初始化为该类型的零值。
- 主要用途：创建Struct结构体、Array数组等类型的值，这些类型需要在堆上分配内存。

**示例：**

```go
type Person struct {
    Name string
    Age  int
}

p := new(Person) // p 指向一个新分配的 Person 结构体
p.Name = "Alice"
p.Age = 30
```





### 声明、初始化、定义、赋值

- 声明：

  声明变量存在的类型，但还未分配空间。

  - 例如：`var p *int` 声明了一个指向 `int` 类型变量的指针 `p`。

    

- 初始化：

  为变量赋予初始值

  - 例如：`city = "合肥"` 初始化变量 `city` 为字符串 "合肥"。

  `new 和 make 也可看作是初始化`

  

- 赋值：

  将新值赋予已声明的变量。

  - 例如：`city = "莆田"` 将变量 `city` 的值修改为 "莆田"。







### 指针、`*`、`&`

- 指针：

   

  指针变量用于存储其他变量的内存地址。

  - 例如：`var p *int` 声明了一个指向 `int` 类型变量的指针 `p`。

- `*`：

   

  用于解引用指针，获取指针指向的内存地址中的值。

  - 例如：`*p = 666` 将指针 `p` 指向的内存地址中的值设置为 666。
  - 例如：`add := *p` 将指针 `p` 指向的内存地址中的值赋值给变量 `add`。

- `&`：

   

  用于获取变量的内存地址。

  - 例如：`a := 10` 声明一个 `int` 类型的变量 `a` 并赋值为 10。
  - 例如：`p2 := &a` 获取变量 `a` 的内存地址并赋值给指针 `p2`。







### 值传递

- Go 语言中的函数参数传递为值传递，这意味着函数会复制参数的值到函数内部的局部变量中。
- 如果需要在函数内部修改参数的值，并使修改后的值在函数外部生效，则需要使用指针类型（引用）。







### 数组

- 初始化：
  - `b := [5]int{1,2,3,4,5}` 初始化一个长度为 5 的 `int` 数组 `b`，并赋值为 {1, 2, 3, 4, 5}。
  - `c := [5]int{1,2,3}` 初始化一个长度为 5 的 `int` 数组 `c`，并赋值为 {1, 2, 3, 0, 0}。
  - `d := [5]int{2: 10, 4: 20}` 初始化一个长度为 5 的 `int` 数组 `d`，并赋值为 {0, 0, 10, 0, 20}。







### 随机数

- 设置种子：`rand.Seed(time.Now().UnixNano())` 以当前时间为种子，确保每次生成的随机数不同。
- 产生随机数：`rand.Intn(n)` 生成 `n` 以内的随机 `int` 数。







### Slice 切片

- 切片是对数组的引用，可以动态调整长度和容量。
- 切片的内存中有三个值：指针、长度、容量。
- 初始化：
  - `var array = [...]int{1,2,3,4}` 初始化一个 `int` 数组 `array`。
  - `slice := array[1:3:4]` 从 `array` 的索引 1 开始，取长度为 2 的切片，容量为 4。



方法

1. `append(slice, n)`:

   在切片末尾追加成员，不改变slice本身，但返回修改后的切片。

   - 特点：如果扩容超过底层数组容量，则会以 2 倍容量进行扩容。

2. **`copy(slice1, slice2)`:** 将 `slice2` 的内容复制到 `slice1` 中，覆盖 `slice1` 的原有内容。

**示例：**

```go
srcSlice := []int{1, 2}
dstSlice := []int{6, 6, 6, 6, 6, 6}
copy(dstSlice, srcSlice) // dstSlice: [1, 2, 6, 6, 6, 6]
```







### 数组与切片的区别

- **数组：** 内存大小固定，类型为 `[N]T`，其中 `N` 是数组长度，`T` 是元素类型。
- **切片：** 内存动态分配，类型为 `[]T`，`T` 是元素类型。
- **关系：** 切片是对原数组的部分引用，可以访问和修改原数组中的元素。







### make 函数

- 用于创建并初始化切片、映射和通道这三种内置数据结构。
- 格式：`make(类型, 长度, 容量)`
- 例如：`slice := make([]int, 5, 5)` 创建长度和容量都为 5 的 `int` 切片。





### new和make

| 特性   | `new`                         | `make`                           |
| :----- | :---------------------------- | :------------------------------- |
| 作用   | 分配内存并返回指针            | 创建内置数据结构                 |
| 返回值 | 指针                          | 数据结构本身，值nil              |
| 初始化 | 初始化为零值                  | 初始化内部数据                   |
| 用途   | 创建结构体struct、数组array等 | 创建切片slice、映射map、通道chan |

- new用于定义 值类型； 返回的是指针

- make用于定义引用类型； 返回的是结构本身（引用：结构中包含指针）

~~~go
package main

import "fmt"

func main() {
	type structing struct {
		s string
	}

	var s []int
	var m map[int]string
	var c chan int
	var a *[2]string
	var b *structing

	s = make([]int, 10)
	m = make(map[int]string)
	c = make(chan int)
	a = new([2]string)
	b = new(structing)
	fmt.Println(s, m, c, *a, *b)  
    //[0 0 0 0 0 0 0 0 0 0] map[] 0xc00001a0c0 [ ] {}

	s = append(s, 1)
	m[0] = "ling"
	a[0] = "zi"
	b.s = "chen"
	fmt.Println(s, m, c, *a, *b)  
    //[0 0 0 0 0 0 0 0 0 0 1] map[0:ling] 0xc00001a0c0 [zi ] {chen}

}

~~~









### map 字典/映射  

### key -> value

~~~go
info := map[int]string {
    110: "报警电话"，
    119："消防电话"
    120: "急救电话"
}
~~~

- map可用make创建，第二个参数的意思是为指定容量，也可不写动态分配

  ~~~go
  m1 := make(map[int]string, 2)
  m1[1] = "maike"
  ~~~

- map是 `无序` 的







### Slice与Map的range遍历

- 使用 `range` 遍历 Slice 时，会返回当前元素的索引和值。

  ~~~go
  package main
  
  import "fmt"
  
  func main() {
    numbers := []int{1, 2, 3, 4, 5}
  
    // 遍历 Slice，获取索引和值
    for index, value := range numbers {
      fmt.Printf("索引: %d, 值: %d\n", index, value)
    }
  }
  ~~~

- 使用 `range` 遍历 Map 时，会返回当前键和值。（无序）

  ~~~go
  package main
  
  import "fmt"
  
  func main() {
    person := map[string]int{"Alice": 25, "Bob": 30}
  
    // 遍历 Map，获取键和值
    for key, value := range person {
      fmt.Printf("姓名: %s, 年龄: %d\n", key, value)
    }
  }
  ~~~



- 方法

  1.判断键值是否存在： 	value, ok := map1[xxx]

  2.删除键值对：				delete(map1, xxx)		//删除key为xxx的内容







### Slice与Map都是为引用类型

可直接作为函数参数传递

`虽然函数参数默认值传递，但是slice与map的值（内存）都是为储存指向底层数据结构的指针`





### 定义函数为类型做参数传递

~~~go
package main

import "fmt"

func printFunc(text string) bool {
	fmt.Println(text)
	return true
}

type onlyFunc func(string) bool

func usePrint(f onlyFunc) {
	f("we will win")
}

func main() {
	usePrint(printFunc)
}

~~~







### Struct结构体（值类型）

- 结构体就是一种类型，不过是由自己定义的： type Student struct {xxx}

- 初始化

  ~~~go
  package main
  
  import "fmt"
  
  type Person struct {
    Name string
    Age  int
  }
  
  func main() {
    // 使用字面量初始化
    p1 := Person{"Alice", 25}
    fmt.Println(p1) // 输出: {Alice 25}
      
    // 使用键值对初始化
    p2 := Person{
      Name: "Bob",
      Age:  30,
    }
    fmt.Println(p2) // 输出: {Bob 30}
      
    // 使用零值初始化
    p3 := Person{}
    fmt.Println(p3) // 输出: {}
      
    // 使用指针初始化
    p4 := &Person{"David", 40}
    fmt.Println(*p4) // 输出: {David 40}
  }
  ~~~



`在golang中，结构体通过指针操作成员： p.vari 与 (*p).vari 完全等价`	 //p是为结构体指针





### 可见性与大小写名（函数名/结构体名/结构体成员）

- **首字母大写：** 如果首字母大写，则在包外可见，可以被其他包的代码调用。
- **首字母小写：** 如果首字母小写，则在包外不可见，只能在当前包内部使用。

~~~go
// package mypackage

type MyStruct struct { // 结构体名首字母大写，在包外可见
    Name string
    Age  int
}

func MyFunc() { // 函数名首字母大写，在包外可见
    fmt.Println("Hello from MyFunc")
}

func myFunc() { // 函数名首字母小写，在包外不可见
    fmt.Println("Hello from myFunc")
}
~~~

~~~go
// package main

import "mypackage"

func main() {
    // 可以访问 MyStruct 和 MyFunc
    myStruct := mypackage.MyStruct{"John Doe", 30}
    mypackage.MyFunc()

    // 无法访问 myFunc
    // mypackage.myFunc() // 编译错误
}
~~~





### go mod 与module导包问题

- 在每个包下都取相同的 go mod init test

  ![image-20240828214619414](C:\Users\86156\AppData\Roaming\Typora\typora-user-images\image-20240828214619414.png)

  

- 在需要用包的同级mod文件中写入 require 包，和replace命名

  ~~~go
  module test
  
  go 1.22.5
  
  require vis v0.0.0
  
  replace vis => ./vis
  
  ~~~

  





### 匿名字段（继承）

- 继承与初始化

~~~go
package main

import "fmt"

type Person struct {
	name string
	sex  byte
	age  int
}

type Student struct {
	Person //继承Person 结构，没有名字匿名字段，包含Person的所有成员
	id     int
	addr   string
}

func (tmp *Person) PrintInfo() {		//写*T和T都可以，结构体默认隐式转化，一般用*T，轻量且引用能修改操作
	fmt.Println(tmp.name, tmp.sex, tmp.age)
}

func main() {
	//顺序初始化
	var s1 Student = Student{Person{"mike", 'm', 20}, 1, "合肥"} //属性继承
	s1.PrintInfo()                                             //方法继承
	//指定成员初始化, 其他默认自动为0
	s3 := Student{id: 1}
	fmt.Println(s3)

}

~~~

- 同名字段

  当Person和Student中都有同名字段，进行赋值时就本作用域赋值成员，没有再找继承中字段



### 结构体方法(封装)

go中可为任意自定义类型添加方法（封装）

在继承中，重名与重写，都默认采用就近原则。需显示调用继承的方法 s.Person.PrintInfo()

 ~~~go
 package main
 
 import "fmt"
 
 type long int
 //自定义类型封装方法
 func (l long) Add(x long) long {	//为某个类型绑定函数，即叫做方法。
     							//总共两个参数，l为自身参数，x为变化参数
 	return l + x
 }
 
 //结构体封装方法
 type Student struct {
     name string
     sex byte
 }
 
 func (tmp Student) PrintInfo() {
     fmt.Println(tmp)
 }
 func (tmp *Student) SetInfo(name string, sex byte) { 	//结构体类型与结构体指针类型可看作为同一封装方法（共用）
     												//receiver 可以时 T 或者 *T，但T不能是接口或指针
 	tmp.name = name
     tmp.sex = sex
 }
 
 func main() {
 	var a long
 	a = 5
 
 	fmt.Println(a.Add(4))
     
     s := Student{name: "zz", sex: 'm'}
     s.PrintInfo()
     
     var s3 = new(Student)
     s3.SetInfo("s3", 'w')
     fmt.Println(*s3)
     
     var s4 Student
 	s4.SetInfo("s4", 'm')	//隐式转化，只在调用方法时发生，不受方法集约束。参数传递时就不会发生隐式转化
 	fmt.Println(s4)
 }
 
 
 
 ~~~





### 接口（多态）

接口类型就是描述了一系列方法的集合。

- 各输出线相当于方法
- 接口就是集合多输出线的接口终端，方法的集合
- 各结构可以套上接口，通过统一接口实现方法
- 统一接口， 相同方法表现出不同结构输出，即为多态

~~~go
package main

import "fmt"

type Animal interface {
	Humaner //继承接口
}

type Humaner interface {
	eat()
}

type Student struct {
	name string
	age  int
}

func (s *Student) eat() {
	fmt.Printf("%v eating\n", *s)
}

type Teacher struct {
	name string
	age  int
	sex  byte
}

func (t *Teacher) eat() {
	fmt.Printf("%v eating\n", *t)
}

type Animals struct {
	types string
}

func (a *Animals) eat() {
	fmt.Printf("%v eating\n", *a)
}

func main() {
	s := &Student{"zl", 20}
	t := &Teacher{"hf", 30, 'm'}
	a := &Animals{"cat"}

	var h Animal
	array := [3]Animal{s, t, a}

	for _, data := range array {
		h = data
		h.eat()
	}
}

~~~



接口相当于聚焦化方法。方法由多向少转化，结构->超集->子集

一个结构可以套上多个接口。

- 空接口

  空接口不包含任何方法。故所有类型都能实现（套上）空接口

~~~go
package main

import "fmt"

func main() {
	// 创建数组
	var data [3]interface{}
	data[0] = "Hello"
	data[1] = 10
	data[2] = true

	// 访问数组元素
	for _, v := range data {
		fmt.Println(v)
	}
}
~~~







### 类型断言

- 反推interface{}空接口中的对应各类型

~~~go
package main

import "fmt"

type Student struct {
	name string
	id   int
}

func main() {
	i := make([]interface{}, 3)
	i[0] = 1
	i[1] = "hello golang"
	i[2] = Student{"mike", 8}
	//if 断言
	for index, data := range i {
		if value, ok := data.(int); ok {
			fmt.Printf("x[%d] 类型为int， 内容为%d\n", index, value)
		} else if value, ok := data.(string); ok {
			fmt.Printf("x[%d] 类型为string， 内容为%s\n", index, value)
		} else if value, ok := data.(Student); ok {
			fmt.Printf("x[%d] 类型为Student， 内容为%+v\n", index, value)
		}
	}

	//switch 断言
	for index, data := range i {
		switch data.(type) {
		case int:
			fmt.Printf("x[%d] 类型为int， 内容为%d\n", index, data.(int))
		case string:
			fmt.Printf("x[%d] 类型为string， 内容为%s\n", index, data.(string))
		case Student:
			fmt.Printf("x[%d] 类型为Student， 内容为%+v\n", index, data.(Student))
		}
	}
}

~~~





### 异常处理

- error接口 (源码)

~~~go
package errors

func New(text string) error { 	//函数： 根据字符串New 创造error接口类型
    return &errorString{string}
}
	
type errorString struct {		//error对象
	s string
}

func (e *errorStirng) Error() string {	//error对象的封装方法，返回具体的错误信息
	return e.s
}

//现在的error源码接口
type error interface {
    Error() string
}
~~~



- err应用

~~~go
package main

import (
	"errors"
	"fmt"
)

func myDiv(a, b int) (result int, err error) {
	if b == 0 {
		err = errors.New("b connot be zero")
	} else {
		result = a / b
	}
	return
}

func main() {
	result, err := myDiv(10, 0)
	if err != nil {
		fmt.Printf(err.Error())
	} else {
		fmt.Print(result)
	}
}

~~~



- panic函数奔溃

  ~~~go
  package main
  
  import "fmt"
  
  func main() {
    // 尝试访问一个不存在的数组元素
    arr := []int{1, 2, 3}
    fmt.Println(arr[4]) // 这里会触发 panic
  
    // 即使 panic 发生，defer 语句仍然会执行
    defer fmt.Println("defer 语句执行")
  
    // 这里代码不会执行
    fmt.Println("这行代码不会被执行")
  }
  ~~~

  

- recover函数

  使得panic出现的时候不崩溃 

  ~~~go
  defer func() {
      if err := recover(); if err != nil {
          fmt.Println(err)
      }
  }
  ~~~

  

  

### 文本文件处理

- 字符串操作
- 字符串转化
- 正则表达式



### JSON处理

- 通过结构体生成json（编码）

  ~~~go
  package main
  
  import (
  	"encoding/json"
  	"fmt"
  )
  
  type IT struct {	//可二次编码，使得json输出的是小写
  	Company  string   `json:"-"`	//不输出该字段内容
  	Subjects []string `json:"subjects"` //二次编码
  	Isok     bool     `json:",string"`	//value格式转化成string再输出
  	Price    float64  `json:"price,string"`
  }
  
  func main() {
  	s := IT{"zlican", []string{"go", "js", "python", "mysql"}, true, 666}
      jsonS, err := json.Marshal(s)  //json.MarshalIndent(s, "", " ") 这是格式化json
      if err != nil {
          fmt.Println(err)
          return
      }
      
  	fmt.Printf("%+v\n", string(jsonS)) 
  }
  
  ~~~

- 通过map生成json（编码）

- json解析到结构体（解析）

  ~~~go
  package main
  
  import (
  	"encoding/json"
  	"fmt"
  )
  
  type IT struct {
  	Company  string   `json:"company"`
  	Subjects []string `json:"subjects"`
  	Isok     bool     `json:"isok"`
  	Price    float64  `json:"price"`
  }
  
  func main() {
  	jsonBuf := `{"company":"zlican","subjects":["go","js","python","mysql"],"isok":true,"price":666}` //json
  
  	var tmp IT
  	err := json.Unmarshal([]byte(jsonBuf), &tmp) 
  	if err != nil {
  		fmt.Println("err = ", err)
  		return
  	}
  	fmt.Println(tmp)
  }
  
  ~~~

- json解析到map（解析）





### 文件操作接口

- 创建
- 写
- 读
- 删除





### goroutine协程

~~~go
package main

import (
	"fmt"
	"time"
)

func routine() {
	for {
		fmt.Println("goroutine")
		time.Sleep(time.Second)
	}
}

func main() {
	go routine()

	for {
		fmt.Println("main")
		time.Sleep(time.Second)
	}
}

~~~



### runtime包

- Gosched函数：让出时间片

~~~go
package main

import (
	"fmt"
	"runtime"
)

func main() {
	go func() {
		for {
			fmt.Println("go")
		}
	}()

	go func() {
		for {
			fmt.Println("routine")
		}
	}()

	for i := 0; i < 3; i++ {
		runtime.Gosched()
		fmt.Println("hello")
	}
}

~~~



- Goexit函数：结束当前协程

  ~~~go
  package main
  
  import (
  	"fmt"
  	"runtime"
  )
  
  func test() {
  	defer fmt.Println("ccc")
  
  	runtime.Goexit()
  
  	fmt.Println("ddd")
  }
  
  func main() {
  	go func() {
  		fmt.Println("aaa")
  
  		test()
  
  		fmt.Println("bbb")
  	}()
  
  	for {
  
  	}
  
  }
  
  ~~~

  

- GOMAXPROCS函数：设置多核CPU并行运行

~~~go
package main

import (
	"fmt"
	"runtime"
)

func main() {
	n := runtime.GOMAXPROCS(16) //16核运行
	fmt.Println(n)	//上次运行所使用的核数

	for {
		go fmt.Print(1)

		fmt.Print(0)
	}
}

~~~





### channel 

使用通道，解决多资源竞争问题（channel中没数据时则会阻塞）

利用chan的阻塞性质可以起到沟通不同协程的作用

- `var c chan int` 只是声明了一个通道变量，但没有创建通道。
- `var c = make(chan int)` 既声明了通道变量，又创建了通道。

~~~go
package main

import (
	"fmt"
	"time"
)

var c = make(chan bool)

func PrintBot(s string) {
	for _, data := range s {
		fmt.Printf("%c", data)
		time.Sleep(time.Second)
	}
}

func person1() {
	PrintBot("Hello")
	c <- true
}

func person2() {
    <-c 	//利用chan阻塞来完成同步操作，在person1没print完时 person2不会执行
	PrintBot("World")
}

func main() {
	go person1()
	go person2()

	for {

	}
}

~~~



- 无缓存通道与有缓存通道

  ~~~go
  func main() {
      c := make(chan int, 0) //无缓存
      d := make(chan int, 3  ) //有缓存
  }
  ~~~

  无缓存的chan：写了没读， c<- 就阻塞；是空要读 <-c 阻塞

  有缓存的chan：写满了， c<- 才阻塞；是空要读 <-c 阻塞

 

- 关闭channel

  ~~~go
  var ch = make(chan int)
  
  go func() {
      for i := 0; i < 5; i++ {
  		ch <- i 
      }
      close(ch)
  }()
  
  for {
      if num, ok := <-ch; ok == true {
          fmt.Println(num)
      } else {
  		break //通道关闭
      }
  }
  ~~~

  关闭channel后无法再发送数据，但可以继续接收数据

  对与nil channel，无论收发都会被阻塞（var ch chan int）



- 遍历channel

  ~~~go
  for num := range ch {
      fmt.Println(num)
  }
  ~~~

  



### 定时器

- Timer（只对通道写入一次）

  ~~~go
  package main
  
  import (
  	"fmt"
  	"time"
  )
  
  func main() {
  	//创建一个定时器，设置时间为2s,2s后，往timer通道写内容（当前时间）
  	timer := time.NewTimer(2 * time.Second)
  	fmt.Println("当前时间", time.Now())
  
  	t := <-timer.C
  	fmt.Println(t)
  }
  
  ~~~

  

- time.After()的源码就是利用通道定时器

  ~~~go
  func After(d Duration) <-chan Time {	
      return NewTimer(d).C	//创建通道，并送入d，返回通道对象
  }
  ~~~

  ~~~go
  //外部使用
  <-time.After(2 * time.Second)   //阻塞，待得从通道读取内容
  fmt.Println("print")
  ~~~

  

- Ticker（定时对通道写入）

  ~~~go
  package main
  
  import (
  	"fmt"
  	"time"
  )
  
  func main() {
  	ticker := time.NewTicker(1 * time.Second)
  
  	i := 0
  	for {
  		<-ticker.C
  		i++
  		fmt.Println(i)
          
          if i == 5 {
              ticker.Stop()  //结束定时器
              break
          }
  	}
  }
  
  ~~~

  

  

### Select 的使用

`主要用来监听channel上的数据流动`： 检测通道的状态；配合for一起使用，持续监听

~~~go
package main

import "fmt"

func fibonacci(ch chan<- int, quit <-chan bool) {
	x, y := 1, 1
	for {
		select { 	//一次只能运行一个case；for起到持续监听功能
		case ch <- x:
			x, y = y, x+y
		case flag := <-quit:
			fmt.Println(flag)
			return
		}
	}
}

func main() {
	ch := make(chan int)    //斐波那契数字通道
	quit := make(chan bool) //使用通信分享结束消息

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-ch)
		}
		quit <- true //main结束任务协程后通过通道传递结束信息
	}()

	fibonacci(ch, quit)
}

~~~



- select实现超时

  ~~~go
  package main
  
  import (
  	"fmt"
  	"time"
  )
  
  func main() {
  	ch := make(chan int)
  	quit := make(chan bool)
  
  	go func() {
  		for {
  			select {
  			case <-ch:
  				fmt.Print("success")
  			case <-time.After(3 * time.Second):
  				fmt.Println("3秒超时")
  				quit <- true
  			}
  		}
  	}()
  
  	go func() {
  		for {
  			time.Sleep(time.Second)
  			ch <- 1
  		}
  	}()
  
  	<-quit
  	fmt.Println("程序结束")
  }
  
  ~~~

  

### net tcp并发编程

- server.go

  ~~~go
  package main
  
  import (
  	"fmt"
  	"net"
  	"strings"
  )
  
  func HandleConn(conn net.Conn) {
  	defer conn.Close()
  
  	addr := conn.RemoteAddr().String()
  	fmt.Println("connection success", addr)
  
  	//读取用户数据
  	buf := make([]byte, 2048)
  	for {
  		n, err := conn.Read(buf)
  		if err != nil {
  			fmt.Println(err)
  			return
  		}
  
  		if string(buf[:n-2]) == "exit" {
  			fmt.Printf("[%s] quit\n", addr)
  			return
  		}
  
  		fmt.Printf("[%s]: %s\n", addr, string(buf[:n]))
  
  		conn.Write([]byte(strings.ToUpper(string(buf[:n]))))
  	}
  }
  
  func main() {
  	//监听
  	listener, err := net.Listen("tcp", "127.0.0.1: 8080")
  	if err != nil {
  		fmt.Println(err)
  		return
  	}
  	defer listener.Close()
  
  	//连接用户
  	for {
  		conn, err := listener.Accept()
  		if err != nil {
  			fmt.Println(err)
  			return
  		}
  
  		//处理用户请求,新建协程
  		//因为每个用户间的通信都是无穷的，所以需要单独开协程，否则阻塞
  		go HandleConn(conn)
  	}
  }
  
  ~~~

- client.go

  ~~~go
  package main
  
  import (
  	"fmt"
  	"net"
  	"os"
  )
  
  func main() {
  	//主动连接服务器
  	conn, err := net.Dial("tcp", "127.0.0.1:8080")
  	if err != nil {
  		fmt.Println(err)
  		return
  	}
  
  	ch := make(chan bool)
  
  	//main 结束时关闭连接
  	defer conn.Close()
  
  	//接收服务器回复的数据
  	go func() {
  		buf := make([]byte, 1024)
  		for {
  			n, err := conn.Read(buf)
  			if err != nil {
  				fmt.Println(err)
  				return
  			}
  			fmt.Println(string(buf[:n]))
  		}
  	}()
  
  	//读取键盘输入，发送给服务器
  	go func() {
  		str := make([]byte, 1024)
  		for {
  			n, err := os.Stdin.Read(str)
  			if err != nil {
  				fmt.Println(err)
  				return
  			}
  
  			if string(str[:n-2]) == "exit" {
  				ch <- true
  			}
  
  			conn.Write([]byte(str[:n]))
  		}
  
  	}()
  
  	<-ch
  }
  
  ~~~

  





### 什么情况下要开启Go routine

`一次性结束不了的，需要挂载引入时间维度的函数`





### CS并发聊天室

~~~go
package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

type Client struct {
	C    chan string //客户监听管道，接收消息
	Name string      //用户名
	Addr string      //用户地址
}

var onlineMap = make(map[string]Client) //在线用户
var messageChan = make(chan string)
var isQuit = make(chan bool)
var hasData = make(chan bool)

func makeMessage(cli Client, msg string) string {
	return "[" + cli.Name + "]" + msg
}

func Manager() { //处理分发消息chan给每一个用户，广播
	for {
		msg := <-messageChan
		for _, cli := range onlineMap {
			cli.C <- msg
		}
	}
}

func writeChanMessage(cli Client, conn net.Conn) {
	for msg := range cli.C {
		conn.Write([]byte(msg + "\n"))
	}
}

func writeClientMessage(cli Client, conn net.Conn) { //服务器端拿到用户的数据，广播
	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}
		msg := string(buf[:n-1])
		if len(msg) == 0 {
			isQuit <- true
		}
		if n == 4 && msg == "who" { //查询在线用户
			conn.Write([]byte("user list:\n"))
			for _, cli := range onlineMap {
				conn.Write([]byte(cli.Name + "|" + cli.Addr + "\n"))
			}
			continue
		} else if len(msg) >= 8 && msg[:6] == "rename" {
			name := strings.Split(msg, "|")[1]
			cli.Name = name
			onlineMap[cli.Addr] = cli
			conn.Write([]byte("rename " + cli.Name + " success\n"))
			continue
		} else {
			messageChan <- makeMessage(cli, msg)
		}
		hasData <- true
	}
}

// 处理用户链接
func handleConn(conn net.Conn) {
	defer conn.Close()

	cliAddr := conn.RemoteAddr().String()
	var c = make(chan string)
	cli := Client{
		C:    c,
		Name: string(cliAddr),
		Addr: string(cliAddr),
	}
	go writeChanMessage(cli, conn) //监听打印自身管道消息

	onlineMap[cli.Name] = cli //将存在用户写入服务器
	messageChan <- makeMessage(cli, "用户登入成功")
	cli.C <- makeMessage(cli, "i am here")

	go writeClientMessage(cli, conn) //广播用户的输入消息

	for {
		select {
		case <-isQuit:
			delete(onlineMap, cliAddr)
			messageChan <- makeMessage(cli, "退出聊天")
			return
		case <-hasData:

		case <-time.After(60 * time.Second): //超时机制
			delete(onlineMap, cliAddr)
			messageChan <- makeMessage(cli, "超时退出")
			return
		}
	}
}

func main() {
	//开启监听服务器
	listenner, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer listenner.Close()

	go Manager() //监听message消息广播

	for { //监听客户端与用户建立连接
		conn, err := listenner.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		// 处理用户链接
		go handleConn(conn)
	}
}

~~~







### http通信与传统tcp通信的差别

- http通信的格式，请求有相应的请求行、请求头、空行、请求体；响应有响应的响应行、响应头、空行、响应体

- 若用tcp传统方法处理http通信则繁琐冗余，于是有了相应的net/http开发库

  



### http编程

- serve

  ~~~go
  package main
  
  import (
  	"fmt"
  	"net"
  )
  
  func main() {
  
  	listener, err := net.Listen("tcp", "127.0.0.1:8080")
  	if err != nil {
  		return
  	}
  	defer listener.Close()
  
  	for {
  		conn, err := listener.Accept()
  		if err != nil {
  			return
  		}
  		defer conn.Close()
  
  		buf := make([]byte, 1024)
  		n, err2 := conn.Read(buf)
  		if n == 0 {
  			fmt.Println(err2)
  			return
  		}
  		fmt.Printf("#%v#", string(buf[:n]))
  	}
  }
  
  ~~~

  

- http_server

  ~~~go
  package main
  
  import (
  	"fmt"
  	"net/http"
  )
  
  func HandConn(w http.ResponseWriter, r *http.Request) {
  	w.Write([]byte("Hello, world!"))
  	fmt.Println(r.Method, r.RemoteAddr, r.Header, r.Body)
  }
  
  func main() {
  	http.HandleFunc("/", HandConn)
  	//用户连接处理函数
  
  	http.ListenAndServe(":8080", nil)
  }
  
  ~~~

  

- http_client

  ~~~go
  package main
  
  import (
  	"fmt"
  	"net/http"
  )
  
  func main() {
  	resp, err := http.Get("http://www.baidu.com")
  	if err != nil {
  		return
  	}
  
  	defer resp.Body.Close()
  
  	fmt.Println(resp.Status, resp.StatusCode, resp, resp.Header)
  
  	buf := make([]byte, 4096)
  	var tmp string
  
  	for {
  		n, err := resp.Body.Read(buf)
  		if n == 0 {
  			fmt.Println(err)
  			break
  		}
  
  		tmp += string(buf[:n])
  	}
  	fmt.Println(tmp)
  }
  
  ~~~







### 并发爬虫

~~~go
package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func HttpGet(url string) (cont string, err2 error) {
	resp, err := http.Get(url)
	if err != nil {
		err2 = err
		return
	}

	defer resp.Body.Close()

	buf := make([]byte, 4096)
	for {
		n, err := resp.Body.Read(buf)
		if n == 0 {
			fmt.Println(err)
			break
		}
		cont += string(buf[:n])
	}
	return
}

func DoWork(start int, end int) {
	//https://tieba.baidu.com/f?kw=%E5%8C%BA%E5%9D%97%E9%93%BE&ie=utf-8&pn=50

	for i := start; i <= end; i++ {
		//在每次for循环中爬取单个网页
		go func() { //因为此处在for过程中，操作不能一次性立刻结束，需要引入时间维度等待http响应，故因采用go routine
			cont, err := HttpGet("https://tieba.baidu.com/f?kw=%E5%8C%BA%E5%9D%97%E9%93%BE&ie=utf-8&pn=" +
				strconv.Itoa(i*50))

			if err != nil {
				return
			}

			//把内容写入文件
			fileName := strconv.Itoa(i) + ".html"
			f, err1 := os.Create(fileName)
			if err1 != nil {
				fmt.Println(err1)
				return
			}
			f.WriteString(cont)
			f.Close()
		}()
	}
}

func main() {
	var start, end int //控制爬取范围
	fmt.Printf("请输入起始页（ >= 1）：")
	fmt.Scan(&start)
	fmt.Printf("请输入终止页（ >= 起始页）：")
	fmt.Scan(&end)

	DoWork(start, end)
	for {

	}
}

~~~

