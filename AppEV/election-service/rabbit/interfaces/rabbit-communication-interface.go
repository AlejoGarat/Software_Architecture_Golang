package interfaces

type RabbitCommunication interface {
	Send([]byte) error
	SendSignal([]byte) error
}
