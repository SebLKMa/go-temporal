package fundtransfer

const FundTransferTaskQueueName = "TRANSFER_FUND_TASK_QUEUE"

type PaymentDetails struct {
	SourceAccount string
	TargetAccount string
	Amount        int
	ReferenceID   string
}
