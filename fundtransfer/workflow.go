package fundtransfer

import (
	"fmt"
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

func MoneyTransfer(ctx workflow.Context, input PaymentDetails) (string, error) {

	// RetryPolicy specifies how to automatically handle retries if an Activity fails.
	retrypolicy := &temporal.RetryPolicy{
		InitialInterval:        time.Second,
		BackoffCoefficient:     2.0,
		MaximumInterval:        100 * time.Second,
		MaximumAttempts:        500, // 0 is unlimited retries
		NonRetryableErrorTypes: []string{"InvalidAccountError", "InsufficientFundsError"},
	}

	options := workflow.ActivityOptions{
		// Timeout options specify when to automatically timeout Activity functions.
		StartToCloseTimeout: time.Minute,
		// Optionally provide a customized RetryPolicy.
		// Temporal retries failed Activities by default.
		RetryPolicy: retrypolicy,
	}

	ctx = workflow.WithActivityOptions(ctx, options)

	// Withdraw, Deposit, Refund are activities

	var withdrawOutput string
	withdrawErr := workflow.ExecuteActivity(ctx, Withdraw, input).Get(ctx, &withdrawOutput)
	if withdrawErr != nil {
		return "", withdrawErr
	}

	var depositOutput string
	depositErr := workflow.ExecuteActivity(ctx, Deposit, input).Get(ctx, &depositOutput)
	if depositErr != nil {
		// Transfer failed, put money back to source account
		var result string
		refundErr := workflow.ExecuteActivity(ctx, Refund, input).Get(ctx, &result)
		if refundErr != nil {
			return "", fmt.Errorf("Deposit: failed to deposit money into %v: %v. Money could not be returned to %v: %w",
				input.TargetAccount, depositErr, input.SourceAccount, refundErr)
		}

		return "", fmt.Errorf("Deposit: failed to deposit money into %v: Money returned to %v: %w",
			input.TargetAccount, input.SourceAccount, depositErr)
	}

	result := fmt.Sprintf("Transfer complete (transaction IDs: %s, %s)", withdrawOutput, depositOutput)
	return result, nil
}
