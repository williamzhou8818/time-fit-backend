package test

import (
	"sync"
	"testing"

	"github.com/williamzhou8818/time-fit-backend/database"
)

// 测试高并发下gorm.DB是否还是单例
func TestGetBlogDBConnection(t *testing.T) {
	const C = 100
	wg := sync.WaitGroup{}
	wg.Add(C)
	for i := 0; i < C; i++ {
		go func() {
			defer wg.Done()
			database.GetTimeToFitDBConnection()
		}()
	}
	wg.Wait()
}

// go test -v .\database\test\ -run=^TestGetBlogDBConnection$ -count=1
