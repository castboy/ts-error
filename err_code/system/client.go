package system

import "ts-error/err_code"

type ClientType err_code.ErrCodeSubSub

const (
	Kafka ClientType = iota + 1
	Mysql
	Redis
)

var ClientTypeMsg = map[ClientType]string {
	Kafka: "Kafka",
	Mysql: "Mysql",
	Redis: "Redis",
}
