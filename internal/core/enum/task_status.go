package enum

type TaskStatus int

const (
	Initiates  TaskStatus = 0
	InProgress TaskStatus = 1
	Done       TaskStatus = 2
	Failed     TaskStatus = 3
)
