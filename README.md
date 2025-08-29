# go-temporal
Applications using temporal for fault tolerance.

Go to your temporal cli directory,  
```sh
./temporal server start-dev --db-filename your_temporal.db --ui-port 8080
```

## fundtransfer

### Start the workflow
A workflow contains the activity tasks, e.g. withdraw, deposit etc.  
Starting a workflow enqueues the tasks.  
```sh
go run start/main.go
```

### Start the worker
Worker executes the tasks from the queue.  
```sh
go run worker/main.go
```

### Test transaction recovery

To test auto-recovery:  

[Recover from a server crash](https://learn.temporal.io/getting_started/go/first_program_in_go/#recover-from-a-server-crash)  

[Recover from an unknown error in an Activity](https://learn.temporal.io/getting_started/go/first_program_in_go/#recover-from-an-unknown-error-in-an-activity)  
***Just fix the issue, re-start the worker.***  

## References

https://learn.temporal.io/getting_started/go/first_program_in_go/  


## [An example of URL Shortening Service](./README_surl.md)