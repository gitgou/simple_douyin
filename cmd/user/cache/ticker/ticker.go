package ticker

import (
	"time"
	"github.com/gitgou/simple_douyin/cmd/user/cache"
)

func Ticker3M() {
	ticker1M := time.NewTicker(60 * time.Second)
	defer ticker1M.Stop()

	for range ticker1M.C {
		cache.StoreDB()
	}
}
