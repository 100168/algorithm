package learning

import (
	"fmt"
	"sync"
	"testing"
)

func TestMutex(t *testing.T) {
	var mu sync.Mutex
	count := 0
	//定义10个线程组，超了会panic
	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		i := i
		go func() {
			fmt.Println(i)
			//执行完一个线程组就少一个
			defer wg.Done()
			fmt.Println(count)
			for i := 0; i < 10000; i++ {
				mu.Lock()
				count++
				mu.Unlock()
			}
		}()
	}
	//等待线程执行完毕
	wg.Wait()
	fmt.Println(count)
}
