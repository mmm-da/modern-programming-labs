package payment

type PaymentType int64

const (
	regularStr      = "Regular"
	preferentialStr = "Preferential"
	governmentStr   = "Government"
	innerStr        = "Inner"
)
const (
	Regular PaymentType = iota
	Preferential
	Government
	Inner
)

var (
	paymentTypeMap = map[PaymentType]string{
		Regular:      regularStr,
		Preferential: preferentialStr,
		Government:   governmentStr,
		Inner:        innerStr,
	}
)

func (t PaymentType) String() string {
	v, ok := paymentTypeMap[t]
	if !ok {
		panic("unknown type")
	}
	return v
}

func countOfPaymentType() int {
	return len(paymentTypeMap)
}
