package main

func main() {}

type CheckerWebsite func(string) bool
type resultChannelStruct struct {
	string
	bool
}

func VerifyWebsite(vw CheckerWebsite, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan resultChannelStruct)

	for _, url := range urls {
		go func(url string) {
			resultChannel <- resultChannelStruct{url, vw(url)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		result := <-resultChannel
		results[result.string] = result.bool
	}

	return results
}
