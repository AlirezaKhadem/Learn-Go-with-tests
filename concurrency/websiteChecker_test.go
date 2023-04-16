package concurrency

import (
	"github.com/bxcodec/faker/v3"
	"reflect"
	"testing"
	"time"
)

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		"waat://example.com",
	}

	expectedOutput := map[string]bool{
		"http://google.com":          true,
		"http://blog.gypsydave5.com": true,
		"waat://example.com":         false,
	}

	output := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(expectedOutput, output) {
		t.Fatalf(
			"wanted %v, got %v",
			expectedOutput,
			output,
		)
	}
}

func BenchmarkCheckWebsites(b *testing.B) {
	numberOfWebsites := 100
	urls := make([]string, numberOfWebsites)
	for index := 0; index < numberOfWebsites; index++ {
		urls[index] = "https://www." + faker.DomainName()
	}
	b.ResetTimer()

	for index := 0; index < numberOfWebsites; index++ {
		_ = CheckWebsites(mockWebsiteChecker, urls)
	}
}

func mockWebsiteChecker(url string) bool {
	if url == "waat://example.com" {
		return false
	}
	return true
}

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(time.Second * 20)
	return false
}
