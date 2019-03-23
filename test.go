package main

import (
	"errors"
	"fmt"
	"time"
)

/*
<<<<<<< HEAD
go outline 
	语法简单
	编译速度快
	内置并发机制
		goroutine
			配置在一组逻辑处理器,每个逻辑处理器绑定到一个操作系统线程
		channel 
			通信，自带同步
	类型系统简单高效
			组合设计模式
			接口隐式实现
	自带gc
		编译器去处理
web
	http://play.golang.org 
go 快速浏览
	包
		定义编译单元
		import 
		包 init 
	符号导出
		首字母大写
	引用类型
		channel, map, slice 
	func 
		匿名函数
		闭包
	defer 
	接口interface 
		value-receiver: use value or ptr 
		pointer-receiver: use value or ptr 
	goroutine 
	channel
	
go 打包和工具链
	包名惯例
		目录名字
		不要求所有的包名称都不同
	main包
		main方法
	导入
		本地导入
			标准库:GOROOT
			第三方:GOPATH 
		远程导入
			go get 
		命名导入
			多个包具有相同名字，使用命名导入
			import myfmt "mylib/fmt"
	函数init
		导入包的时候触发
	go工具
		go build //构建 
		go clean //清除
		go run 　//构建＋清除
		go vet //错误检查
		go fmt //格式调整
		go 语言文档
			go doc 
				命令行方式
			godoc -http:=6060
				web方式
	go开发者合作
		包处于代码库的根目录
		包可以非常小
		对代码执行go fmt 
		给代码写文档
	依赖管理
		godep
		vender
		gopkg.in 
		go module //1.11

	空白标识符
		忽略变量，忽略包名(避免编译错误)


=======
go包
	首字母大写的字符才能导出
>>>>>>> 71aeb557c0dc494fcfaf9db1cfa078f177c1b9de
go-type
	bool
		true/false
	number
		uint 6/16/32/64
		int 6/16/32/64
		byte //alias for uint8
		rune //alias for int32
		complex 32/64
		uintptr
	string

	derived
		pointer/array/struct/channel/function/slice/interface/map
zero-value 
	0 for number
	false for bool
	"" for string
var define
	var var_name var_type = var_value
	var var_name = var_value
	var_name := var_value
	var (
		var_name1 var_type1
		var_name2 var_type2
	)
value-type
	int/string/float/bool
ref-type
	&var_name
init-declare
	var_name := var_value
empty-identifier
	_
const-var
	const var_name [var_type] = var-value
	const (
		var_name1 = var_value1
		var_name2 = var_value2
		var_name3 = var_value3
	)

	iota
	
	can't use ":="
operator
	math-operator
		+ - * / % ++ --
	relation-operator
		== != > >= < <=
	logical-operator
		&& || !
	bit-operator
		& ^ | << >>
	assign-operator
		= += -= * / /= %= <<= &= ^= |=
	other
		& *
if-cond-statement
	if cond-expr{
		...
	}

	if cond-expr{
		...
	}else if cond-expr{
		...
	}else{
		...
	}
	
	if short-statement; cond-expr{
		...
	}

switch-cond-statement
	switch var1{
		case v1:
			...
		case v2:
			...
		case v3:
			...
		default:
			...
	}
	
	"break” is added automatically
	"case" may not be const var
	"var1" may be anytype
	
	switch-with-no-condition
		same as "switch true"
	
	type-switch
	switch x.(type){
		case type1:
			...
		case type2:
			...
		case type3:
			...
		default:
			...
	}

	fallthrough

loop-statement
	for init; cond; post {}
	for cond {}
	for {}
loop-control
	break
	continue
	goto
function
	func function_name( [parameter list] ) [return_types] {
   		函数体
	}
	name-return-values
	
	param
		value-pass
			default
			func swap(x, y int) int { }
		ref-pass
			pointer
			func swap(x *int, y *int) { }
	function-use
		function-var
			getSquareRoot := func(x float64) float64 {
				  return math.Sqrt(x)
			   }
		lambda-exp
			func getSequence() func() int {
			   i:=0
			   return func() int {
				  i+=1
				 return i
			   }
		method
			func (variable_name variable_data_type) function_name() [return_type]{
			}
var-domain
	local-var
		in function
	global-var
		out function
	pattern-param
defer
	推迟函数的执行
	stacking-defer
array
	declare
		var variable_name [SIZE] variable_type
	init
		var balance = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
		var balance = [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	access
		var val = balance[1]

	multi-dim declare
		var variable_name [SIZE1][SIZE2]...[SIZEN] variable_type
	multi-dim init
		var variable_name [SIZE1][SIZE2]...[SIZEN] variable_type = { {...}, {...}, {...}}
	multi-dim acess
		variable_name[dim1][dim2]
	pass to function
		myFunction(param [10] int)
		myFunction(param [] int)


}

pointer
	var ptr *type;

	define pointer
		var ptr*type
	assign pointer
		ptr = &var
	access pointer
		*ptr
struct
	type struct_variable_type struct {
	   member definition
	   member definition
	   ...
	   member definition
	}

	var-init
		variable_name := structure_variable_type {value1, value2...valuen}
		variable_name := structure_variable_type { key1: value1, key2: value2..., keyn: valuen}
	access member
		variable_name.member
	pass to function
		func printBook( book Books )
	struct pointer
slice
	abstract of array
		flexible view into the element of array
	slice likes the references of array of underlying array
	define
		var var_name []type
	create
		make([]type, len [, capacity])
	init
		s := [] int{1,2,3}
		s := array[[startIdx]:[endIdx]]

	slice-default
		zero for low-bound
		length for high-bound
	len()
		the length of slice, you can re-slice to extend or shrink
	cap()

	nil-slice
		zero of slice whith len and cap of 0
	slice-cut
		slice_name[low_bound:high_bound]
	slices of slices
	append
		the backing array may be allocated
	copy
Range
	like iterator for array/slice/channel/map
Map
	implement by hash_map
	define
		var map_variable map[key_data_type]value_data_type
		map_variable := make(map[key_data_type]value_data_type)
	delete
		delete(map_name, key)
recursive-func

method
	function with receiver-arg
	method for struct or no-struct
	pointer-receiver
		can change the variable
		you can use both pointer or no-pointer 
	no-pointer-receiver
		you can use both pointer or no-pointer 
	choose pointer-receiver or no-pointer-receiver?
		always choose pointer-receiver for modification or value-copy
interface
	a set of method signarues
	interfaces are implemented implicitly
	interface-value 
		(value, concrete-type)
	interface with nil underlying value 
		nil-receiver
	nil-interface 
		with no-type and no-value
	empty-interface 
		interface {} //interface with zero method
		used for unknow type
		
	type interface_name interface {
		method_name1 [return_type]
		method_name2 [return_type]
		method_name3 [return_type]
		...
		method_namen [return_type]
     	}

	type struct_name struct {
	}

	func (struct_name_variable struct_name) method_name1() [return_type] {
	}
	...
	func (struct_name_variable struct_name) method_namen() [return_type] {
	}
type-assertions
	providing access to underlying value of interface 
	t, ok := i.(T)
	
	if i is not type T, ok is false
type-switch
	serveral type-assertion in series
type-convert
	type_name(expression)
stringers
	builtin-inteface
	type Stringer interface {
    		String() string
	}
error
	builtin-interface 
	type error interface {
		Error() string
	}
reader 
	func (T) Read(b []byte) (n int, err error)
	
goroutine
	//lightweight thread managed by go runtime
	go 函数名( 参数列表 )
channel
	pipe for sender/receiver data 
	by default 
		the sender/recever is blocked until the other side is ready
	ch := make(chan int)
    	ch <- v    // 把v发送到通道ch
	v := <-ch  // 从ch接收数据, 并把值赋给v
	
	buffered-channel
		ch := make(chan int, cap)
	range-close
		v, ok := <-ch
		if ok is false, then ch is closed, sending to a closed ch may cause panic

select
	let a goroutint wait on multiple communication operations
	default-select
sync.mutex
other:
	IDE
		LiteIDE
		Sublime
		goland
		vim
		vscode
	gdb for debug
	go command tool
		build/clean/doc/env/fix/fmt/generate/install/list
		run/test/tool/version/vet
		others

logmain
*/

const (
	a = iota
	b = iota
	c = iota
	d
)

func max(val1, val2 int) int {
	var res int
	if val1 < val2 {
		res = val2
	} else {
		res = val1
	}
	return res
}

func swap(x, y string) (string, string) {
	return y, x
}

func getSequence() func() int {
	i:=0
	return func() int {
		i+=1
		return i
	}
}

type Circle struct {
	radius float64
}

func (c Circle) getArea() float64 {
	return 3.14 * c.radius * c.radius
}

type Books struct {
	title string
	author string
	subject string
	book_id int
}

func printSlice(x []int){
	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}

func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-2) + fibonacci(n-1)
}

type Phone interface {
	call()
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}

type IPhone struct {
}

func (iPhone IPhone) call() {
	fmt.Println("I am iPhone, I can call you!")
}


type error interface {
	Error() string
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("math: square root of negative number")
	}

	return 1, nil
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 把 sum 发送到通道 c
}

func fibonacci_channel(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

//testmain


func main() {
	/*
	fmt.Println("Hello, World!")

	var value1 int = 10
	var value2 = 10
	value3 := 10
	fmt.Print(value1, value2, value3)
	fmt.Print(&value1, &value2, &value3)

	fmt.Print(a, b, c, d)

	for a := 0; a < 10; a++ {
		fmt.Printf("a 的值为: %d\n", a)
	}

	switch true {
	case false:
		fmt.Println("1、case 条件语句为 false")
		fallthrough
	case true:
		fmt.Println("2、case 条件语句为 true")
		fallthrough
	case false:
		fmt.Println("3、case 条件语句为 false")
		fallthrough
	case true:
		fmt.Println("4、case 条件语句为 true")
	case false:
		fmt.Println("5、case 条件语句为 false")
		fallthrough
	default:
		fmt.Println("6、默认 case")
	}

	a, b := swap("Mahesh", "Kumar")
	fmt.Println(a, b)

	nextNumber := getSequence()
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())

	nextNumber1 := getSequence()
	fmt.Println(nextNumber1())
	fmt.Println(nextNumber1())
	fmt.Println(nextNumber1())

	var c1 Circle
	c1.radius = 10.00
	fmt.Println("圆的面积 = ", c1.getArea())

	var var1 int = 10;
	var ptr *int
	fmt.Printf("var1[%d],ptr[%x]\n", var1, ptr)

	ptr = &var1
	fmt.Printf("var1[%d],ptr[%x]\n", var1, ptr)

	fmt.Println(Books{"Go 语言", "www.runoob.com", "Go 语言教程", 6495407})
	fmt.Println(Books{title: "Go 语言", author: "www.runoob.com", subject: "Go 语言教程", book_id: 6495407})
	fmt.Println(Books{title: "Go 语言", author: "www.runoob.com"})

	var book Books
	book.author = "author"
	book.book_id = 100
	book.subject = "test"
	book.title = "test"
	fmt.Println(book)

	var ptr *Books = &book
	fmt.Println(ptr.title)

	var numbers = make([]int, 3, 5)
	fmt.Printf("len[%d], cap[%d], slice[%v]\n", len(numbers), cap(numbers), numbers)

	var numbers []int
	printSlice(numbers)

	numbers = append(numbers, 0)
	printSlice(numbers)

	numbers = append(numbers, 1)
	printSlice(numbers)

	numbers = append(numbers, 2,3,4)
	printSlice(numbers)

	numbers1 := make([]int, len(numbers), (cap(numbers))*2)

	copy(numbers1,numbers)
	printSlice(numbers1)

	//for slice
	nums := []int{2, 3, 4}
	for idx, num := range nums {
		fmt.Printf("idx[%d], num[%d]\n", idx, num)
	}

	//for map
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("key[%s] -> val[%s]\n", k, v)
	}

	//for string
	for i, c := range "go" {
		fmt.Printf("idx[%d], val[%d]\n", i, c)
	}

	//var countryCapitalMap map[string]string
	countryCapitalMap := make(map[string]string)

	countryCapitalMap [ "France" ] = "Paris"
	countryCapitalMap [ "Italy" ] = "罗马"
	countryCapitalMap [ "Japan" ] = "东京"
	countryCapitalMap [ "India " ] = "新德里"

	for country, captal := range countryCapitalMap {
		fmt.Printf("country[%s], captal[%s]\n", country, captal)
	}

	captal, ok := countryCapitalMap [ "美国" ]
	fmt.Printf("captal[%s],ok[%d]\n", captal, ok)
	if (ok) {
		fmt.Println("美国的首都是", captal)
	} else {
		fmt.Println("美国的首都不存在")
	}

	//
	fmt.Println("------------------------------")
	delete(countryCapitalMap, "Japan")
	for country, captal := range countryCapitalMap {
		fmt.Printf("country[%s], captal[%s]\n", country, captal)
	}

	fmt.Printf("fibonacci[%d]", fibonacci(10))

	var phone Phone
	phone = new(NokiaPhone)
	phone.call()
	phone = new(IPhone)
	phone.call()

	var sum int = 17
	var count int = 5
	var mean float32

	mean = float32(sum)/float32(count)
	fmt.Printf("mean 的值为: %f\n",mean)

	result, err:= Sqrt(-1)
	fmt.Println(result, err)

	go say("world_1")
	go say("world_2")
	say("world_main")

	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)

	x, y := <-c, <-c // 从通道 c 中接收
	fmt.Println(x, y, x+y)

	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	*/

	c := make(chan int, 10)
	go fibonacci_channel(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}
