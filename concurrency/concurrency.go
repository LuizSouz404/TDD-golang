package main

func main() {}

type CheckerWebsite func(string) bool

func VerifyWebsite(vw CheckerWebsite, urls []string) map[string]bool {
	result := make(map[string]bool)

	for _, url := range urls {
		go func(url string) {
			result[url] = vw(url)
		}(url)
	}

	return result
}
