package context

// 自定义中间层，模拟gin的中间件

type H map[string]any

type Task struct {
	Data H
	funcIndex int // 表示进行到第funcIndex个函数
	funcList []func(*Task)
}

func NewTask(data H,procedure ...func(*Task)) Task {
	if data == nil {
		data = make(H)
	}
	return Task{funcList: procedure,Data: data}
}

func (task *Task) Next() {
	task.funcIndex ++
	task.Do()
}

func (task *Task) Do() {
	procedureLen := len(task.funcList)
	for task.funcIndex < procedureLen {
		task.funcList[task.funcIndex](task)
		task.funcIndex ++
	}
}
