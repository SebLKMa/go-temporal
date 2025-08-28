# go-temporal
Applications using temporal for fault tolerance.

## fundtransfer
```sh
temporal server start-dev --db-filename your_temporal.db --ui-port 8080
```

```sh
go run start/main.go
```

### Test transaction recovery

To test auto-recovery:  

[Recover from a server crash](https://learn.temporal.io/getting_started/go/first_program_in_go/#recover-from-a-server-crash)  

[Recover from an unknown error in an Activity](https://learn.temporal.io/getting_started/go/first_program_in_go/#recover-from-an-unknown-error-in-an-activity)  
***Just fix the issue, re-start the worker.***  

## References

https://learn.temporal.io/getting_started/go/first_program_in_go/  


## [An example of URL Shortening Service](./README_surl.md)