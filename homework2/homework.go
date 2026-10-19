package main

import ("fmt";"sync"; "time";"math";"sync/atomic")

// IncreaseByTen: 接收一个整数指针，将指针指向的值增加10
func IncreaseByTen(ptr *int){
	*ptr += 10
	
	fmt.Printf("IncreaseByTen，ptr 的值为: %d\n", *ptr)
}

func main(){
	fmt.Printf("IncreaseByTen-----------------------------------------\n")
	// 1. 声明一个整数变量num
	num := 5
	fmt.Printf("调用函数前，num 的值为: %d\n", num)
	
	// 2.调用函数，传入num变量的地址（&num）
	IncreaseByTen(&num)
	
	// 3. 输出修改后的值
	fmt.Printf("调用函数后，num 的值为: %d\n", num)
	
	
	fmt.Printf("DoubleSlice-----------------------------------------\n")
	// 1. 声明一个整数切片
	nums := []int{1,2,3,4,5}
	fmt.Printf("调用函数前: %v\n", nums)
	
	// 2. 调用函数，传入切片的地址
	DoubleSlice(&nums)
	
	// 3.输出修改后的切片
	fmt.Printf("调用函数后：%v\n", nums)
	
	
	fmt.Printf("PrintOdd and PrintEven-----------------------------------------\n")
	// 1. 声明一个 WaitGroup 变量，用于等待协程结束
	var wg sync.WaitGroup
	
	// 2. 设置需要等待的协程数量为 2
	wg.Add(2)
	
	// 3. 使用 go 关键字启动两个协程
	go PrintOdd(&wg)
	go PrintEven(&wg)
	
	//4.主协程阻塞等待，直到wg的计数变为0即两个子协程都执行完毕）
	wg.Wait()
	fmt.Printf("所有协程执行完毕，主程序退出。\n")
	
	
	// TaskScheduler
	fmt.Printf("任务调度器-------------------------------------------------------------\n")
	// 1. 创建调度器
	scheduler := NewTaskScheduler()
	
	// 2. 添加任务
	scheduler.AddTask("快速任务", taskFast)
	scheduler.AddTask("慢速任务", taskSlow)
	scheduler.AddTask("中速任务", taskMedium)
	
	// 3. 开始执行
	fmt.Printf("开始并发执行任务...")
	start := time.Now()
	scheduler.Run()
	fmt.Printf("\n总耗时: %v\n", time.Since(start))
	
	// ==================== 4. 主函数测试 ====================
	fmt.Printf("定义一个Shape接口，包含Area()和Perimeter()两个方法-------------------------------------------------------------\n")
	// 创建矩形实例
	rect := Rectangle {
		Width:10,
		Height:5}
	
	// 创建圆形实例
	circ := Circle {
		Radius : 7,
	}
	
	// 声明一个Shape接口类型的切片，用来演示多态
	shapes := []Shape{rect, circ}
	
	// 遍历切片，调用接口方法
	for _, Shape := range shapes {
		// 使用 %T 打印出具体的类型（Rectangle 或 Circle）
		fmt.Printf("形状类型: %T\n", Shape)
		fmt.Printf("  面积: %.2f\n", Shape.Area())
		fmt.Printf("  周长: %.2f\n", Shape.Perimeter())
		fmt.Printf("--------------------------------")
	}
	
	
	// 使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段
	fmt.Printf("使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段-------------------------------------------------------------\n")
	// 1.创建 Employee 实例
	emp := Employee {
		Person: Person {
			Name : "张三",
			Age : 28,
		},
		EmployeeID: "EMP-001",
	}
	// 2.调用方法
	emp.PrintInfo()
	
	// 题目：编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
	// 考察点：通道的基本使用、协程间通信。
	fmt.Printf("使用通道实现两个协程之间的通信-------------------------------------------------------------\n")
	// 3. 创建一个无缓冲的整数通道
	fmt.Printf("创建一个无缓冲的整数通道\n")
	ch := make(chan int)
	
	// 4. 启动两个协程，并发执行
	go Producer(ch)
	go Consumer(ch)
	
	// 5. 主协程休眠 1 秒，等待两个子协程执行完毕
	time.Sleep(1 * time.Second)
	fmt.Println("主程序退出。")
	
	
	
	fmt.Printf("实现一个带有缓冲的通道，生产者协程向通道中发送100个整数-------------------------------------------------------------\n")
	// 1. 创建一个容量为 10 的带缓冲通道
	// 这意味着生产者可以一口气连发 10 个数据而不被阻塞
	ch2 := make(chan int, 10)
	
	// 2. 使用 WaitGroup 优雅地等待协程结束（替代 time.Sleep）
	var wg2 sync.WaitGroup
	// 有两个协程需要等待
	wg2.Add(2)
	
	// 3. 启动生产者和消费者协程
	go Producer2(ch2, &wg2)
	go Consumer2(ch2, &wg2)
	
	// 4. 等待两个协程都执行完毕
	wg2.Wait()
	fmt.Printf("主程序退出。")
	fmt.Printf("\n")
	
	
	
	// 题目 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
	fmt.Printf("编写一个程序，使用 sync.Mutex 来保护一个共享的计数器-------------------------------------------------------------\n")
	// 1. 使用 WaitGroup 等待 10 个协程全部执行完毕
	var wg3 sync.WaitGroup
	wg3.Add(10) // 启动 10 个协程
	
	// 2. 循环启动 10 个协程
	for i:=0; i<10; i++ {
		go func(){
			defer wg3.Done() //每个协程结束时，计数器减 1
			
			// 每个协程对计数器递增 1000 次
			for j:=0; j<1000; j++ {
				// 加锁：每次修改前必须获取锁
				mu3.Lock()
				counter++
				// 解锁：修改完毕后立刻释放锁，让其他协程有机会获取
				mu3.Unlock()
			}
		}()
	}
	
	// 3.等待所有 10 个协程执行完毕
	wg3.Wait()
	
	// 4.输出最终结果
	fmt.Printf("sync.Mutex最终计数器的值: %d\n", counter)
	
	
	
	// 题目 ：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
	// 1. 声明一个 int64 类型的计数器（注意：原子操作要求类型必须严格匹配，通常是 int64）
	fmt.Printf("使用原子操作（ sync/atomic 包）实现一个无锁的计数器-------------------------------------------------------------\n")
	var counter2 int64
	
	// 2. 使用 WaitGroup 等待 10 个协程执行完毕
	var wg4 sync.WaitGroup
	wg4.Add(10)
	
	// 3. 启动 10 个协程
	for i:=0; i<10; i++ {
		go func() {
			defer wg4.Done()
			
			// 每个协程循环 1000 次
			for j:=0; j<1000; j++ {
				// 核心操作：原子递增，无需加锁，底层由硬件保证同步
				// 传入的是 counter 的内存地址（&counter）
				atomic.AddInt64(&counter2,1)
			}
		}()
	}
	
	// 4. 等待所有协程完成
	wg4.Wait()
	
	// 5. 输出最终结果
	finalCount := atomic.LoadInt64(&counter2)
	fmt.Printf("sync/atomic最终计数器的值: %d\n", finalCount)
}



// 题目 ：实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
// 考察点 ：指针运算、切片操作。
// DoubleSlice 接收一个整数切片的指针，将切片中每个元素乘以 2
func DoubleSlice(ptr *[]int){
	// 1.解引用指针，拿到切片本身
	slice := *ptr
	
	// 2.遍历切片，修改每个元素
	for i:=0; i < len(slice); i++ {
		slice[i] *=2 
	}
}




// 题目 ：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
// 考察点 ： go 关键字的使用、协程的并发执行。
// PrintOdd 打印 1 到 10 的奇数
func PrintOdd(wg *sync.WaitGroup){
	defer wg.Done()
	for i:=1; i<= 10; i+=2{
		fmt.Printf("奇数协程: %d\n", i)
	}
}

// PrintEven 打印 2 到 10 的偶数
func PrintEven(wg *sync.WaitGroup){
	defer wg.Done()
	for i:=2; i<=10; i+=2 {
		fmt.Printf("偶数协程: %d\n", i)
	}
}




// TaskFunc 定义任务类型：一个无参数，无返回值的函数
type TaskFunc func()

// TaskScheduler 任务调度器
type TaskScheduler struct {
	tasks []TaskFunc  // 任务列表
	results map[string]time.Duration // 存储每个任务的执行时间
	mu   sync.Mutex   // 保护 results 的互斥锁
}


// NewTaskScheduler 创建一个新的调度器
func NewTaskScheduler() *TaskScheduler{
	return &TaskScheduler {
		tasks : []TaskFunc{},
		results: make(map[string]time.Duration),
	}
}

// AddTask 添加任务
func (s *TaskScheduler) AddTask(name string, task TaskFunc){
	// 这里使用闭包，把 name 和 task 绑定在一起
	s.tasks = append(s.tasks, func(){
		// 记录开始时间
		start := time.Now()
		
		// 执行任务
		task()
		
		// 计算耗时
		elapsed := time.Since(start)
		
		// 加锁保护 map 的写入
		s.mu.Lock()
		s.results[name] = elapsed
		s.mu.Unlock()
		
		fmt.Printf("任务 [%s] 执行完毕，耗时: %v\n", name, elapsed)
		
	})
}

// Run 并发执行所有任务
func (s *TaskScheduler) Run() {
	var wg sync.WaitGroup
	
	// 1.设置需要等待的协程数量
	wg.Add(len(s.tasks))
	
	// 2. 并发执行每个任务
	for _, task := range s.tasks {
		go func(t TaskFunc){
			defer wg.Done()
			t()
		}(task)
	}
	
	// 3. 等待所有任务完成
	wg.Wait()
	
	// 4. 打印汇总结果
	fmt.Println("\n===== 任务执行汇总 =====")
	for name, duration := range s.results {
		fmt.Printf("%s: %v\n", name ,duration)
	}
}

// ==================== 测试用的任务 ====================
func taskFast() {
	time.Sleep(100 * time.Millisecond)
}

func taskSlow() {
	time.Sleep(500 * time.Millisecond)
}

func taskMedium() {
	time.Sleep(300 * time.Millisecond)
}






// ==================== 1. 定义接口 ====================
// Shape 接口：包含了求面积和求周长两个方法
type Shape interface {
	Area() float64
	Perimeter() float64
}

// ==================== 2. 定义一个Rectangle(矩形)结构体 ====================
type Rectangle struct {
	Width  float64
	Height float64
}

// 为 Rectangle 实现 Area 方法
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// 为 Rectangle 实现 Perimeter 方法
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// ==================== 3. 定义Circle(圆形)结构体 ====================
type Circle struct {
	Radius float64
}

// 为 Circle 实现 Area 方法
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// 为 Circle 实现 Perimeter 方法
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}


// 题目 ：使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。
// 为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。

// 1. 定义 Person 结构体
type Person struct {
	Name string
	Age int
}

// 2. 定义 Employee 结构体，组合（嵌入）Person
type Employee struct {
	Person // 匿名嵌入 Person
	EmployeeID string
	
}

// 3. 为 Employee 实现 PrintInfo 方法
func (e Employee) PrintInfo() {
	// 因为匿名嵌入了 Person，可以直接使用 e.Name 和 e.Age 访问
	fmt.Printf("员工信息: \n")
	fmt.Printf("姓名: %s\n", e.Name)
	fmt.Printf("年龄: %d\n", e.Age)
	fmt.Printf("工号: %s\n", e.EmployeeID)
}



// 题目 ：编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
// 考察点 ：通道的基本使用、协程间通信。

// ==================== 1. 生产者：生成数据并发送到通道 ====================
// ch chan<- int 表示这是一个“只写/发送”通道
func Producer(ch chan<- int) {
	for i:= 1; i<=10; i++ {
		fmt.Printf("[生产者] 发送数据：%d\n", i)
		ch <- i 
	}
	close(ch)
}

// ==================== 2. 消费者：从通道接收数据并打印 ====================
// ch <-chan int 表示这是一个“只读/接收”通道
func Consumer(ch <-chan int) {
	for num := range ch {
		fmt.Printf("[消费者] 接收到数据： %d\n", num)
	}
	fmt.Printf("[消费者] 通道已关闭，接收完毕。")
}



// 题目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
// 考察点 ：通道的缓冲机制。
// Producer 生产者：向带缓冲的通道发送 100 个整数
func Producer2(ch2 chan<-int, wg2 *sync.WaitGroup) {
	defer wg2.Done() // 任务结束，计数器减 1
	
	for i :=1; i<=100; i++ {
		ch2 <-i
	}
	// 生产完毕，关闭通道，通知消费者没有更多数据了
	close(ch2)
}

// Consumer 消费者：从通道中接收数据并打印
func Consumer2(ch2 <-chan int, wg2 *sync.WaitGroup) {
	defer wg2.Done()
	
	count := 0
	
	// 遍历通道，直到通道被关闭且数据全部取完
	for num:= range ch2 {
		fmt.Printf("[消费者] 接收到数据: %d\n", num)
		count++
	}
	fmt.Printf("[消费者] 接收完毕，共接收了 %d 个数据。\n", count)
}



// 题目 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
// 考察点 ： sync.Mutex 的使用、并发数据安全。

// 定义一个全局共享的计数器
var counter int

// 定义一个全局互斥锁
var mu3 sync.Mutex

