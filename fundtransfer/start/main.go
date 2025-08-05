package main

import (
	"context"
	"log"

	"github.com/seblkma/go-temporal/fundtransfer"
	"go.temporal.io/sdk/client"
)

func main() {

	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create Temporal client:", err)
	}
	defer c.Close()

	input := fundtransfer.PaymentDetails{
		SourceAccount: "85-150",
		TargetAccount: "43-812",
		Amount:        250,
		ReferenceID:   "12345",
	}

	options := client.StartWorkflowOptions{
		ID:        "paynow-401",
		TaskQueue: fundtransfer.FundTransferTaskQueueName,
	}

	log.Printf("Starting transfer from account %s to account %s for %d", input.SourceAccount, input.TargetAccount, input.Amount)

	wex, err := c.ExecuteWorkflow(context.Background(), options, fundtransfer.MoneyTransfer, input)
	if err != nil {
		log.Fatalln("Unable to start the Workflow:", err)
	}

	log.Printf("WorkflowID: %s RunID: %s\n", wex.GetID(), wex.GetRunID())

	var result string

	err = wex.Get(context.Background(), &result)

	if err != nil {
		log.Fatalln("Unable to get Workflow result:", err)
	}

	log.Println(result)

}
