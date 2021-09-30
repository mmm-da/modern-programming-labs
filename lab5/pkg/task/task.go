package task

import "time"

type PaymentInterface interface {
	New() PaymentInterface
}

type Payment struct {
	PaymentType string
	Value       uint64
	VaultS      string
	VaultD      string
	Timestamp   time.Time
}
