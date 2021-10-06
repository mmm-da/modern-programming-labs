package payment

import (
	"lab5/pkg/errorhandlers"
	"math/rand"
	"time"

	uuid "github.com/nu7hatch/gouuid"
)

type Payment struct {
	Type      PaymentType
	Value     uint64
	VaultSrc  string
	VaultDst  string
	Timestamp time.Time
}

func New(paymentType PaymentType, srcVault, dstVault string, value uint64, time time.Time) Payment {
	return Payment{Type: paymentType, VaultSrc: srcVault, VaultDst: dstVault, Value: value, Timestamp: time}
}

func GenerateMockPayments(count int) []Payment {
	rand.Seed(time.Now().Unix())
	resultList := make([]Payment, 0)

	for i := 0; i < count; i++ {

		srcVault, err := uuid.NewV4()
		errorhandlers.IfErrorStartPanic(err)

		dstVault, err := uuid.NewV4()
		errorhandlers.IfErrorStartPanic(err)

		value := rand.Uint64()
		paymentType := rand.Int63n(int64(countOfPaymentType()))

		time := time.Unix(rand.Int63(), 0)

		resultList = append(resultList, Payment{Type: PaymentType(paymentType), VaultSrc: srcVault.String(), VaultDst: dstVault.String(), Value: value, Timestamp: time})
	}

	return resultList
}
