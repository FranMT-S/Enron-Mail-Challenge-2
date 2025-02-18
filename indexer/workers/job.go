package workers

type Result[Out any] struct {
	Value Out
	Err   error
}

type Job[In any, Out any] func(In) (result Out, err error)
