package interfaces

type RabbitCommunication interface {
	Send([]byte) error
}
