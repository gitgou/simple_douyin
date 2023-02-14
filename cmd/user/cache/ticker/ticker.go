package ticker

import (
	"time"
	"github.com/gitgou/simple_douyin/cmd/user/cache"
)

func Ticker5() {
	ticker5 := time.NewTicker(1 * time.Second)
	defer ticker5.Stop()

	for range ticker5.C {
		cache.StoreDB()
	}
}
