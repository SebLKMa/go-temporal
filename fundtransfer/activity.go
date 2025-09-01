package fundtransfer

import (
	"context"
	"fmt"
	"log"
)

func Withdraw(ctx context.Context, data PaymentDetails) (string, error) {
	log.Printf("Withdrawing $%d from account %s.\n\n",
		data.Amount,
		data.SourceAccount,
	)

	referenceID := fmt.Sprintf("%s-withdrawal", data.ReferenceID)
	bank := BankingService{"some-API-from-some-bank.com"}
	txnId, err := bank.Withdraw(data.SourceAccount, data.Amount, referenceID)
	return txnId, err
}

func Deposit(ctx context.Context, data PaymentDetails) (string, error) {
	return DepositSuccess(ctx, data)
}

func DepositSuccess(ctx context.Context, data PaymentDetails) (string, error) {
	log.Printf("Depositing $%d into account %s.\n\n",
		data.Amount,
		data.TargetAccount,
	)

	referenceID := fmt.Sprintf("%s-deposit", data.ReferenceID)
	bank := BankingService{"some-API-from-some-bank.com"}
	// Uncomment the next line and comment the one after that to simulate an unknown failure
	//txnId, err := bank.DepositThatFails(data.TargetAccount, data.Amount, referenceID)
	txnId, err := bank.Deposit(data.TargetAccount, data.Amount, referenceID)
	return txnId, err
}

func DepositFailure(ctx context.Context, data PaymentDetails) (string, error) {
	log.Printf("Depositing $%d into account %s.\n\n",
		data.Amount,
		data.TargetAccount,
	)

	referenceID := fmt.Sprintf("%s-deposit", data.ReferenceID)
	bank := BankingService{"some-API-from-some-bank.com"}
	txnId, err := bank.DepositThatFails(data.TargetAccount, data.Amount, referenceID)
	return txnId, err
}

func Refund(ctx context.Context, data PaymentDetails) (string, error) {
	return RefundSuccess(ctx, data)
}

func RefundSuccess(ctx context.Context, data PaymentDetails) (string, error) {
	log.Printf("Depositing $%d into account %s.\n\n",
		data.Amount,
		data.TargetAccount,
	)

	referenceID := fmt.Sprintf("%s-deposit", data.ReferenceID)
	bank := BankingService{"some-API-from-some-bank.com"}
	// Uncomment the next line and comment the one after that to simulate an unknown failure
	//txnId, err := bank.DepositThatFails(data.TargetAccount, data.Amount, referenceID)
	txnId, err := bank.Deposit(data.TargetAccount, data.Amount, referenceID)
	return txnId, err
}

func RefundFailure(ctx context.Context, data PaymentDetails) (string, error) {
	log.Printf("Depositing $%d into account %s.\n\n",
		data.Amount,
		data.TargetAccount,
	)

	referenceID := fmt.Sprintf("%s-deposit", data.ReferenceID)
	bank := BankingService{"some-API-from-some-bank.com"}
	txnId, err := bank.DepositThatFails(data.TargetAccount, data.Amount, referenceID)
	return txnId, err
}
