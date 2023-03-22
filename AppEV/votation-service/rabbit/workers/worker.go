package workers

type Worker interface {
	
	Use(...WorkerMiddlewareFunc)
	
	Health() error
	
	Config(address string) error
	
	Close() error
	
	Send(queue string, messages ...[]byte) error
}

type WorkerMiddlewareFunc func(name string) WorkerMiddleware

type WorkerMiddleware interface {
	Start()
	Stop(error)
}

// Channel that can be closed in order to stop listening without having to close the whole worker
type Channel interface {
	Close() error
}

func BuildRabbitWorker(address string) (worker Worker, err error) {
	worker = new(rabbitWorker)
	err = worker.Config(address)
	return
}
