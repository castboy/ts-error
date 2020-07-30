package kafka

import "ts-error/err_code"

type KafkaOpType err_code.ErrCodeSub
type KafkaOpObj err_code.ErrCodeSubSub

const (
	Produce KafkaOpType = iota + 1
	Consume
)

var KafkaOpTypeMsg = map[KafkaOpType]string{
	Produce: "Produce",
	Consume: "Consume",
}

const (
	Order KafkaOpObj = iota + 1
	Event
)

var KafkaOpObjMsg = map[KafkaOpObj]string{
	Order: "Order",
	Event: "Event",
}

