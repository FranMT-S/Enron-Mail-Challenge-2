package workers

type WorkerState string

const (
	InProcess WorkerState = "InProcess"
	Finish    WorkerState = "Finish"
	Waiting   WorkerState = "Waiting"
	Error     WorkerState = "Error"
)
