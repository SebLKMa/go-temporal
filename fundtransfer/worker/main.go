package main

import (
	"log"

	"github.com/seblkma/go-temporal/fundtransfer"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create Temporal client:", err)
	}
	defer c.Close()

	w := worker.New(c, fundtransfer.FundTransferTaskQueueName, worker.Options{})
	w.RegisterWorkflow(fundtransfer.MoneyTransfer)
	w.RegisterActivity(fundtransfer.Withdraw)
	w.RegisterActivity(fundtransfer.Deposit)
	w.RegisterActivity(fundtransfer.Refund)

	// Start listening to the Task Queue.
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("unable to start Worker", err)
	}
}
