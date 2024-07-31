package main

import (
	"errors"
	"fmt"
	"strconv"

	"visit/computer"
)

func memoryRead(memory *[]int,out *[]int,params ...int) error {
	// 第一个参数是起始指针
	if len(params) != 1 {
		return errors.New("参数数量应为1个")
	}
	if len(*memory) - 1 < params[0]{
		return errors.New("指针越界，已开辟内存空间范围是0~"+strconv.Itoa(len(*memory)))
	}
	*out = []int{(*memory)[params[0]]}
	return nil
}

// func memoryApply(memory *[]int,out *[]int,params ...int) error {
// 	// 第一个参数是长度，第二个参数是初始化的值
// 	if len(params) != 2 {
// 		return errors.New("参数数量应为2个")
// 	}
// 	rear := make([]int,params[0])
// 	for i := range rear {
// 		rear[i] = params[1]
// 	}
// 	*memory = append(*memory, rear...)
// 	return nil
// }


func memoryWrite(memory *[]int,out *[]int,params ...int) error  {
	// 第一个参数是指针，第二个是写入值
	if len(params) != 2 {
		return errors.New("参数数量应为2个")
	}
	if params[0] <0 || params[0] >= len(*memory) {
		return errors.New("指针越界，已开辟内存空间范围是0~"+strconv.Itoa(len(*memory)))
	}
	(*memory)[params[0]] = params[1]
	return nil
}

func cpuAdd(memory *computer.Memory,params ...int) error {
	// 第一个参数是左值指针，第二个参数是右值指针，加完的结果会放到左值指针上
	if len(params) != 2 {
		return errors.New("参数数量应为2个")
	}

	err := memory.Calc("Read",params[0])
	if err != nil {
		return err
	}
	left := memory.Out[0]

	err = memory.Calc("Read",params[1])
	if err != nil {
		return err
	}
	right := memory.Out[0]

	err = memory.Calc("Write",params[0],left+right)
	if err != nil {
		return err
	}
	return nil
}

func cpuPrint(memory *computer.Memory,params ...int) error {
	// 第一个参数是要展示的数据的指针
	if len(params) != 1 {
		return errors.New("参数数量应为1个")
	}
	err := memory.Calc("Read",params[0])
	if err != nil {
		return err
	}
	fmt.Println(memory.Out[0])
	return nil
}

func cpuAssign(memory *computer.Memory,params ...int) error {
	// 内存写的wrapper
	err := memory.Calc("Write",params...)
	return err
}

func CPUUpdate(cpu *computer.CPU) {
	cpu.Update("Add",cpuAdd)
	cpu.Update("Print",cpuPrint)
	cpu.Update("Assign",cpuAssign)
}

func MemoryUpdate(memory *computer.Memory) {
	memory.Update("Read",memoryRead)
	// memory.Update("Apply",memoryApply)
	memory.Update("Write",memoryWrite)
}

func main() {
	com := computer.NewComputer()
	updatePack := computer.Visitor{VisitCPU: CPUUpdate,VisitMemory: MemoryUpdate}
	com.Update(updatePack)

	err := com.Cpu.Calc("Assign",0,1)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = com.Cpu.Calc("Assign",1,1)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = com.Cpu.Calc("Add",0,1)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = com.Cpu.Calc("Print",0)
	if err != nil {
		fmt.Println(err)
		return
	}
}