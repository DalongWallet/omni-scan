package scan


type Task struct{
	Value string
}

type TaskQueue struct {
	tasks []*Task
	rptr int
}

func (tq *TaskQueue) GetTask() *Task {
	if tq.AllFinished() {
		return nil
	}
	return tq.tasks[0]
}

func (tq *TaskQueue) MarkTaskDone() {
	if tq.AllFinished()	{
		return
	}
	tq.tasks[0], tq.tasks[tq.rptr] = tq.tasks[tq.rptr], tq.tasks[0]
	tq.rptr--
}

func (tq *TaskQueue) AllFinished() bool {
	if len(tq.tasks) <= 0 || tq.rptr < 0 {
		return true
	}
	return false
}

func NewTaskQueue(values []string) *TaskQueue {
	var tasks []*Task
	for _, v := range values {
		tasks = append(tasks, &Task{
			Value:v,
		})
	}
	return &TaskQueue{
		tasks: tasks,
		rptr: len(tasks)-1,
	}
}
