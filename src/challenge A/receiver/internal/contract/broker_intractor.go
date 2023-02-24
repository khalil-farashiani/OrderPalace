package contract

type BrokerIntractor interface {
	ReceiveOrder(data chan []byte)
}
