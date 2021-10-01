package paymentchain

import (
	"fmt"
	"lab5/pkg/payment"
)

type PaymentChainLink interface {
	Execute(p payment.Payment) error
}

type InitJob struct{}

func (j InitJob) Execute(p payment.Payment) error {
	fmt.Print("\n\nInitJob -> ")
	return nil
}

type StartJob struct{}

func (j StartJob) Execute(p payment.Payment) error {
	fmt.Print("StartJob -> ")
	return nil
}

type LogJob struct {
	text string
}

func (j LogJob) Execute(p payment.Payment) error {
	fmt.Print("LogJob -> ", j.text)
	return nil
}

type CalcJob struct{}

func (j CalcJob) Execute(p payment.Payment) error {
	fmt.Print("CalcJob -> ")
	return nil
}

type TransitJob struct{}

func (j TransitJob) Execute(p payment.Payment) error {
	fmt.Print("TransitJob")
	fmt.Printf("transfer %d value from %s to %s -> ", p.Value, p.VaultSrc, p.VaultDst)
	return nil
}

type RewardJob struct {
	rewardSize uint64
}

func (j RewardJob) Execute(p payment.Payment) error {
	fmt.Print("RewardJob -> ")
	fmt.Printf("revard %d value from %s -> ", j.rewardSize, p.VaultSrc)
	return nil
}

type FinishJob struct{}

func (j FinishJob) Execute(p payment.Payment) error {
	fmt.Print("FinishJob")
	return nil
}

type ReportJob struct{}

func (j ReportJob) Execute(p payment.Payment) error {
	fmt.Print(" -> ReportJob")
	return nil
}
