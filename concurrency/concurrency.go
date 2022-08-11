package main

func main() {}

type CheckerWebsite func(string) bool

func VerifyWebsite(vw CheckerWebsite, urls []string) map[string]bool {
	result := make(map[string]bool)

	for _, url := range urls {
		result[url] = vw(url)
	}

	return result
}
