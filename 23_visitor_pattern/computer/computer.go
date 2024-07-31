package computer

import "errors"

// 为了方便直接使用切片替代寄存器

type CPU struct {
	methods map[string]func(*Memory,...int) error
	memory *Memory
}

func (cpu *CPU) accept(visitor Visitor) {
	visitor.VisitCPU(cpu)
}

func (cpu *CPU) Update(methodName string,method func(*Memory,...int) error) error {
	_,ok := cpu.methods[methodName]
	if ok {
		return errors.New("操作名已被占用")
	}
	cpu.methods[methodName] = method
	return nil
}

func (cpu *CPU) Calc(methodName string, param ...int) error {
	method,ok := cpu.methods[methodName]
	if ok {
		method(cpu.memory,param...)
		return nil
	} else {
		return errors.New("操作未找到，操作名为 "+methodName)
	}
}

type Memory struct {
	memory  []int  // 使用一维数组模拟内存空间
	methods map[string]func(*[]int,*[]int,...int) error
	Out []int  // 内存输出
}

func (memory *Memory) accept(visitor Visitor) {
	visitor.VisitMemory(memory)
}

func (memory *Memory) Update(methodName string,method func(*[]int,*[]int,...int) error ) error{
	_,ok := memory.methods[methodName]
	if ok {
		return errors.New("操作名已被占用")
	}
	memory.methods[methodName] = method
	return nil
}


func (memory *Memory) Calc(methodName string, params ...int) error {
	method,ok := memory.methods[methodName]
	if ok {
		method(&memory.memory,&memory.Out,params...)
		return nil
	} else {
		return errors.New("操作未找到，操作名为 "+methodName)
	}
}


type Computer struct {
	Cpu CPU
	memory Memory
}

func (computer *Computer) Update(visitor Visitor) {
	computer.Cpu.accept(visitor)
	computer.memory.accept(visitor)
} 

func NewComputer() *Computer{
	memory := Memory{memory: make([]int, 10),methods: make(map[string]func(*[]int, *[]int, ...int) error)} // 初始开辟十个字的内存
	cpu := CPU{memory: &memory,methods: make(map[string]func(*Memory, ...int) error)}
	computer := &Computer{Cpu: cpu,memory: memory}
	return computer
}


type Visitor struct {
	VisitCPU func(*CPU)
	VisitMemory func(*Memory)
}

