Concurrent batch
===

Implement a function that will load users from the database  
`getBatch(n int64, pool int64) res []user`  

The function takes two arguments - the number of users and the number of goroutines in which users will concurrently load​
```
for i := 0; i < number; i++ {
      go func() {
         user := getOne()...
      }
}
```
The function returns an array of received users  
It is planned to review the solution for the interview  

NOTE
---
Due to autocode VM limitations, run tests locally. It is possible that autocode tests will pass with an invalid solution​

Tips & tricks
---
* Don't forget the data race
* To limit concurrently running goroutines you can use:
  * Semaphore pattern
  * Worker pool
  * errgroup
  * etc ...

