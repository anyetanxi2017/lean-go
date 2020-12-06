package concurrency

import (
	"testing"
	"time"
)

func slowStuWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Microsecond)
	return true
}
func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}
	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStuWebsiteChecker, urls)
	}
}
