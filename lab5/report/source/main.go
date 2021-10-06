package main

import (
	"lab5/pkg/payment"
	"lab5/pkg/paymentchain"
)

// Паттерн Chain of responsibility.
// Прохождение платежа через банковскую систему сопровождается целым рядом действий: фиксирующих,
// контролирующих, снимающих процент банка и прочие вычеты и действия.
// Построить цепочки для различного вида платежей (обычных, льготных, государственных, внутрибанковских)
// в соответствии с предметной областью и разработать модель системы.

// Regular chain
// initJob -> startJob -> logJob -> transitJob -> rewardJob -> logJob -> finishJob
// Preferential chain
// initJob -> startJob -> logJob -> calcJob -> transitJob -> rewardJob -> logJob -> finishJob
// Government chain
// initJob -> startJob -> logJob -> transitJob -> rewardJob -> logJob -> finishJob -> reportJob
// Inner chain
// initJob -> startJob -> logJob -> transitJob -> logJob -> finishJob

func getRegularChain() []paymentchain.PaymentChainLink {
	chain := make([]paymentchain.PaymentChainLink, 0)

	chain = append(chain, paymentchain.InitJob{})
	chain = append(chain, paymentchain.StartJob{})
	chain = append(chain, paymentchain.LogJob{})
	chain = append(chain, paymentchain.TransitJob{})
	chain = append(chain, paymentchain.RewardJob{})
	chain = append(chain, paymentchain.LogJob{})
	chain = append(chain, paymentchain.FinishJob{})

	return chain
}

func getPreferentialChain() []paymentchain.PaymentChainLink {
	chain := make([]paymentchain.PaymentChainLink, 0)

	chain = append(chain, paymentchain.InitJob{})
	chain = append(chain, paymentchain.StartJob{})
	chain = append(chain, paymentchain.LogJob{})
	chain = append(chain, paymentchain.CalcJob{})
	chain = append(chain, paymentchain.TransitJob{})
	chain = append(chain, paymentchain.LogJob{})
	chain = append(chain, paymentchain.FinishJob{})

	return chain
}

func getGovermentChain() []paymentchain.PaymentChainLink {
	chain := make([]paymentchain.PaymentChainLink, 0)

	chain = append(chain, paymentchain.InitJob{})
	chain = append(chain, paymentchain.StartJob{})
	chain = append(chain, paymentchain.LogJob{})
	chain = append(chain, paymentchain.TransitJob{})
	chain = append(chain, paymentchain.LogJob{})
	chain = append(chain, paymentchain.FinishJob{})
	chain = append(chain, paymentchain.ReportJob{})

	return chain
}

func getInnerChain() []paymentchain.PaymentChainLink {
	chain := make([]paymentchain.PaymentChainLink, 0)

	chain = append(chain, paymentchain.InitJob{})
	chain = append(chain, paymentchain.StartJob{})
	chain = append(chain, paymentchain.LogJob{})
	chain = append(chain, paymentchain.TransitJob{})
	chain = append(chain, paymentchain.LogJob{})
	chain = append(chain, paymentchain.FinishJob{})

	return chain
}

func main() {
	paymentsCount := 100

	chainMap := map[payment.PaymentType][]paymentchain.PaymentChainLink{
		payment.Regular:      getRegularChain(),
		payment.Preferential: getPreferentialChain(),
		payment.Government:   getGovermentChain(),
		payment.Inner:        getInnerChain(),
	}

	payments := payment.GenerateMockPayments(paymentsCount)

	for _, p := range payments {
		for _, job := range chainMap[p.Type] {
			job.Execute(p)
		}
	}
}
