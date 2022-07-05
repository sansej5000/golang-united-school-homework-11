package batch

import (
	"sync"
	"time"
)

type user struct {
	ID int64
}

func getOne(id int64) user {
	time.Sleep(time.Millisecond * 100)
	return user{ID: id}
}

func getBatch(n int64, pool int64) (res []user) {

	mu := sync.Mutex{}
	var wg sync.WaitGroup
	sem := make(chan struct{}, pool)
	for i := 0; i < int(n); i++ {
		wg.Add(1)
		sem <- struct{}{}
		go func(j int64) {
			user := getOne(j)
			mu.Lock()
			res = append(res, user)
			mu.Unlock()
			<-sem
			wg.Done()
		}(int64(i))
	}
	wg.Wait()
	return
}
