package omnicore



type TaskQueue struct {
	tasks []string
	rptr int
}

func (tq *TaskQueue) GetTask() string {
	if tq.AllFinished() {
		return ""
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

// TODO: add most retry times
func NewTaskQueue(tasks []string) *TaskQueue {
	return &TaskQueue{
		tasks: tasks,
		rptr: len(tasks)-1,
	}
}
